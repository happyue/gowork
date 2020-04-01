package service

import (
	"gowork/config"
	"gowork/log"
	"gowork/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite
)

// DB 数据库连接
var db *gorm.DB

// Models represents all models..
var Models = []interface{}{
	&model.User{}, &model.Todo{}, &model.QiniuStorage{}, &model.Gwcfg{},
	&model.Machine{}, &model.SSHLog{},
}

// ConnectDB connects to the database.
func initDB() {
	var err error
	log.Debug("sqlite3:" + config.ServerConfig.SQLite)
	db, err = gorm.Open("sqlite3", config.ServerConfig.SQLite)
	if err != nil {
		log.Error("opens database failed: " + err.Error())
	}
	if err = db.AutoMigrate(Models...).Error; nil != err {
		log.Error("auto migrate tables failed: " + err.Error())
	}

	if config.ServerConfig.Env == model.DevelopmentMode {
		db.LogMode(config.ServerConfig.ShowSQL)
	}
	db.DB().SetMaxIdleConns(config.ServerConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.ServerConfig.MaxOpenConns)
}

// DisconnectDB disconnects from the database.
func DisconnectDB() {
	if err := db.Close(); nil != err {
		log.Error("Disconnect from database failed: " + err.Error())
	}
}

func initAdminUser() {

	var user model.User
	if err := db.Where("name = ?", "admin").Find(&user).Error; err == nil {
		if user.Name == "admin" {
			log.Debug("存在管理员 " + user.Name)
			return
		}
	}

	var newUser model.User

	newUser.Name = "admin"
	newUser.Pass = newUser.EncryptPassword("admin888", newUser.Salt())
	newUser.AvatarURL = model.AvatarURL
	newUser.Role = "admin"

	if err := db.Create(&newUser).Error; err != nil {
		log.Error("init Admin User error")
	}

	var cfg model.Gwcfg
	cfg.RegStatus = config.ServerConfig.OpenRegister
	if err := db.Create(&cfg).Error; err != nil {
		log.Error("init gowork config error")
	}
}

func init() {
	initDB()
	initAdminUser()
	go model.GetIPInfo()
}
