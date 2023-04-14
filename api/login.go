package api

import (
	"blog/common"
	"blog/models"
	"blog/service"
	"net/http"
)

func (a *ApiHandler) Login(w http.ResponseWriter, r *http.Request) {
	//用户，密码
	name := r.FormValue("name")
	password := r.FormValue("password")

	service.Login(name, password)
	//验证用户密码
	common.Success(w, models.User{})
}
