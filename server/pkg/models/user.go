package models

import (
	"crypto/md5"
	"fmt"
	"github.com/jameskeane/bcrypt"
	"simple-blog/pkg/global"
	"strings"
	"time"
)

type User struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	DisplayName   string    `json:"display_name"`
	Password      string    `json:"-"`
	Email         string    `json:"-"`
	Status        int       `json:"-"`
	LastLoginTime time.Time `json:"-"`
	LastLoginIp   string    `json:"-"`
}

func (User) TableName() string { return global.Config.DbTablePrefix + "user" }

func CheckPass(loginPass string, savePass string) bool {
	sig := md5.Sum([]byte(global.Options["password_salt"] + loginPass))
	pass := strings.ToLower(strings.ToTitle(fmt.Sprintf("%x", sig)))
	match := bcrypt.Match(pass, savePass)
	return match
}

func GetUserByName(name string) (*User, error) {
	var user User
	has, err := global.Store.Where("name = ?", name).Get(&user)
	if has {
		return &user, err
	}
	return nil, err
}

func GetUserById(id int32) (*User, error) {
	var user User
	has, err := global.Store.Where("id = ?", id).Get(&user)
	if has {
		return &user, err
	}
	return nil, err
}

func UpdateUser(user User) {
	_, err := global.Store.Update(user, User{Id: user.Id})
	if err != nil {
		panic(err)
	}
}
