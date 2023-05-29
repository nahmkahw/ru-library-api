package service

import (
	"ru-library-api/entity"
	"ru-library-api/repository"

	"github.com/go-redis/redis/v8"
)

type (
	sierraService struct {
		sierraRepo  repository.SierraRepoInterface
		redis_cache *redis.Client
	}

	SierraServiceInterface interface {
		PatronById(personId string) (*[]entity.Patron, error)
		FineById(personId string) (*[]entity.Fine, error)
	}
)

func NewSierraService(sierraRepo repository.SierraRepoInterface, redis_cache *redis.Client) SierraServiceInterface {
	return &sierraService{
		sierraRepo:  sierraRepo,
		redis_cache: redis_cache,
	}
}
