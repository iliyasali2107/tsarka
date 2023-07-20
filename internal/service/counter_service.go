package service

import (
	"context"

	"tsarka/internal/repository"
)

type CounterService interface {
	Increment(context.Context, int64) (int64, error)
	Decrement(context.Context, int64) (int64, error)
	GetValue(context.Context) (int64, error)
}

type counterService struct {
	CounterRepo repository.CounterRepository
}

func NewCounterService(counterRepo repository.CounterRepository) CounterService {
	return &counterService{
		CounterRepo: counterRepo,
	}
}

func (cs *counterService) Increment(ctx context.Context, n int64) (int64, error) {
	res, err := cs.CounterRepo.Increment(ctx, n)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (cs *counterService) Decrement(ctx context.Context, n int64) (int64, error) {
	res, err := cs.CounterRepo.Decrement(ctx, n)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (cs *counterService) GetValue(ctx context.Context) (int64, error) {
	res, err := cs.CounterRepo.GetValue(ctx)
	if err != nil {
		return 0, err
	}

	return res, nil
}
