package models

type Category struct {
	Cid       int    `json:"cid"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


type CategoryRes struct {
	HomeData
	CategoryName string 
}