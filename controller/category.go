package controller

import (
	"blog/common"
	"blog/service"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (Html *HtmlApi) Category(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	categoryTemp := common.Template.Category
	cIdStr := strings.TrimPrefix(path, "/c/")

	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		log.Println("Category:", err)
		categoryTemp.ExecuteError(w, errors.New("不识别此请求路径！"))
		return
	}
	if err := r.ParseForm(); err != nil {
		categoryTemp.ExecuteError(w, errors.New("服务器错误"))
		return
	}

	pageStr := r.Form.Get("page")

	if pageStr == "" {
		pageStr = "1"
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		categoryTemp.ExecuteError(w, errors.New("服务器错误"))
		return
	}
	pageSize := 10

	categoryResponse,err := service.GetPostByCategoryId(cId, page, pageSize)

	if err != nil {
		categoryTemp.ExecuteError(w,errors.New("服务器错误"))
		return
	} 

	categoryTemp.ExecuteTemp(w, categoryResponse)

}
