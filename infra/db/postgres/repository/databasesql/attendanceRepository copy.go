package PGRepository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cocoasterr/kalbe_test/models"
)

type AttendanceInterface interface {
	GetReportAttendandRepository(ctx context.Context, startDate, endDate string) (map[string]interface{}, error)
}

type AttendanceRepository struct {
	Repository
}

func NewAttendanceRepository(db *sql.DB) *AttendanceRepository {
	return &AttendanceRepository{
		Repository{
			Db:       db,
			Model:    &models.Attendance{},
			PrefixID: "attendance",
		},
	}
}
func (r *AttendanceRepository) GetReportAttendandRepository(ctx context.Context, startDate, endDate string) (map[string]interface{}, error) {

	query := fmt.Sprintf(`
    SELECT
        a.Created_at AS Date,
        e.Employee_code,
        e.Employee_name,
        d.Department_name,
        p.Position_name,
        l.Location_name,
        a.Absent_in,
        a.Absent_out
    FROM
        Attendance a
    JOIN
        Employee e ON a.Employee_id = e.Employee_id
    JOIN
        MasterDepartment d ON e.Department_id = d.Department_id
    JOIN
        MasterPosition p ON e.Position_id = p.Position_id
    JOIN
        MasterLocation l ON a.Location_id = l.Location_id
    WHERE
        a.Absent_in BETWEEN '%s' AND '%s'
    ORDER BY
        a.Created_at
    `, startDate, endDate)

	data, err := r.FindCustomQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	return data, nil
}
