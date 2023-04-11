package router

import (
	"blog/api"
	"blog/config"
	"blog/controller"
	"net/http"
)

func Router() {
	// 1.页面 2.json 3.静态资源
	// http.HandleFunc("/", func(w http.Resp去2111·	1·1111111onseWriter, r *http.Request) {
	// 	w.Write([]byte("hello world"))
	// })

	// view
	http.HandleFunc("/c/", controller.Html.Category)
	http.HandleFunc("/login", controller.Html.Login)
	http.HandleFunc("/", controller.Html.Index)
	http.HandleFunc("/api/v1/login", controller.Html.Index) 
	//api
	http.HandleFunc("/api", api.Api.Index)
	// 静态资源
	http.Handle("/resource/", http.StripPrefix("/resource/",
		http.FileServer(http.Dir(config.Config.SystemConfig.CurrentDir+"/public/resource"))))
}
