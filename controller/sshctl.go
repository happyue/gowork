package controller

import (
	"gowork/config"
	"gowork/controller/common"
	"gowork/model"
	"gowork/service"
	"gowork/utils"
	"gowork/utils/utsessions"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//SSHDeleMachine 删除machine配置
func SSHDeleMachine(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	session := utsessions.GetSession(c)

	type MachineData struct {
		ID uint `json:"id"`
	}
	var machineBind MachineData

	if err := c.ShouldBindWith(&machineBind, binding.JSON); err != nil {
		SendErrJSON("参数无效", c)
		return
	}

	var machine model.Machine

	machine.ID = machineBind.ID

	machine.UserID = session.User.ID

	err := service.Ssh.DeleMachine(machine)
	if err != nil {
		SendErrJSON("删除服务器配置错误！", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SHOWSUCCESS,
		"msg":   "成功删除服务器配置！",
	})
	return

}

//SSHUpdateMachine 修改machine配置
func SSHUpdateMachine(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	session := utsessions.GetSession(c)

	type MachineData struct {
		ID             uint   `json:"id"`
		ServerName     string `json:"servername"`
		ServerHost     string `json:"serverhost"`
		ServerUserName string `json:"serverusername"`
		ServerPassWord string `json:"serverpassword"`
		ServerPort     int    `json:"serverport"`
		ServerType     string `json:"servertype"`
	}
	var machineBind MachineData

	if err := c.ShouldBindWith(&machineBind, binding.JSON); err != nil {
		SendErrJSON("参数无效", c)
		return
	}

	var machine model.Machine
	machine.ID = machineBind.ID
	machine.Name = machineBind.ServerName
	machine.Host = machineBind.ServerHost
	machine.Port = machineBind.ServerPort
	machine.User = machineBind.ServerUserName
	machine.Password = machineBind.ServerPassWord
	machine.Type = machineBind.ServerType
	machine.UserID = session.User.ID

	// log.Debugf("machine [%#v]", machine)
	err := service.Ssh.UpdateMachine(machine)
	if err != nil {
		SendErrJSON("更新服务器配置错误！", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SHOWSUCCESS,
		"msg":   "修改配置成功！",
	})
	return

}

//GetSSHList 获取主机列表
func GetSSHList(c *gin.Context) {
	SendErrJSON := common.SendErrJSON

	type SSHData struct {
		Page  int    `json:"page"`
		Limit int    `json:"limit"`
		Sort  string `json:"sort"`
	}
	var sshBind SSHData

	if err := c.ShouldBindWith(&sshBind, binding.JSON); err != nil {
		SendErrJSON("参数无效", c)
		return
	}
	session := utsessions.GetSession(c)

	var machinelist []*model.Machine
	var err error
	var total int
	if sshBind.Sort == "-id" {
		machinelist, total, err = service.Ssh.GetMachineListDESC(sshBind.Limit, sshBind.Page, 20, session.User.ID)
	} else {
		machinelist, total, err = service.Ssh.GetMachineListASC(sshBind.Limit, sshBind.Page, 20, session.User.ID)
	}

	if err != nil {
		SendErrJSON("获取服务器列表错误", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data": gin.H{
			"items": machinelist,
			"total": total,
			"addr":  model.ServerIPInfo.IntranetIP + ":" + strconv.Itoa(config.ServerConfig.Port),
		},
	})
	return
}

//SSHAddMachine 新增主机
func SSHAddMachine(c *gin.Context) {
	SendErrJSON := common.SendErrJSON

	type MachineData struct {
		ServerName     string `json:"servername"`
		ServerHost     string `json:"serverhost"`
		ServerUserName string `json:"serverusername"`
		ServerPassWord string `json:"serverpassword"`
		ServerPort     int    `json:"serverport"`
		ServerType     string `json:"servertype"`
	}
	var machineBind MachineData

	if err := c.ShouldBindWith(&machineBind, binding.JSON); err != nil {
		SendErrJSON("参数无效", c)
		return
	}
	// log.Debugf("machineBind [%#v]", machineBind)
	session := utsessions.GetSession(c)

	var machine model.Machine
	machine.Name = machineBind.ServerName
	machine.Host = machineBind.ServerHost
	machine.Port = machineBind.ServerPort
	machine.User = machineBind.ServerUserName
	machine.Password = machineBind.ServerPassWord
	machine.Type = machineBind.ServerType
	machine.UserID = session.User.ID
	err := service.Ssh.AddMachine(machine)
	if err != nil {
		SendErrJSON("新增服务器错误！", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SHOWSUCCESS,
		"msg":   "成功新增SSH服务器！",
	})
	return
}

//NewSSHShell 新开ssh连接
func NewSSHShell(c *gin.Context) {
	SendErrJSON := common.SendErrJSON

	type SSHData struct {
		Cols int  `json:"cols"`
		Rows int  `json:"rows"`
		ID   uint `json:"id"`
	}
	var sshBind SSHData
	var err error
	// log.Debug("enter new ssh shell")
	wsConn, err := model.UpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		SendErrJSON(err.Error(), c)
		return
	}
	defer wsConn.Close()

	sshBind.Cols, err = strconv.Atoi(c.DefaultQuery("cols", "80"))
	if err != nil {
		SendErrJSON(err.Error(), c)
		return
	}
	sshBind.Rows, err = strconv.Atoi(c.DefaultQuery("rows", "40"))
	if err != nil {
		SendErrJSON(err.Error(), c)
		return
	}
	sshBind.ID, err = utils.ParseParamID(c)
	if err != nil {
		SendErrJSON(err.Error(), c)
		return
	}

	session := utsessions.GetSession(c)

	mc, err := service.Ssh.FindMachineByID(session.User.ID, sshBind.ID)
	if err != nil {
		SendErrJSON(err.Error(), c)
		return
	}

	client, err := mc.NewSSHClient()
	if err != nil {
		SendErrJSON(err.Error(), c)
		return
	}
	defer client.Close()

	startTime := time.Now()

	sws, err := model.NewLogicSSHWsSession(sshBind.Cols, sshBind.Rows, true, client, wsConn)
	if err != nil {
		SendErrJSON(err.Error(), c)
		return
	}
	defer sws.Close()

	sws.WaitGroup.Wrap(func() { sws.ReceiveWsMsg() })
	sws.WaitGroup.Wrap(func() { sws.SendComboOutput() })

	sws.WaitGroup.Wrap(func() { sws.Wait() })
	sws.WaitGroup.Wait()

	//write logs
	xtermLog := model.SSHLog{
		StartedAt: startTime,
		UserID:    session.User.ID,
		Log:       sws.LogString(),
		MachineID: sshBind.ID,
		// ClientIp:  cIp,
	}
	if err := service.Ssh.AddSSHLog(xtermLog); err != nil {
		SendErrJSON(err.Error(), c)
		return
	}
}
