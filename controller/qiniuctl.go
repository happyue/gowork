package controller

import (
	"math"
	"net/http"

	"gowork/config"
	"gowork/model"
	"gowork/service"
	"gowork/utils/utsessions"

	"gowork/controller/common"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

//GetQiniuUploadToken 获取七牛对象存储的token
func GetQiniuUploadToken(c *gin.Context) {
	// SendErrJSON := common.SendErrJSON

	putPolicy := storage.PutPolicy{
		Scope: config.ServerConfig.QiniuBucket,
	}
	putPolicy.Expires = 7200 //2小时有效期
	mac := qbox.NewMac(config.ServerConfig.QiniuAccessKey, config.ServerConfig.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  upToken,
	})
	return
}

//GetQiniuMedia 获取七牛对象存储的媒体文件列表
func GetQiniuMedia(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	session := utsessions.GetSession(c)

	var filelist []*model.QiniuStorage
	var err error
	filelist, err = service.Qiniu.GetQiniuMediaList(math.MaxInt64, session.User.ID)
	if err != nil {
		SendErrJSON("获取媒体文件列表错误", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  filelist,
	})
	return
}

//AddQiniuMedia 增加媒体文件
func AddQiniuMedia(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	session := utsessions.GetSession(c)

	type File struct {
		UID  uint   `json:"uid"`
		Name string `json:"name"`
	}
	var fileBind File

	if err := c.ShouldBindWith(&fileBind, binding.JSON); err != nil {
		SendErrJSON("服务器错误，解析media失败！", c)
		return
	}

	var file model.QiniuStorage

	file.UID = fileBind.UID
	file.Name = fileBind.Name
	file.URL = config.ServerConfig.QiniuExDomainName + "/" + fileBind.Name
	file.UserID = session.User.ID

	err := service.Qiniu.AddQiniuMedia(file)
	if err != nil {
		SendErrJSON("新增media错误！", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SHOWSUCCESS,
		"msg":   "上传成功！",
	})
	return

}

//DeleQiniuMedia 删除
func DeleQiniuMedia(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	session := utsessions.GetSession(c)

	type File struct {
		UID  uint   `json:"uid"`
		Name string `json:"name"`
	}
	var fileBind File

	if err := c.ShouldBindWith(&fileBind, binding.JSON); err != nil {
		SendErrJSON("参数无效", c)
		return
	}

	var file model.QiniuStorage

	file.UID = fileBind.UID
	file.Name = fileBind.Name
	file.UserID = session.User.ID

	err := service.Qiniu.DeleQiniuMedia(file)
	if err != nil {
		SendErrJSON("删除media错误！", c)
		return
	}

	mac := qbox.NewMac(config.ServerConfig.QiniuAccessKey, config.ServerConfig.QiniuSecretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}

	// log.Debug(file.Name)
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err = bucketManager.Delete(config.ServerConfig.QiniuBucket, file.Name)
	if err != nil {
		SendErrJSON("七牛删除media错误！", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SHOWSUCCESS,
		"msg":   "删除成功！",
	})
	return

}
