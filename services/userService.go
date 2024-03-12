package services

import (
	PGRepository "github.com/cocoasterr/kalbe_test/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/kalbe_test/models"
)

type UserService struct {
	BaseService
}

func NewUserService(mutation PGRepository.Repository) *UserService {
	return &UserService{
		BaseService{
			Repo:  mutation,
			Model: &models.User{},
		},
	}
}
