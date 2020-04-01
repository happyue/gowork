package middleware

import (
	"errors"
	"fmt"

	"gowork/config"
	"gowork/controller/common"
	"gowork/log"
	"gowork/model"
	"gowork/service"
	"gowork/utils/utsessions"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) (model.User, error) {
	var user model.User
	// tokenString, cookieErr := c.Cookie("token")
	tokenString := c.Request.Header.Get("Token")
	if tokenString == "" {
		tokenString = c.DefaultQuery("Token", "")
	}
	// log.Debug("token:" + tokenString)
	// if cookieErr != nil {
	if tokenString == "" {
		return user, errors.New("未登录")
	}
	// log.Debug("getUser...")
	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.ServerConfig.TokenSecret), nil
	})

	if tokenErr != nil {
		return user, errors.New("未登录")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["id"].(float64))
		session := utsessions.GetSession(c)
		if nil != session.User {
			// log.Debugf("sid=%d,userid=%d", session.User.ID, userID)
			if int(session.User.ID) == userID {
				getUser := service.UserSrv.GetUserByUserID(session.User.ID)
				if getUser == nil {
					// if err := modeldb.DB.Where("`id` = ?", session.User.ID).Find(&user).Error; err != nil {
					session := sessions.Default(c)
					session.Options(sessions.Options{
						Path:   "/",
						MaxAge: -1,
					})
					session.Clear()
					if err := session.Save(); nil != err {
						log.Errorf("get user saves session failed: " + err.Error())
					}
					return user, errors.New("未登录")
				}
				user = *getUser
				return user, nil
			}
		}
	}
	return user, errors.New("未登录")
}

// SetContextUser 给 context 设置 user
func SetContextUser(c *gin.Context) {
	var user model.User
	var err error
	if user, err = getUser(c); err != nil {
		c.Set("user", nil)
		c.Next()
		return
	}
	c.Set("user", user)
	c.Next()
}

// SigninRequired 必须是登录用户
func SigninRequired(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	var user model.User
	var err error
	// log.Debug("enter sr!!!!")
	if user, err = getUser(c); err != nil {
		SendErrJSON("未登录", model.ErrorCode.LoginTimeout, c)
		return
	}
	c.Set("user", user)
	c.Next()
}

// // EditorRequired 必须是网站编辑
// func EditorRequired(c *gin.Context) {
// 	SendErrJSON := common.SendErrJSON
// 	var user model.User
// 	var err error
// 	if user, err = getUser(c); err != nil {
// 		SendErrJSON("未登录", model.ErrorCode.LoginTimeout, c)
// 		return
// 	}
// 	if user.Role == model.UserRoleEditor || user.Role == model.UserRoleAdmin || user.Role == model.UserRoleCrawler || user.Role == model.UserRoleSuperAdmin {
// 		c.Set("user", user)
// 		c.Next()
// 	} else {
// 		SendErrJSON("没有权限", c)
// 	}
// }

// AdminRequired 必须是管理员
func AdminRequired(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	var user model.User
	var err error
	if user, err = getUser(c); err != nil {
		SendErrJSON("未登录", model.ErrorCode.LoginTimeout, c)
		return
	}
	if user.Role == model.UserRoleAdmin || user.Role == model.UserRoleCrawler {
		c.Set("user", user)
		c.Next()
	} else {
		SendErrJSON("没有权限", c)
	}
}
