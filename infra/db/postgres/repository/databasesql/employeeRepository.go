package PGRepository

import (
	"database/sql"

	"github.com/cocoasterr/kalbe_test/models"
)

type EmployeeRepository struct {
	Repository
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{
		Repository{
			Db:       db,
			Model:    &models.Employee{},
			PrefixID: "employee",
		},
	}
}
