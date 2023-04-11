package controller

import (
	"blog/common"
	"net/http"
)



func (*HtmlApi) Login(w http.ResponseWriter, r *http.Request){
	login := common.Template.Login
	login.ExecuteTemp(w, nil)
}