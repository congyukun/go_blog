package dao

import "log"

func GetUserNameById(id int) string {
	var name string
	err := DB.QueryRow("select user_name from blog_user where uid=?", id).Scan(&name)
	if err != nil {
		log.Println("GetUserNameById 查询出错:", err)
	}
	return name
}
