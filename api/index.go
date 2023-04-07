package api

import (
	"encoding/json"
	"net/http"
)

var Api = &ApiHandler{}
type ApiHandler struct {

}

func (api *ApiHandler) Index(w http.ResponseWriter, r *http.Request) {
	json.Marshal(map[string]interface{}{
		"code": 200,
		"msg":  "success",
	})
}