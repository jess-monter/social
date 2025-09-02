package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	ID        int64  `json:"id"`
	PostID    int64  `json:"post_id"`
	UserID    int64  `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) GetCommentsByPostID(ctx context.Context, postID int64) ([]*Comment, error) {
	var comments []*Comment
	query := `
		SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, u.username FROM comments c
		JOIN users u ON c.user_id = u.id
		WHERE c.post_id = $1
		ORDER BY c.created_at DESC
	`
	rows, err := s.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c Comment
		if err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt, &c.User.Username); err != nil {
			return nil, err
		}
		comments = append(comments, &c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
