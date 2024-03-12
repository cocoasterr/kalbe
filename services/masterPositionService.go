package services

import (
	PGRepository "github.com/cocoasterr/kalbe_test/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/kalbe_test/models"
)

type MasterPositionService struct {
	BaseService
}

func NewMasterPositionService(mutation PGRepository.Repository) *MasterPositionService {
	return &MasterPositionService{
		BaseService{
			Repo:  mutation,
			Model: &models.MasterPosition{},
		},
	}
}
