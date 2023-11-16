package usecase

import (
	"context"
	"task-for-docker/internal/entity"
	"task-for-docker/internal/repository/postgres"
	"time"
)

type UserUseCase struct {
	userRepo postgres.TimeRepository
}

func NewTimerUseCase(userRepo postgres.TimeRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (u *UserUseCase) CreateTime(ctx context.Context) error {
	tm := time.Now()
	err := u.userRepo.CreateTime(ctx, tm)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) GetTime(ctx context.Context) (entity.Data, error) {
	t := &entity.Data{}

	err := u.userRepo.GetTime(ctx, t)
	if err != nil {
		return entity.Data{}, err
	}
	return *t, err
}
