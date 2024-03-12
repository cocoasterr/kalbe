package PGRepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/cocoasterr/kalbe_test/models"
)

type RepoInterface interface {
	Create(ctx context.Context, trx *sql.Tx, model map[string]interface{}) error
	Update(repository *Repository, ctx context.Context, trx *sql.Tx, id string, model map[string]interface{}) error
	Index(ctx context.Context, page, limit int, tbName string) ([]interface{}, interface{}, error)
	FindBy(ctx context.Context, tbName, key string, value interface{}) (map[string]interface{}, error)
	Delete(ctx context.Context, tbname, key string, value interface{}) error
}
type Repository struct {
	Db       *sql.DB
	Model    models.BaseModel
	PrefixID string
}

func (r *Repository) Create(ctx context.Context, model map[string]interface{}) error {
	tbName := r.Model.TbName()
	var key, value []string
	var values []interface{}
	var i int
	for k, v := range model {
		key = append(key, k)
		value = append(value, fmt.Sprintf("$%d", i+1))
		values = append(values, v)
		i++
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tbName, strings.Join(key, ", "), strings.Join(value, ", "))

	trx, err := r.Db.Begin()
	if err != nil {
		return err
	}
	_, err = trx.ExecContext(ctx, query, values...)
	if err != nil {
		trx.Rollback()
		return err
	}
	err = trx.Commit()
	if err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func (r *Repository) Index(ctx context.Context, page, limit int) ([]map[string]interface{}, interface{}, error) {
	tbName := r.Model.TbName()
	offset := (page - 1) * limit
	query := fmt.Sprintf("SELECT * FROM %s WHERE deleted_at IS NULL LIMIT %d OFFSET %d", tbName, limit, offset)
	rows, _ := r.Db.QueryContext(ctx, query)
	if rows == nil {
		return nil, nil, nil
	}
	data, _ := getResIndex(rows)

	var total int
	query_total := fmt.Sprintf("SELECT COUNT(%s_id) FROM %s WHERE deleted_at IS NULL", r.PrefixID, tbName)
	if err := r.Db.QueryRowContext(ctx, query_total).Scan(&total); err != nil {
		return nil, nil, err
	}

	return data, total, nil
}
func (r *Repository) FindBy(ctx context.Context, key string, value interface{}) ([]map[string]interface{}, error) {
	tbName := r.Model.TbName()

	typeValue := reflect.ValueOf(value)
	if typeValue.Kind() == reflect.String {
		value = fmt.Sprintf("'%s'", value)
	}
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s_%s = %v", tbName, r.PrefixID, key, value)
	rows, _ := r.Db.QueryContext(ctx, query)
	if rows == nil {
		return nil, nil
	}
	data, _ := getResIndex(rows)

	return data, nil
}

func (r *Repository) FindCustumQuery(ctx context.Context, query string) ([]map[string]interface{}, error) {
	// query = "select * from attendance"
	rows, _ := r.Db.QueryContext(ctx, query)
	if rows == nil {
		return nil, nil
	}
	data, _ := getResIndex(rows)

	return data, nil
}

func (r *Repository) Update(ctx context.Context, id string, model map[string]interface{}) error {
	tbName := r.Model.TbName()

	existingData, err := r.FindBy(ctx, "id", id)
	if err != nil {
		return err
	}
	if len(existingData) == 0 {
		return errors.New("data not found")
	}

	var updates []string
	var values []interface{}
	var i int
	for k, v := range model {
		updates = append(updates, fmt.Sprintf("%s=$%d", k, i+1))
		values = append(values, v)
		i++
	}
	values = append(values, id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s_id = $%d", tbName, strings.Join(updates, ", "), r.PrefixID, i+1)

	trx, err := r.Db.Begin()
	if err != nil {
		return err
	}
	_, err = trx.ExecContext(ctx, query, values...)
	if err != nil {
		trx.Rollback()
		return err
	}
	err = trx.Commit()
	if err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func (r *Repository) FindCustomQuery(ctx context.Context, query string) (map[string]interface{}, error) {
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := getResIndex(rows)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	return data[0], nil
}
func getResIndex(rows *sql.Rows) ([]map[string]interface{}, error) {
	lisColumn, _ := rows.Columns()
	var res []map[string]interface{}
	var sampleSliceUint8 []uint8
	for rows.Next() {
		dest := make([]interface{}, len(lisColumn))
		for i := range lisColumn {
			dest[i] = new(interface{})
		}
		err := rows.Scan(dest...)
		if err != nil {
			return nil, err
		}
		itermap := make(map[string]interface{})
		for i, colName := range lisColumn {
			value := *dest[i].(*interface{})
			switch reflect.TypeOf(value) {
			case reflect.TypeOf(time.Time{}):
				toStringFormatDate := time.Date(2012, time.February, 23, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
				value = toStringFormatDate
			case reflect.TypeOf(sampleSliceUint8):
				stringValue := string(value.([]uint8))
				floatValue, _ := strconv.ParseFloat(stringValue, 64)
				value = int(floatValue)
			}
			itermap[colName] = value
		}
		res = append(res, itermap)
	}
	return res, nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	tbName := r.Model.TbName()

	existingData, err := r.FindBy(ctx, "id", id)
	if err != nil {
		return err
	}
	if len(existingData) == 0 {
		return errors.New("data not found")
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tbName)
	trx, err := r.Db.Begin()
	if err != nil {
		return err
	}
	_, err = trx.ExecContext(ctx, query, id)
	if err != nil {
		trx.Rollback()
		return err
	}
	err = trx.Commit()
	if err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func (r *Repository) BeginTransaction() (*sql.Tx, error) {
	trx, err := r.Db.Begin()
	if err != nil {
		return nil, err
	}
	return trx, nil
}

func (r *Repository) CommitTransaction(trx *sql.Tx) error {
	return trx.Commit()
}

func (r *Repository) RollbackTransaction(trx *sql.Tx) error {
	return trx.Rollback()
}
