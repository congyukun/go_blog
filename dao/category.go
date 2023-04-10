package dao

import (
	"blog/models"
	"log"
)

func GetAllCategory() ([]models.Category, error) {

	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("GetAllCategory 查询出错:", err)
		return nil, err
	}
	var categories []models.Category

	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Cid, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			log.Println("GetAllCategory 查询出错:", err)
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func GetCategoryNameById(id int) string {
	var name string
	err := DB.QueryRow("select name from blog_category where cid=?", id).Scan(&name)
	if err != nil {
		log.Println("GetCategoryNameById 查询出错:", err)
	}
	return name
}

