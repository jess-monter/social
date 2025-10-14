package store

import (
	"context"
	"database/sql"
	"time"
)

var QueryTimeoutDuration = time.Second * 5

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		Delete(context.Context, int64) error
		Update(context.Context, *Post) error
		GetPostByID(context.Context, int64) (*Post, error)
		GetUserFeed(context.Context, int64) ([]FeedPost, error)
	}
	Users interface {
		Create(context.Context, *User) error
		GetUserByID(context.Context, int64) (*User, error)
	}
	Comments interface {
		GetCommentsByPostID(context.Context, int64) ([]*Comment, error)
		Create(context.Context, *Comment) error
	}
	Followers interface {
		Follow(context.Context, int64, int64) error
		Unfollow(context.Context, int64, int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db},
		Followers: &FollowerStore{db},
	}
}
