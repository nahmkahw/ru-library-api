package repository

import (
	"ru-library-api/entity"

	"github.com/jmoiron/sqlx"
)

type (
	sierraRepoDB struct {
		postgres_db *sqlx.DB
	}

	SierraRepoInterface interface {
		FineId(fines *[]entity.FineRepo, personId string) error
		PatronId(fines *[]entity.PatronRepo, personId string) error
	}
)

func NewSierraRepo(postgres_db *sqlx.DB) SierraRepoInterface {
	return &sierraRepoDB{postgres_db: postgres_db}
}
