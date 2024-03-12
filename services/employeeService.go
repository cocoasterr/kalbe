package services

import (
	PGRepository "github.com/cocoasterr/kalbe_test/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/kalbe_test/models"
)

type EmployeeService struct {
	BaseService
}

func NewEmployeeService(mutation PGRepository.Repository) *EmployeeService {
	return &EmployeeService{
		BaseService{
			Repo:  mutation,
			Model: &models.Employee{},
		},
	}
}
