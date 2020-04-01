package controller

import (
	"gowork/config"
	"gowork/log"
	"gowork/middleware"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// MapRoutes 路由函数
func MapRoutes() *gin.Engine {
	ret := gin.New()
	ret.Use(gin.Recovery())
	ret.Use(gin.Logger())
	store := cookie.NewStore([]byte(config.ServerConfig.SessionSecret))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   config.ServerConfig.SessionMaxAge,
		Secure:   strings.HasPrefix(config.ServerConfig.Host+":"+strconv.Itoa(config.ServerConfig.Port), "https"),
		HttpOnly: true,
	})
	ret.Use(sessions.Sessions("gowork", store))

	// 加载静态文件路径
	ret.Static("/static", config.PathPlatformStatic)
	ret.Static("/favicon.ico", config.PathPlatformFavicon)

	// 代理news通信，新闻接口数据，直接从平台服务器获取
	microGroup := ret.Group("/platform")
	microGroup.Any("/*action", WithHeader)

	indexGroup := ret.Group("")
	indexGroup.GET("", showIndexAction)

	apiPrefix := config.ServerConfig.APIPrefix
	api := ret.Group(apiPrefix, middleware.RefreshTokenCookie)
	{
		api.POST("/login/login", Signin)
		api.POST("/user/signup", Signup)
		api.GET("/user/info", GetUserInfo, middleware.SigninRequired)
		api.POST("/login/logout", middleware.SigninRequired, Signout)
		api.POST("/user/updatepassword", middleware.SigninRequired, UpdatePassword)
		api.GET("/computerinfo", middleware.SigninRequired, GetComputerInfo)
		api.GET("/weatherinfo", middleware.SigninRequired, GetWeatherInfo)
		api.GET("/getTodoList", middleware.SigninRequired, GetTodoList)
		api.POST("/addTodo", middleware.SigninRequired, AddTodo)
		api.POST("/deleTodo", middleware.SigninRequired, DeleTodo)
		api.POST("/updateTodo", middleware.SigninRequired, UpdateTodo)
		api.POST("/uploadAvatar", middleware.SigninRequired, UpdateAvatar)
		api.GET("/qiniu/upload/token", middleware.SigninRequired, GetQiniuUploadToken)
		api.GET("/qiniu/getMediaInfo", middleware.SigninRequired, GetQiniuMedia)
		api.POST("/qiniu/addMediaInfo", middleware.SigninRequired, AddQiniuMedia)
		api.POST("/qiniu/deleMediaInfo", middleware.SigninRequired, DeleQiniuMedia)
		api.GET("/getRegStatus", middleware.SigninRequired, GetRegStatus)
		api.POST("/setRegStatus", middleware.SigninRequired, SetRegStatus)

		//ssh
		api.GET("/newSsh/:id", middleware.SigninRequired, NewSSHShell)
		api.POST("/sshAddMachine", middleware.SigninRequired, SSHAddMachine)
		api.POST("/getSSHList", middleware.SigninRequired, GetSSHList)
		api.POST("/sshUpdateMachine", middleware.SigninRequired, SSHUpdateMachine)
		api.POST("/sshDeleMachine", middleware.SigninRequired, SSHDeleMachine)

	}

	return ret
}

//新闻接口数据，直接从服务器获取
var simpleHostProxy = httputil.ReverseProxy{
	Director: func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = "47.100.54.157:80"
		req.Host = "47.100.54.157:80"
	},
}

//WithHeader 可改写头
func WithHeader(ctx *gin.Context) {
	// ctx.Request.Header.Add("requester-uid", "id")
	log.Debugf("ctx.Request.url:" + ctx.Request.URL.Path)
	// ctx.Request.URL.Path = strings.Replace(ctx.Request.URL.Path, model.Conf.AxiosBaseURL, "", 1)
	// ctx.Request.URL.Path = "/"
	// log.Debugf("ctx.Request.url:" + ctx.Request.URL.Path)
	simpleHostProxy.ServeHTTP(ctx.Writer, ctx.Request)
}
