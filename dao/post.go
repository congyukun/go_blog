package dao

import "blog/models"

func GetAllPost(page, pageSize int) ([]models.PostMore, error) {
	page = (page - 1) * pageSize
	row, err := DB.Query("select * from blog_post limit ?,?", page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.PostMore
	for row.Next() {
		var post models.PostMore
		err := row.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetTotal() (int, error) {
	var total int
	err := DB.QueryRow("select count(pid) from blog_post").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
