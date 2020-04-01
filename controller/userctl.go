package controller

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"gowork/config"
	"gowork/controller/common"

	"gowork/log"
	"gowork/model"
	"gowork/service"

	"gowork/utils"
	"gowork/utils/utsessions"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//GetRegStatus 获取注册状态
func GetRegStatus(c *gin.Context) {
	SendErrJSON := common.SendErrJSON

	regStatus, err := service.Cfg.GetGwcfg()
	if err != nil {
		SendErrJSON("获取注册状态错误", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  regStatus.RegStatus,
	})
	return
}

//SetRegStatus 设置注册状态
func SetRegStatus(c *gin.Context) {
	SendErrJSON := common.SendErrJSON

	type RegStatus struct {
		Status bool `json:"status"`
	}
	var regBind RegStatus

	if err := c.ShouldBindWith(&regBind, binding.JSON); err != nil {
		SendErrJSON("参数无效", c)
		return
	}

	var cfg model.Gwcfg

	cfg.ID = 1
	cfg.RegStatus = regBind.Status

	err := service.Cfg.SetRegStatus(cfg)
	if err != nil {
		SendErrJSON("设置注册状态错误", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SHOWSUCCESS,
		"msg":   "更改注册状态成功",
	})
	return
}

//GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	SendErrJSON := common.SendErrJSON

	session := utsessions.GetSession(c)
	if nil != session.User {
		if user := service.UserSrv.GetUserByUserID(session.User.ID); user == nil {
			session := sessions.Default(c)

			session.Options(sessions.Options{
				Path:   "/",
				MaxAge: -1,
			})

			session.Clear()
			if err := session.Save(); nil != err {
				log.Error("saves session failed: " + err.Error())
			}
		}
		roles := []string{session.User.Role} //为了前端的检测权限
		c.JSON(http.StatusOK, gin.H{
			"errNo": model.ErrorCode.SUCCESS,
			"msg":   "success",
			"data": gin.H{
				"name":   session.User.Name,
				"roles":  roles,
				"avatar": session.User.AvatarURL,
			},
		})
		return
	}
	SendErrJSON("账号不存在", c)
	return
}

// Signin 用户登录
func Signin(c *gin.Context) {
	SendErrJSON := common.SendErrJSON

	type UsernameLogin struct {
		Username string `json:"username" binding:"required,min=4,max=20"`
		Password string `json:"password" binding:"required,min=6,max=20"`
	}

	var usernameLogin UsernameLogin

	if err := c.ShouldBindWith(&usernameLogin, binding.JSON); err != nil {
		SendErrJSON("用户名或密码错误:"+err.Error(), c)
		return
	}

	username := usernameLogin.Username
	password := usernameLogin.Password

	// var user model.User
	user := &model.User{}

	if user = service.UserSrv.GetUserByUsername(username); user == nil {
		SendErrJSON("账号不存在", c)
		return
	}

	if user.CheckPassword(password) {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": user.ID,
		})
		tokenString, err := token.SignedString([]byte(config.ServerConfig.TokenSecret))
		if err != nil {
			log.Errorf(err.Error())
			SendErrJSON("内部错误", c)
			return
		}

		session := &utsessions.SessionData{
			Version: config.ServerConfig.Version,
			User:    user,
		}
		if err := session.Save(c); nil != err {
			SendErrJSON("saves session failed: "+err.Error(), c)
		}

		c.JSON(http.StatusOK, gin.H{
			"errNo": model.ErrorCode.SHOWSUCCESS,
			"msg":   "登录成功",
			"data": gin.H{
				"token": tokenString,
				"user":  user,
			},
		})
		return
	}
	SendErrJSON("账号或密码错误", c)
}

// Signup 用户注册
func Signup(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	type UserReqData struct {
		Name     string `json:"name" binding:"required,min=4,max=20"`
		Password string `json:"password" binding:"required,min=6,max=20"`
	}

	var userData UserReqData
	if err := c.ShouldBindWith(&userData, binding.JSON); err != nil {
		SendErrJSON("用户名或密码无效", c)
		return
	}

	userData.Name = utils.AvoidXSS(userData.Name)
	userData.Name = strings.TrimSpace(userData.Name)

	regStatus, err := service.Cfg.GetGwcfg()
	if err != nil {
		SendErrJSON("获取注册状态错误", c)
		return
	}
	if regStatus.RegStatus == false {
		SendErrJSON("注册功能已关闭，请联系管理员", c)
		return
	}

	if user := service.UserSrv.GetUserByUsername(userData.Name); user != nil {
		SendErrJSON("用户名 "+user.Name+" 已被注册", c)
		return
	}

	var newUser model.User
	newUser.Name = userData.Name
	newUser.Pass = newUser.EncryptPassword(userData.Password, newUser.Salt())
	newUser.Role = model.UserRoleNormal
	newUser.AvatarURL = model.AvatarURL

	if err := service.UserSrv.AddUser(&newUser); err != nil {
		SendErrJSON("服务器错误", c)
		return
	}

	session := &utsessions.SessionData{
		Version: config.ServerConfig.Version,
		User:    &newUser,
	}
	// log.Debugf("configurations [%#v]", session)
	if err := session.Save(c); nil != err {
		SendErrJSON("saves session failed: "+err.Error(), c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SHOWSUCCESS,
		"msg":   "注册成功",
	})
}

// Signout 退出登录
func Signout(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})
	session.Clear()
	if err := session.Save(); nil != err {
		log.Error("saves session failed: " + err.Error())
		SendErrJSON("saves session failed: "+err.Error(), c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  gin.H{},
	})

}

