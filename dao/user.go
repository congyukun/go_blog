package dao

import (
	"blog/models"
	"log"
)

func GetUserNameById(id int) string {
	var name string
	err := DB.QueryRow("select user_name from blog_user where uid=?", id).Scan(&name)
	if err != nil {
		log.Println("GetUserNameById 查询出错:", err)
	}
	return name
}

func GetUserUserInfo(userName, password string) (*models.UserInfo, error) {
	var userInfo = &models.UserInfo{}
	err := DB.QueryRow("select uid,user_name,avatar from blog_user where user_name=? and password=?", userName, password).Scan(&userInfo.Uid, &userInfo.UserName, &userInfo.Avatar)
	if err != nil {
		log.Println("GetUserInfo 查询出错:", err)
		return nil, err
	}
	return userInfo, nil
}

func GetUser(userName, password string) (*models.User, error) {
	var user = &models.User{}
	err := DB.QueryRow("select uid,user_name,password from blog_user where user_name=? and password=?", userName, password).Scan(
		&user.Uid, &user.UserName, &user.Password, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println("GetUser 查询出错:", err)
		return nil, err
	}
	return user, nil
}
