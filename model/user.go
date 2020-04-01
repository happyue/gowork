package model

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"

	"gowork/config"
)

// User 用户
type User struct {
	Model

	Name      string `json:"name"`
	Pass      string `json:"-"`
	Role      string `json:"role"`      //角色
	AvatarURL string `json:"avatarURL"` //头像
}

// CheckPassword 验证密码是否正确
func (user User) CheckPassword(password string) bool {
	if password == "" || user.Pass == "" {
		return false
	}
	return user.EncryptPassword(password, user.Salt()) == user.Pass
}

// Salt 每个用户都有一个不同的盐
func (user User) Salt() string {
	var userSalt string
	if user.Pass == "" {
		userSalt = strconv.Itoa(int(time.Now().Unix()))
	} else {
		userSalt = user.Pass[0:10]
	}
	return userSalt
}

// EncryptPassword 给密码加密
func (user User) EncryptPassword(password, salt string) (hash string) {
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	hash = salt + password + config.ServerConfig.PassSalt
	hash = salt + fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	return
}

const (
	// UserRoleNormal 普通用户
	UserRoleNormal = "user"

	// UserRoleAdmin 管理员
	UserRoleAdmin = "admin"

	// UserRoleCrawler 爬虫，网站编辑或管理员登陆后台后，操作爬虫去抓取文章
	// 这时，生成的文章，其作者是爬虫账号。没有直接使用爬虫账号去登陆的情况.
	UserRoleCrawler = "crawler"
)

const (

	// MaxUserNameLen 用户名的最大长度
	MaxUserNameLen = 20

	// MinUserNameLen 用户名的最小长度
	MinUserNameLen = 4

	// MaxPassLen 密码的最大长度
	MaxPassLen = 20

	// MinPassLen 密码的最小长度
	MinPassLen = 6
)
