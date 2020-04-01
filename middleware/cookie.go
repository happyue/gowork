package middleware

import (
	"gowork/config"
	"gowork/controller/common"
	"gowork/utils/utsessions"

	"github.com/gin-gonic/gin"
	// "gowork/model"
)

// RefreshTokenCookie 刷新过期时间
func RefreshTokenCookie(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	// tokenString, err := c.Cookie("token")
	// log.Debug("enter refresh token")
	tokenString := c.Request.Header.Get("Token")

	// fmt.Println(err)
	// if tokenString != "" && err == nil {
	if tokenString != "" {
		// c.SetCookie("token", tokenString, config.ServerConfig.TokenMaxAge, "/", "", true, true)
		if user, err := getUser(c); err == nil {
			session := &utsessions.SessionData{
				Version: config.ServerConfig.Version,
				User:    &user,
			}
			if err := session.Save(c); nil != err {
				SendErrJSON("saves session failed: "+err.Error(), c)
			}
		}
	}
	c.Next()
}
