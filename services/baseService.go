package services

import (
	"context"
	"database/sql"

	PGRepository "github.com/cocoasterr/kalbe_test/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/kalbe_test/models"
)

type BaseServiceInterface interface {
	CreateService(ctx context.Context, trx *sql.Tx, payload map[string]interface{}) error
	IndexService(ctx context.Context, page, limit int) ([]map[string]interface{}, interface{}, error)
	FindByService(ctx context.Context, key string, value interface{}) (map[string]interface{}, error)
	UpdateService(ctx context.Context, trx *sql.Tx, payload map[string]interface{}, id string) error
	DeleteProduct(ctx context.Context, id string) error
}

type BaseService struct {
	Repo  PGRepository.Repository
	Model models.BaseModel
}

func NewService(productRepo PGRepository.Repository) *BaseService {
	return &BaseService{
		Repo: productRepo,
	}
}

func (s *BaseService) CreateService(ctx context.Context, trx *sql.Tx, payload map[string]interface{}) error {
	createPayload := s.Model.ModelCreate(payload)

	// trx, err := s.Repo.BeginTransaction()
	// if err != nil {
	// 	return err
	// }

	// defer func() {
	// 	if commitErr := s.Repo.CommitTransaction(trx); commitErr != nil {
	// 		logrus.WithError(commitErr).Errorf("Error committing transaction")
	// 	}
	// 	if err != nil {
	// 		rollbackErr := s.Repo.RollbackTransaction(trx)
	// 		if rollbackErr != nil {
	// 			logrus.WithError(rollbackErr).Error("Error rolling back transaction")
	// 		}
	// 	}
	// }()

	return s.Repo.Create(ctx, createPayload)
}

func (s *BaseService) IndexService(ctx context.Context, page, limit int) ([]map[string]interface{}, interface{}, error) {
	datares, total, err := s.Repo.Index(ctx, page, limit)

	if err != nil {
		return nil, nil, err
	}

	return datares, total, nil
}

func (s *BaseService) FindByService(ctx context.Context, key string, value interface{}) (map[string]interface{}, error) {
	datares, err := s.Repo.FindBy(ctx, key, value)
	if err != nil {
		return nil, err
	}
	if len(datares) != 1 {
		return nil, err
	}
	return datares[0], nil
}

func (s *BaseService) FindCustomQueryService(ctx context.Context, query string) (map[string]interface{}, error) {

	return s.Repo.FindCustomQuery(ctx, query)
}

func (s *BaseService) UpdateService(ctx context.Context, trx *sql.Tx, payload map[string]interface{}, id string) error {
	updatePayload := s.Model.ModelUpdate(payload)

	return s.Repo.Update(ctx, id, updatePayload)
}

func (s *BaseService) SoftDelService(ctx context.Context, trx *sql.Tx, payload map[string]interface{}, id string) error {
	softDelPayload := s.Model.ModelSoftDel(payload)

	return s.Repo.Update(ctx, id, softDelPayload)
}
