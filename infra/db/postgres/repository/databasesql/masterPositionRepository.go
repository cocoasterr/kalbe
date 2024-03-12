package PGRepository

import (
	"database/sql"

	"github.com/cocoasterr/kalbe_test/models"
)

type MasterPositionRepository struct {
	Repository
}

func NewMasterPositionRepository(db *sql.DB) *MasterPositionRepository {
	return &MasterPositionRepository{
		Repository{
			Db:       db,
			Model:    &models.MasterPosition{},
			PrefixID: "position",
		},
	}
}
