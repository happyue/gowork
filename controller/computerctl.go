package controller

import (
	"fmt"
	"gowork/config"
	"gowork/controller/common"
	"gowork/model"
	"gowork/utils"
	"net/http"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

//GetComputerInfo 获取计算机详细信息
func GetComputerInfo(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	computerInfo := &model.ComputerInfo{}
	computerInfo.CpuInfo = &model.CpuInfo{}
	var err error

	if computerInfo.InfoStat, err = host.Info(); err != nil {
		SendErrJSON("获取host信息出错！", c)
		return
	}
	computerInfo.CpuInfo.Counts = runtime.NumCPU()
	// if computerInfo.CpuInfo.Counts, err = cpu.Counts(false); err != nil {
	// 	SendErrJSON("获取cpu信息出错！", c)
	// 	return
	// }

	if computerInfo.CpuInfo.Percent, err = cpu.Percent(0, false); err != nil {
		SendErrJSON("获取cpu信息出错！", c)
		return
	}
	computerInfo.CpuInfo.Percent[0], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", computerInfo.CpuInfo.Percent[0]), 64)

	if computerInfo.VirtualMemoryStat, err = mem.VirtualMemory(); err != nil {
		SendErrJSON("获取memory信息出错！", c)
		return
	}
	computerInfo.VirtualMemoryStat.Total = computerInfo.VirtualMemoryStat.Total / uint64(GB)
	computerInfo.VirtualMemoryStat.Available = computerInfo.VirtualMemoryStat.Available / uint64(GB)
	computerInfo.VirtualMemoryStat.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", computerInfo.VirtualMemoryStat.UsedPercent), 64)

	if computerInfo.UsageStat, err = disk.Usage("/"); err != nil {
		SendErrJSON("获取disk信息出错！", c)
		return
	}
	computerInfo.UsageStat.Total = computerInfo.UsageStat.Total / uint64(GB)
	computerInfo.UsageStat.Free = computerInfo.UsageStat.Free / uint64(GB)

	// log.Debugf("CpuInfo [%#v]", computerInfo.CpuInfo)
	// log.Debugf("InfoStat [%#v]", computerInfo.InfoStat)
	// log.Debugf("VirtualMemoryStat [%#v]", computerInfo.VirtualMemoryStat)
	// log.Debugf("UsageStat [%#v]", computerInfo.UsageStat)

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  computerInfo,
	})
	return
}

//GetWeatherInfo 获取15天天气信息
func GetWeatherInfo(c *gin.Context) {
	SendErrJSON := common.SendErrJSON

	// log.Debug("ip : " + ipinfo.ExternalIP)
	// log.Debug("city : " + ipinfo.PconlineInfo.City)
	// log.Debugf("ipinfo [%#v]", ipinfo)
	// log.Debugf("PconlineInfo [%#v]", ipinfo.PconlineInfo)
	cityCode := config.CodeData[model.ServerIPInfo.PconlineInfo.City]
	// log.Debug("cityCode : " + cityCode)
	// "佳木斯市":"101050401",
	// weather, err := utils.GetWeather("101050401")
	if cityCode == "" {
		cityCode = "101020100"
	}
	weather, err := utils.GetWeather(cityCode)
	if err != nil {
		SendErrJSON("获取城市天气出错，请刷新重试！", c)
		return
	}

	// log.Debugf("weather [%#v]", weather)
	eweather := utils.TranWeatherToEWeather(weather)

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  eweather,
	})
	return
}

// func GetComputerInfo(c *gin.Context) {
// 	SendErrJSON := common.SendErrJSON
// 	computerInfo := &model.ComputerInfo{}

// 	computerInfo.CurrentOS = runtime.GOOS
// 	computerInfo.NumCPU = runtime.NumCPU()
// 	computerInfo.Arch = runtime.GOARCH
// 	hostname, err := os.Hostname()
// 	if err != nil {
// 		SendErrJSON("操作系统不支持", c)
// 		return
// 	}
// 	computerInfo.HostName = hostname

// 	switch computerInfo.CurrentOS {
// 	case "darwin":
// 		diskStatus := DiskUsage("/")
// 		computerInfo.DiskStatus = &diskStatus

// 	case "linux":
// 		diskStatus := DiskUsage("/home")
// 		computerInfo.DiskStatus = &diskStatus
// 	case "windows":
// 		//		diskStatus := DiskUsageWindows("C:")
// 		//		computerInfo.DiskStatus = &diskStatus
// 		SendErrJSON("操作系统不支持", c)
// 	default:
// 		SendErrJSON("操作系统不支持", c)
// 		return
// 	}

// 	logger.Debugf("DiskStatus [%#v]", computerInfo.DiskStatus)
// 	logger.Debugf("computerInfo [%#v]", computerInfo)

// 	c.JSON(http.StatusOK, gin.H{
// 		"errNo": model.ErrorCode.SUCCESS,
// 		"msg":   "success",
// 		"data":  computerInfo,
// 	})
// 	return
// }
