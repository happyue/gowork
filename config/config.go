package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"

	"gowork/log"
	"gowork/utils"

	"github.com/jinzhu/gorm"
)

var jsonData map[string]interface{}

//CodeData 城市气象代码
var CodeData map[string]string

func initJSON() {
	bytes, err := ioutil.ReadFile("./gowork.json")
	if err != nil {
		fmt.Println("loads configuration file [gowork.json] failed: " + err.Error())
		os.Exit(-1)
	}

	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		fmt.Println("parses [gowork.json] failed: " + err.Error())
		os.Exit(-1)
	}
}

func initCityCode() {
	bytes, err := ioutil.ReadFile("./addrcode.json")
	if err != nil {
		fmt.Println("loads configuration file [addrcode.json] failed: " + err.Error())
		os.Exit(-1)
	}

	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &CodeData); err != nil {
		fmt.Println("parses [addrcode.json] failed: " + err.Error())
		os.Exit(-1)
	}
}

type serverConfig struct {
	Version           string
	APIPoweredBy      string
	SiteName          string
	Host              string
	ImgHost           string
	Env               string
	LogDir            string
	BaseDir           string
	HomePath          string
	LogFile           string
	APIPrefix         string
	UploadImgDir      string
	Port              int
	TokenSecret       string
	TokenMaxAge       int
	PassSalt          string
	OpenRegister      bool
	ShowSQL           bool
	LogLevel          string
	SQLite            string
	MaxIdleConns      int
	MaxOpenConns      int
	SessionSecret     string
	SessionMaxAge     int
	TablePrefix       string
	QiniuAccessKey    string
	QiniuSecretKey    string
	QiniuBucket       string
	QiniuExDomainName string
}

var (
	// ServerConfig 服务器相关配置
	ServerConfig serverConfig

	//PathPlatform 平台index路径
	PathPlatform string

	//PathPlatformStatic 平台static路径
	PathPlatformStatic string

	//PathPlatformFavicon 平台favicon路径
	PathPlatformFavicon string
)

func initServer() {
	utils.SetStructByJSON(&ServerConfig, jsonData["go"].(map[string]interface{}))
	sep := string(os.PathSeparator)
	execPath, _ := os.Getwd()
	length := utf8.RuneCountInString(execPath)
	lastChar := execPath[length-1:]
	if lastChar != sep {
		execPath = execPath + sep
	}
	if ServerConfig.UploadImgDir == "" {
		pathArr := []string{"website", "static", "upload", "img"}
		uploadImgDir := execPath + strings.Join(pathArr, sep)
		ServerConfig.UploadImgDir = uploadImgDir
	}

	ymdStr := utils.GetTodayYMD("-")

	ServerConfig.BaseDir = execPath

	if ServerConfig.LogDir == "" {
		ServerConfig.LogDir = execPath
	} else {
		length := utf8.RuneCountInString(ServerConfig.LogDir)
		lastChar := ServerConfig.LogDir[length-1:]
		if lastChar != sep {
			ServerConfig.LogDir = ServerConfig.LogDir + sep
		}
	}
	ServerConfig.LogFile = ServerConfig.LogDir + "gowork" + ymdStr + ".log"

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return ServerConfig.TablePrefix + "_" + defaultTableName
	}

	home, err := utils.UserHome()
	if nil != err {
		fmt.Println("can't find user home directory: " + err.Error())
		os.Exit(-1)
	}
	fmt.Println("${home} : [" + home + "]")
	ServerConfig.HomePath = home

	ServerConfig.SQLite = ServerConfig.SQLite + sep + "gowork.db"
	ServerConfig.SQLite = strings.Replace(ServerConfig.SQLite, "${home}", home, 1)

	log.NewLog(log.Option{
		Level:    ServerConfig.LogLevel,
		FilePath: ServerConfig.LogFile,
		ModeName: ServerConfig.TablePrefix,
		Format:   log.FullFormat,
	})

	PathPlatform = ServerConfig.BaseDir + "dist" + sep + "index.html"
	PathPlatformStatic = ServerConfig.BaseDir + "dist" + sep + "static"
	PathPlatformFavicon = ServerConfig.BaseDir + "dist" + sep + "favicon.ico"
	log.Debugf("config.PathPlatform:%s", PathPlatform)

	log.Debugf("configurations [%#v]", ServerConfig)
}

func init() {
	initJSON()
	initCityCode()
	initServer()
}
