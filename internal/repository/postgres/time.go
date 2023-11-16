package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"task-for-docker/internal/entity"
	"time"
)

var (
	ErrTimeNotExist = errors.New("such time does not exist")
)

type Repo struct {
	db *sqlx.DB
}

func NewTimeRepository(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) CreateTime(ctx context.Context, t time.Time) error {
	var id int64
	id = 1
	query := `INSERT INTO time_table (id, last_time) VALUES ($1, $2)
		ON CONFLICT(id) DO UPDATE SET last_time = EXCLUDED.last_time`
	_, err := r.db.ExecContext(ctx, query, id, t)
	if err != nil {
		return fmt.Errorf("postgres - Repo - CreateTime - r.db.ExecContext: %w", err)
	}
	return nil
}

func (r *Repo) GetTime(ctx context.Context, t *entity.Data) error {
	var id int64
	id = 1

	query := `SELECT last_time FROM time_table WHERE id=$1 `
	err := r.db.GetContext(ctx, &t.Time, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrTimeNotExist // return 0, fmt.Errorf("postgres - UsersRepositoryImpl - GetBalance - tx.GetContext: %w", ErrUserNotExist)
		}
		return fmt.Errorf("postgres - Repo - GetTime - r.db.GetContext: %w", err)
	}

	return nil
}

//func (r Repo) GetBalance(ctx context.Context, id uint64) (float32, error) {
//	var balance float32
//	query := `SELECT balance FROM user_info WHERE id=$1 `
//	err := r.db.GetContext(ctx, &balance, query, id)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return 0, ErrUserNotExist // return 0, fmt.Errorf("postgres - UsersRepositoryImpl - GetBalance - tx.GetContext: %w", ErrUserNotExist)
//		}
//		return 0, fmt.Errorf("postgres - UsersRepositoryImpl - GetBalance - tx.GetContext: %w", err)
//	}
//	return balance, nil
//}
