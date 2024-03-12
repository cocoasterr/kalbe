package PGRepository

import (
	"database/sql"

	"github.com/cocoasterr/kalbe_test/models"
)

type MasterDepartmentRepository struct {
	Repository
}

func NewMasterDepartmentRepository(db *sql.DB) *MasterDepartmentRepository {
	return &MasterDepartmentRepository{
		Repository{
			Db:       db,
			Model:    &models.MasterDepartment{},
			PrefixID: "department",
		},
	}
}
