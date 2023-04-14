package service

import (
	"blog/dao"
	"blog/models"
)

func Login(userName, password string) (models.LoginRes, error) {
	userInfo := dao.GetUser(userName, password)
	var LoginRes models.LoginRes
	LoginRes.UserInfo = userInfo
	LoginRes.Token = "token"
	return LoginRes
}