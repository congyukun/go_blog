package service

import (
	"blog/config"
	"blog/dao"
	"blog/models"
	"html/template"
)

func GetAllIndexInfo(page, pageSize int) (*models.HomeData, error) {
	categories, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}

	posts, err := dao.GetAllPost(page, pageSize)
	if err != nil {
		return nil, err
	}
	var postMores []models.PostMore

	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Markdown:     post.Markdown,
			Content:      template.HTML(content),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.Date(post.CreateAt),
			UpdateAt:     models.Date(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	total, err := dao.GetTotal()
	if err != nil {
		return nil, err
	}
	var pages []int
	pagesCount := (total-1)/10 + 1
	return &models.HomeData{
		Viewer: config.Config.Viewer,
		Categories: categories,
		Posts: postMores,
		Total: total,
		Page: page,
		Pages: pages,
		PageEnd: page != pagesCount,
	}, nil
}
