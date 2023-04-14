package models

type Result struct {
	Error string      `json:"error"`
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}
