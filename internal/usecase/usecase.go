package usecase

import (
	"context"
	"task-for-docker/internal/entity"
	"task-for-docker/internal/repository/postgres"
)

type UseCase struct {
	Timer
}

type Timer interface {
	CreateTime(ctx context.Context) error
	GetTime(ctx context.Context) (entity.Data, error)
}

func New(repo *postgres.Repository) *UseCase {
	return &UseCase{
		Timer: NewTimerUseCase(repo.TimeRepository),
	}
}
