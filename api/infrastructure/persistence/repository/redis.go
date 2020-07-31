package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type redisRepository struct {
	Client *redis.Client
	Ctx    context.Context
}
