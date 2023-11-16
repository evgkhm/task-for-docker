package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"task-for-docker/internal/entity"
	"time"
)

type Repository struct {
	TimeRepository
}

type TimeRepository interface {
	CreateTime(ctx context.Context, t time.Time) error
	GetTime(ctx context.Context, t *entity.Data) error
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		TimeRepository: NewTimeRepository(db),
	}
}
