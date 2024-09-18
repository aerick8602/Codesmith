package repository

import (
	"database/sql"
	"instagram-api/models"
)

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{DB: db}
}
func (r *PostRepository) CreatePost(post *models.Post) error {
	_, err := r.DB.Exec("INSERT INTO posts (caption, image_url, posted_timestamp, user_id) VALUES ($1, $2, $3, $4)",
		post.Caption, post.ImageURL, post.PostedTimestamp, post.UserID)
	return err
}

func (r *PostRepository) GetPost(id int) (*models.Post, error) {
	row := r.DB.QueryRow("SELECT id, caption, image_url, posted_timestamp, user_id FROM posts WHERE id = $1", id)
	var post models.Post
	if err := row.Scan(&post.ID, &post.Caption, &post.ImageURL, &post.PostedTimestamp, &post.UserID); err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) ListUserPosts(userID int) ([]models.Post, error) {
	rows, err := r.DB.Query("SELECT id, caption, image_url, posted_timestamp, user_id FROM posts WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Caption, &post.ImageURL, &post.PostedTimestamp, &post.UserID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
