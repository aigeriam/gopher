package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type PostStore struct {
	db *sql.DB
}
type Post struct {
	ID        int64    `json: "id"`
	Content   string   `json: "content`
	Title     string   `json: "title"`
	UserId    int64    `json: "user_id"`
	CreatedAt string   `json: "created_at`
	UpdatedAt string   `json: "updated_at"`
	Tags      []string `json: "tags"`
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
	INSERT INTO posts (content, title, user_id, tags)
	VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`
	err := s.db.QueryRowContext(
		ctx, query, post.Content, post.Title, post.UserId, pq.Array(post.Tags),
	).Scan(
		&post.ID, ///when creating post they will be added to post sent
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) GetByID(ctx context.Context, postID string) (Post, error) {
	query:=`
	SElECT FROM posts WHERE `
	query:=
}