// UpdatePassword 更新用户密码
func UpdatePassword(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	type userReqData struct {
		Password string `json:"password" binding:"required,min=6,max=20"`
		NewPwd   string `json:"newPwd" binding:"required,min=6,max=20"`
	}
	var userData userReqData
	if err := c.ShouldBindWith(&userData, binding.JSON); err != nil {
		SendErrJSON("参数无效", c)
		return
	}

	userInter, _ := c.Get("user")
	user := userInter.(model.User)

	getUser := service.UserSrv.GetUserByUserID(user.ID)
	if getUser == nil {
		SendErrJSON("error", c)
		return
	}
	user = *getUser

	if user.CheckPassword(userData.Password) {
		user.Pass = user.EncryptPassword(userData.NewPwd, user.Salt())
		if err := service.UserSrv.SaveUser(&user); err != nil {
			SendErrJSON("原密码不正确", c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"errNo": model.ErrorCode.SHOWSUCCESS,
			"msg":   "密码修改成功！",
			"data":  gin.H{},
		})
	} else {
		SendErrJSON("原密码错误", c)
		return
	}
}

// UpdateAvatar 更新用户头像
func UpdateAvatar(c *gin.Context) {
	SendErrJSON := common.SendErrJSON

	file, header, err := c.Request.FormFile("avatar")

	if err != nil {
		SendErrJSON("上传头像失败", c)
		return
	}

	filename := strconv.FormatInt(time.Now().Unix(), 10) + "_" + header.Filename
	avatarURL := "/static/img/" + filename
	realname := config.PathPlatformStatic + "/img/" + filename

	log.Debug("avatarURL:" + avatarURL)
	log.Debug("realname:" + realname)

	//写入文件
	out, err := os.Create(realname)
	if err != nil {
		SendErrJSON("上传头像失败", c)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		SendErrJSON("上传头像失败", c)
		return
	}

	userInter, _ := c.Get("user")
	user := userInter.(model.User)

	if err := service.UserSrv.UpdateAvatar(user.ID, avatarURL); err != nil {
		SendErrJSON("上传头像失败", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": model.ErrorCode.SHOWSUCCESS,
		"msg":   "修改头像成功",
		"data": gin.H{
			"avatarURL": avatarURL,
		},
	})
}

// // AllList 查询用户列表，只有管理员才能调此接口
// func AllList(c *gin.Context) {
// 	SendErrJSON := common.SendErrJSON
// 	role := c.Query("role")
// 	allUserRole := []string{
// 		model.UserRoleNormal,
// 		model.UserRoleEditor,
// 		model.UserRoleAdmin,
// 		model.UserRoleCrawler,
// 		model.UserRoleSuperAdmin,
// 	}
// 	foundRole := false
// 	for _, r := range allUserRole {
// 		if r == role {
// 			foundRole = true
// 			break
// 		}
// 	}

// 	var startTime string
// 	var endTime string

// 	if startAt, err := strconv.Atoi(c.Query("startAt")); err != nil {
// 		startTime = time.Unix(0, 0).Format("2006-01-02 15:04:05")
// 	} else {
// 		startTime = time.Unix(int64(startAt/1000), 0).Format("2006-01-02 15:04:05")
// 	}

// 	if endAt, err := strconv.Atoi(c.Query("endAt")); err != nil {
// 		endTime = time.Now().Format("2006-01-02 15:04:05")
// 	} else {
// 		endTime = time.Unix(int64(endAt/1000), 0).Format("2006-01-02 15:04:05")
// 	}

// 	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
// 	if pageNoErr != nil {
// 		pageNo = 1
// 	}
// 	if pageNo < 1 {
// 		pageNo = 1
// 	}

// 	offset := (pageNo - 1) * model.PageSize
// 	pageSize := model.PageSize

// 	var users []model.User
// 	var totalCount int
// 	if foundRole {
// 		if err := modeldb.DB.Model(&model.User{}).Where("created_at >= ? AND created_at < ? AND role = ?", startTime, endTime, role).
// 			Count(&totalCount).Error; err != nil {
// 			fmt.Println(err.Error())
// 			SendErrJSON("error", c)
// 			return
// 		}
// 		if err := modeldb.DB.Where("created_at >= ? AND created_at < ? AND role = ?", startTime, endTime, role).
// 			Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
// 			fmt.Println(err.Error())
// 			SendErrJSON("error", c)
// 			return
// 		}
// 	} else {
// 		if err := modeldb.DB.Model(&model.User{}).Where("created_at >= ? AND created_at < ?", startTime, endTime).
// 			Count(&totalCount).Error; err != nil {
// 			fmt.Println(err.Error())
// 			SendErrJSON("error", c)
// 			return
// 		}
// 		if err := modeldb.DB.Where("created_at >= ? AND created_at < ?", startTime, endTime).Order("created_at DESC").Offset(offset).
// 			Limit(pageSize).Find(&users).Error; err != nil {
// 			fmt.Println(err.Error())
// 			SendErrJSON("error", c)
// 			return
// 		}
// 	}
// 	var results []interface{}
// 	for i := 0; i < len(users); i++ {
// 		results = append(results, gin.H{
// 			"id":          users[i].ID,
// 			"name":        users[i].Name,
// 			"email":       users[i].Email,
// 			"role":        users[i].Role,
// 			"status":      users[i].Status,
// 			"createdAt":   users[i].CreatedAt,
// 			"activatedAt": users[i].ActivatedAt,
// 		})
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"errNo": model.ErrorCode.SUCCESS,
// 		"msg":   "success",
// 		"data": gin.H{
// 			"users":      results,
// 			"pageNo":     pageNo,
// 			"pageSize":   pageSize,
// 			"totalCount": totalCount,
// 		},
// 	})
// }
