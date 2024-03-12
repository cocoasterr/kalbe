package models

import (
	"time"

	"github.com/google/uuid"
)

type MasterDepartment struct {
	Department_id   string    `json:"department_id"`
	Department_name string    `json:"department_name"`
	Created_at      time.Time `json:"created_at"`
	Created_by      string    `json:"created_by"`
	Updated_at      time.Time `json:"updated_at"`
	Updated_by      string    `json:"updated_by"`
	Deleted_at      time.Time `json:"deleted_at"`
}

func (m *MasterDepartment) Model() interface{} {
	return &MasterDepartment{}
}

func (m *MasterDepartment) ModelCreate(payload map[string]interface{}) map[string]interface{} {
	payload["department_id"] = uuid.New().String()
	payload["created_at"] = time.Now()
	payload["updated_at"] = time.Now()
	return payload
}

func (m *MasterDepartment) ModelUpdate(payload map[string]interface{}) map[string]interface{} {
	if payload["department_name"] != "" {
		payload["updated_at"] = time.Now()
		return payload
	}
	return nil
}
func (m *MasterDepartment) ModelSoftDel(payload map[string]interface{}) map[string]interface{} {
	payload["delete_at"] = time.Now()
	return payload
}
func (m *MasterDepartment) TbName() string {
	return "MasterDepartment"
}
