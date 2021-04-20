package service

import (
	"errors"
	"fmt"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/viper"
	"qn/server/global"
	"qn/server/middlewares"
	"qn/server/models"
	"qn/server/util"
)

type Status int

func GenerateToken() string {
	bucket := viper.GetString("qiniu.bucket")
	putPolicy := storage.PutPolicy{Scope: bucket}
	upToken := putPolicy.UploadToken(global.MAC)
	fmt.Println(upToken)
	return upToken
}

type Login models.Login

//返回所有用户的登录信息
func QueryLogin() []models.Login {
	var logins []models.Login
	result := global.QN_DB.Find(&logins)
	if result.Error != nil {
		return nil
	}
	return logins
}

func TrueLogin(username, password string) (string, error) {
	var login models.Login
	err := global.QN_DB.Where("username = ?", username).Find(&login).Error
	if err != nil {
		return "", errors.New("用户不存在")
	}
	if username == login.Username && util.HashDecrypt(login.Password, password) {
		tokenString, _ := middlewares.GenToken(login.Username, login.Password)
		return tokenString, nil
	} else {
		return "", errors.New("密码不正确")

	}

}

//添加用户
func PostLogin(username, password string) Status {
	hashPwd := util.HashEncryption(password)
	login := &models.Login{
		Username: username,
		Password: hashPwd,
	}

	result := global.QN_DB.Create(&login)
	if result.Error != nil {
		return 401
	} else {
		return 200
	}
}
