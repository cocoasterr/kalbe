package services

import (
	PGRepository "github.com/cocoasterr/kalbe_test/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/kalbe_test/models"
)

type MasterDepartmentService struct {
	BaseService
}

func NewMasterDepartmentService(mutation PGRepository.Repository) *MasterDepartmentService {
	return &MasterDepartmentService{
		BaseService{
			Repo:  mutation,
			Model: &models.MasterDepartment{},
		},
	}
}
