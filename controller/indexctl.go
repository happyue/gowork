package controller

import (
	"net/http"
	"path/filepath"
	"text/template"

	"gowork/config"
	"gowork/log"

	"github.com/gin-gonic/gin"
)

// 平台
func showIndexAction(c *gin.Context) {
	t, err := template.ParseFiles(filepath.ToSlash(config.PathPlatform))
	if nil != err {
		log.Errorf("load platform page failed: " + err.Error())
		c.String(http.StatusNotFound, "load  platform page failed")

		return
	}
	t.Execute(c.Writer, nil)
}

func showGocubePlatformIndexAction(c *gin.Context) {
	log.Debugf("URLpath:" + c.Request.URL.Path)
	t, err := template.ParseFiles(filepath.ToSlash(filepath.Join(config.PathPlatform, c.Request.URL.Path+"/index.html")))
	if nil != err {
		log.Errorf("load  page failed: " + err.Error())
		c.String(http.StatusNotFound, "load  page failed")

		return
	}
	t.Execute(c.Writer, nil)
}
