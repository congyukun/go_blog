package models

type User struct {
	Uid      int    `json:"uid"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
}

type UserInfo struct {
	Uid int `json:"uid"`
	UserName string `json:"userName"`
	Avatar string `json:"avatar"`
}
