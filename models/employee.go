package models

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	Employee_id   string    `json:"employee_id"`
	Employee_code string    `json:"employee_code"`
	Employee_name string    `json:"employee_name"`
	Password      string    `json:"password"`
	Department_id string    `json:"department_id"`
	Position_id   string    `json:"position_id"`
	Superior      int       `json:"superior"`
	Created_at    time.Time `json:"created_at"`
	Created_by    string    `json:"created_by"`
	Updated_at    time.Time `json:"updated_at"`
	Updated_by    string    `json:"updated_by"`
	Deleted_at    time.Time `json:"deleted_at"`
}

func (m *Employee) Model() interface{} {
	return &Employee{}
}

func (m *Employee) ModelCreate(payload map[string]interface{}) map[string]interface{} {
	payload["employee_id"] = uuid.New().String()
	payload["created_at"] = time.Now()
	payload["updated_at"] = time.Now()
	return payload
}

func (m *Employee) ModelUpdate(payload map[string]interface{}) map[string]interface{} {
	if payload["employee_name"] != "" {
		payload["updated_at"] = time.Now()
		return payload
	}
	return nil
}
func (m *Employee) ModelSoftDel(payload map[string]interface{}) map[string]interface{} {
	payload["delete_at"] = time.Now()
	return payload
}
func (m *Employee) TbName() string {
	return "Employee"
}
