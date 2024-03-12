package services

import (
	"context"
	"errors"

	PGRepository "github.com/cocoasterr/kalbe_test/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/kalbe_test/models"
)

type AttendaceService struct {
	BaseService
}
type AttendanceServiceAdditional struct {
	PGRepository.AttendanceInterface
}

func NewAttendaceService(mutation PGRepository.Repository) *AttendaceService {
	return &AttendaceService{
		BaseService{
			Repo:  mutation,
			Model: &models.Attendance{},
		},
	}
}

func (s *AttendanceServiceAdditional) GetReportService(ctx context.Context, startDate, endDate string) (map[string]interface{}, error) {
	data, err := s.GetReportAttendandRepository(ctx, startDate, endDate)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return data, nil
}
