package PGRepository

import (
	"database/sql"

	"github.com/cocoasterr/kalbe_test/models"
)

type MasterLocationRepository struct {
	Repository
}

func NewMasterLocationRepository(db *sql.DB) *MasterLocationRepository {
	return &MasterLocationRepository{
		Repository{
			Db:       db,
			Model:    &models.MasterLocation{},
			PrefixID: "location",
		},
	}
}
