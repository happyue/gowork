package log

import (
	"fmt"
	"os"

	logging "github.com/op/go-logging"
)

/**
 *日志级别
 */
const (
	critical int = iota
	err
	warning
	notice
	info
	debug
)

var (
	//logFormat 日志输出格式
	logFormat = []string{
		`%{shortfunc} ▶ %{level:.4s} %{message}`,
		`%{time:15:04:05.00} %{shortfunc} ▶ %{level:.4s} %{id:03x} %{message}`,
		`%{color}%{time:15:04:05.00} %{shortfunc} %{shortfile} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	}

	//LogLevelMap 日志级别与 string类型映射
	LogLevelMap = map[string]int{
		"CRITICAL": critical,
		"ERROR":    err,
		"WARNING":  warning,
		"NOTICE":   notice,
		"INFO":     info,
		"DEBUG":    debug,
	}
)

var log *logging.Logger

//Format
const (
	EasyFormat = iota
	MidFormat
	FullFormat
)

//Password hide
type Password string

//Redacted hide password
func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

//Option new log option
type Option struct {
	Level    string
	FilePath string
	ModeName string
	Format   int
}

//NewLog new log
func NewLog(logOption Option) {
	log = logging.MustGetLogger(logOption.ModeName)
	log.ExtraCalldepth = 2

	file, err := os.OpenFile(logOption.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	backend1 := logging.NewLogBackend(file, "", 1)
	backend2 := logging.NewLogBackend(os.Stderr, "", 1)

	format := logging.MustStringFormatter(logFormat[logOption.Format])
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.Level(LogLevelMap[logOption.Level]), "")
	logging.SetBackend(backend1Leveled, backend2Formatter)
}

//Infof info log
func Infof(fmtstr string, args ...interface{}) {
	log.Infof(fmtstr, args...)
	return
}

//Warningf info log
func Warningf(fmtstr string, args ...interface{}) {
	log.Warningf(fmtstr, args...)
	return
}

//Errorf info log
func Errorf(fmtstr string, args ...interface{}) {
	log.Errorf(fmtstr, args...)
	return
}

//Debugf info log
func Debugf(fmtstr string, args ...interface{}) {
	log.Debugf(fmtstr, args...)
	return
}

//Info info log
func Info(args ...interface{}) {
	log.Info(args...)
	return
}

//Warning info log
func Warning(args ...interface{}) {
	log.Warning(args...)
	return
}

//Error info log
func Error(args ...interface{}) {
	log.Error(args...)
	return
}

//Debug info log
func Debug(args ...interface{}) {
	log.Debug(args...)
	return
}
