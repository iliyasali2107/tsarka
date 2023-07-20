package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const redisKey = "counter"

type CounterRepository interface {
	Increment(context.Context, int64) (int64, error)
	Decrement(context.Context, int64) (int64, error)
	GetValue(context.Context) (int64, error)
}

type counterRepository struct {
	DB *redis.Client
}

func NewCounterRepository(db *redis.Client) CounterRepository {
	return &counterRepository{
		DB: db,
	}
}

func (cr *counterRepository) Increment(ctx context.Context, n int64) (int64, error) {
	res, err := cr.DB.IncrBy(ctx, redisKey, int64(n)).Result()
	if err != nil {
		return 0, err
	}

	return int64(res), nil
}

func (cr *counterRepository) Decrement(ctx context.Context, n int64) (int64, error) {
	res, err := cr.DB.DecrBy(ctx, redisKey, int64(n)).Result()
	if err != nil {
		return 0, err
	}

	return int64(res), nil
}

func (cr *counterRepository) GetValue(ctx context.Context) (int64, error) {
	res, err := cr.DB.Get(ctx, redisKey).Int64()
	if err != nil {
		return 0, err
	}

	return int64(res), nil
}
