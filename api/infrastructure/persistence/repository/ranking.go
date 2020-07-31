package repository

import (
	"context"
	"encoding/json"
	"influ-dojo/api/domain/apperr"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/dto"

	"github.com/go-redis/redis/v8"
	"golang.org/x/xerrors"
)

const Ranking = "ranking"
const RankingAll = "ranking_all"

type ranking struct {
	redisRepository
}

func NewRanking(client *redis.Client) repository.Ranking {
	return &ranking{redisRepository{
		Client: client,
		Ctx:    context.Background(),
	}}
}

func (repo *ranking) LoadAll() (*dto.RankingAll, error) {
	record, err := repo.Client.HGet(repo.Ctx, Ranking, RankingAll).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, apperr.ErrRecordNotFound
		}

		return nil, xerrors.Errorf("failed to load ranking: %w", err)
	}

	all := new(dto.RankingAll)
	if err := json.Unmarshal([]byte(record), all); err != nil {
		return nil, xerrors.Errorf("failed to decode from JSON: %w", err)
	}

	return all, nil
}

func (repo *ranking) Store(all *dto.RankingAll) error {
	b, err := json.Marshal(all)
	if err != nil {
		return xerrors.Errorf("failed to encode to JSON: %w", err)
	}

	if err := repo.Client.HSet(repo.Ctx, Ranking, map[string]interface{}{RankingAll: string(b)}).Err(); err != nil {
		return xerrors.Errorf("failed to store ranking: %w", err)
	}

	return nil
}
