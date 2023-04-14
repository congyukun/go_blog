package common

import (
	"blog/config"
	"blog/models"
	"encoding/json"
	"net/http"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
    w := sync.WaitGroup{}
    w.Add(1)
    go func() {
        // 耗时
        var err error
        Template, err = models.InitTemplate(config.Config.SystemConfig.CurrentDir + "/template/")
        if err != nil {
            panic(err)
        }
        w.Done()
    }()
    w.Wait()
}


func Success(w http.ResponseWriter, data interface{}) {
    var result models.Result
    result.Error = ""
    result.Code = 200
    result.Msg = "success"
    result.Data = data
    resJson, err := json.Marshal(result)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(resJson)
}

func ErrorRes(w http.ResponseWriter, code int, msg string) {
    var result models.Result
    result.Error = "error"
    result.Code = code
    result.Msg = msg
    result.Data = nil
    resJson, err := json.Marshal(result)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(resJson)
}

func GetRequestJsonParam(r *http.Request) (map[string]interface{}, error) {
    var params map[string]interface{}
    err := json.NewDecoder(r.Body).Decode(&params)
    if err != nil {
        return nil, err
    }
    return params, nil
}