package main

import (
	"blog/common"
	"blog/router"
	"log"
	"net/http"
)

func init() {
	common.LoadTemplate()
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:80",
	}

	router.Router()
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
