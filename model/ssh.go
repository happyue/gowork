package model

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gowork/log"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
)

//UpGrader websocket.Upgrader
var UpGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//Machine ssh服务端结构
type Machine struct {
	Model
	Name     string `json:"name"`
	Host     string `json:"host"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Key      string `json:"key"`
	Type     string `json:"type"`
	UserID   uint   `json:"userID"`
}

//NewSSHClient new ssh client
func (h *Machine) NewSSHClient() (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         time.Second * 5,
		User:            h.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if h.Type == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(h.Password)}
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(h.Key)}
	}
	addr := fmt.Sprintf("%s:%d", h.Host, h.Port)
	c, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	log.Debug(keyPath)
	if err != nil {
		log.Errorf("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Errorf("ssh key file read failed", err)
	}
	// CreateUserOfRole the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Errorf("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}

type safeBuffer struct {
	buffer bytes.Buffer
	mu     sync.Mutex
}

func (w *safeBuffer) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Write(p)
}
func (w *safeBuffer) Bytes() []byte {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Bytes()
}
func (w *safeBuffer) Reset() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.buffer.Reset()
}

const (
	wsMsgCmd    = "cmd"
	wsMsgResize = "resize"
)

type wsMsg struct {
	Type string `json:"type"`
	Cmd  string `json:"cmd"`
	Cols int    `json:"cols"`
	Rows int    `json:"rows"`
}

type waitGroupWrapper struct {
	sync.WaitGroup
}

func (w *waitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}

//LogicSSHWsSession LogicSshWsSession struct
type LogicSSHWsSession struct {
	stdinPipe       io.WriteCloser
	comboOutput     *safeBuffer //ssh 终端混合输出
	logBuff         *safeBuffer //保存session的日志
	inputFilterBuff *safeBuffer //用来过滤输入的命令和ssh_filter配置对比的
	session         *ssh.Session
	wsConn          *websocket.Conn
	isAdmin         bool
	IsFlagged       bool `comment:"当前session是否包含禁止命令"`
	WaitGroup       waitGroupWrapper
	exit            chan bool
}

//NewLogicSSHWsSession new LogicSshWsSession
func NewLogicSSHWsSession(cols, rows int, isAdmin bool, sshClient *ssh.Client, wsConn *websocket.Conn) (*LogicSSHWsSession, error) {
	sshSession, err := sshClient.NewSession()
	if err != nil {
		return nil, err
	}

	stdinP, err := sshSession.StdinPipe()
	if err != nil {
		return nil, err
	}

	comboWriter := new(safeBuffer)
	logBuf := new(safeBuffer)
	inputBuf := new(safeBuffer)
	//ssh.stdout and stderr will write output into comboWriter
	sshSession.Stdout = comboWriter
	sshSession.Stderr = comboWriter

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echo
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	// Request pseudo terminal
	if err := sshSession.RequestPty("xterm", rows, cols, modes); err != nil {
		return nil, err
	}
	// Start remote shell
	if err := sshSession.Shell(); err != nil {
		return nil, err
	}
	return &LogicSSHWsSession{
		stdinPipe:       stdinP,
		comboOutput:     comboWriter,
		logBuff:         logBuf,
		inputFilterBuff: inputBuf,
		session:         sshSession,
		wsConn:          wsConn,
		isAdmin:         isAdmin,
		IsFlagged:       false,
		exit:            make(chan bool),
	}, nil
}

//ReceiveWsMsg  receive websocket msg do some handling then write into ssh.session.stdin
func (sws *LogicSSHWsSession) ReceiveWsMsg() {
	wsConn := sws.wsConn

	for {
		select {
		case <-sws.exit:
			return
		default:
			//read websocket msg
			_, wsData, err := wsConn.ReadMessage()
			if err != nil {
				log.Error("reading webSocket message failed")
				return
			}
			//unmashal bytes into struct
			msgObj := wsMsg{}
			if err := json.Unmarshal(wsData, &msgObj); err != nil {
				log.Error("unmarshal websocket message failed")
			}
			switch msgObj.Type {
			case wsMsgResize:
				//handle xterm.js size change
				if msgObj.Cols > 0 && msgObj.Rows > 0 {
					if err := sws.session.WindowChange(msgObj.Rows, msgObj.Cols); err != nil {
						log.Error("ssh pty change windows size failed")
					}
				}
			case wsMsgCmd:
				//handle xterm.js stdin
				decodeBytes, err := base64.StdEncoding.DecodeString(msgObj.Cmd)
				if err != nil {
					log.Error("websock cmd string base64 decoding failed")
				}
				sws.sendWebsocketInputCommandToSSHSessionStdinPipe(decodeBytes)
			}
		}
	}
}

//sendWebsocketInputCommandToSshSessionStdinPipe
func (sws *LogicSSHWsSession) sendWebsocketInputCommandToSSHSessionStdinPipe(cmdBytes []byte) {
	// log.Debug("cmdBytes:" + string(cmdBytes[:]))
	if _, err := sws.stdinPipe.Write(cmdBytes); err != nil {
		log.Error("ws cmd bytes write to ssh.stdin pipe failed")
	}
}

//SendComboOutput send
func (sws *LogicSSHWsSession) SendComboOutput() {
	wsConn := sws.wsConn

	//every 120ms write combine output bytes into websocket response
	tick := time.NewTicker(time.Millisecond * time.Duration(60))
	//for range time.Tick(120 * time.Millisecond){}
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			if sws.comboOutput == nil {
				return
			}
			bs := sws.comboOutput.Bytes()
			if len(bs) > 0 {
				// log.Debug("bs:" + string(bs[:]))
				err := wsConn.WriteMessage(websocket.TextMessage, bs)
				if err != nil {
					log.Error("ssh sending combo output to webSocket failed")
				}
				_, err = sws.logBuff.Write(bs)
				if err != nil {
					log.Error("combo output to log buffer failed")
				}
				sws.comboOutput.buffer.Reset()
			}

		case <-sws.exit:
			return
		}
	}
}

//Close 关闭
func (sws *LogicSSHWsSession) Close() {
	if sws.session != nil {
		sws.session.Close()
	}
	if sws.logBuff != nil {
		sws.logBuff = nil
	}
	if sws.comboOutput != nil {
		sws.comboOutput = nil
	}
}

func (sws *LogicSSHWsSession) LogString() string {
	return sws.logBuff.buffer.String()
}

//Wait 等待ssh session 错误， 发送exit信号
func (sws *LogicSSHWsSession) Wait() {
	if err := sws.session.Wait(); err != nil {
		log.Info(err.Error() + "---ssh session wait failed")
		// sws.exit <- true
		close(sws.exit)
	}
	log.Debug("remote command to exit.")
	close(sws.exit)
}
