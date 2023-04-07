package controller

import (
	"blog/common"
	"blog/service"
	"errors"
	"log"
	"net/http"
	"strconv"
)

var Html = &HtmlApi{}

type HtmlApi struct {
}

func (*HtmlApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	if err := r.ParseForm(); err != nil {
		index.ExecuteError(w, errors.New("服务器错误1"))
		return
	}
	paseStr := r.Form.Get("page")
	page := 1
	if paseStr != "" {
		page, _ = strconv.Atoi(paseStr)
	}

	//每页显示数量
	pageSize := 10

	// 页面上涉及到的所有的数据，必须有定义
	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("Index:", err)
		index.ExecuteError(w, errors.New("服务器错误"))
		return
	}

	if err != nil {
		panic(err)
	}
	index.ExecuteTemp(w, hr)
}
