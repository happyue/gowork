package main

import (
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"

	"gowork/config"
	"gowork/controller"
	"gowork/log"
	"gowork/model"
)

func init() {

	log.Infof("gin.Version: %s", gin.Version)
	if config.ServerConfig.Env != model.DevelopmentMode {
		gin.SetMode(gin.ReleaseMode)
		// Disable Console Color, you don't need console color when writing the logs to file.
		gin.DisableConsoleColor()
		// Logging to a file.
		logFile, err := os.OpenFile(config.ServerConfig.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Error(err.Error())
			os.Exit(-1)
		}
		gin.DefaultWriter = io.MultiWriter(logFile)
	}
}

func main() {
	router := controller.MapRoutes()
	server := &http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(config.ServerConfig.Port),
		Handler: router,
	}

	// cron.New().Start()

	handleSignal(server)

	log.Debugf("gowork server (v%s) is running [%s]", model.Version, strconv.Itoa(config.ServerConfig.Port))
	server.ListenAndServe()
}

// handleSignal handles system signal for graceful shutdown.
func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		log.Infof("got signal [%s], exiting gowork now", s)
		if err := server.Close(); nil != err {
			log.Errorf("server close failed: " + err.Error())
		}

		log.Infof("gowork exited")
		os.Exit(0)
	}()
}
