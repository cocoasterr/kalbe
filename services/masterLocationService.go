package services

import (
	PGRepository "github.com/cocoasterr/kalbe_test/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/kalbe_test/models"
)

type MasterLocationtService struct {
	BaseService
}

func NewMasterLocationtService(mutation PGRepository.Repository) *MasterLocationtService {
	return &MasterLocationtService{
		BaseService{
			Repo:  mutation,
			Model: &models.MasterLocation{},
		},
	}
}
