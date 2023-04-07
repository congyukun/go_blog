package models

import "blog/config"


type HomeData struct {
	config.Viewer
	Categories []Category
	Posts []PostMore
	Total int
	Page int
	Pages []int
	PageEnd bool
}