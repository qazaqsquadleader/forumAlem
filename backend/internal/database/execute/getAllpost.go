package execute

import (
	"database/sql"
	"forum-backend/internal/models"
)

func GetAllpostSql(db *sql.DB) ([]models.AllPosts, error) {
	var Allposts []models.AllPosts
	query := `SELECT * FROM posts`
	rows, err := db.Query(query)
	if err != nil {
		return Allposts, err
	}
	for rows.Next() {
		var post models.AllPosts
		if err := rows.Scan(&post.PostId, &post.Author, &post.Title, &post.Content, &post.CreationDate); err != nil {
			return nil, err
		}
		Allposts = append(Allposts, post)
	}
	return Allposts, nil
}
