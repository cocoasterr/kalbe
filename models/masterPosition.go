package models

import (
	"time"

	"github.com/google/uuid"
)

type MasterPosition struct {
	Position_id   string    `json:"position_id"`
	Department_id string    `json:"department_id"`
	Position_name string    `json:"position_name"`
	Created_at    time.Time `json:"created_at"`
	Created_by    string    `json:"created_by"`
	Updated_at    time.Time `json:"updated_at"`
	Updated_by    string    `json:"updated_by"`
	Deleted_at    time.Time `json:"deleted_at"`
}

func (m *MasterPosition) Model() interface{} {
	return &MasterPosition{}
}

func (m *MasterPosition) ModelCreate(payload map[string]interface{}) map[string]interface{} {
	payload["position_id"] = uuid.New().String()
	payload["created_at"] = time.Now()
	payload["updated_at"] = time.Now()
	return payload
}

func (m *MasterPosition) ModelUpdate(payload map[string]interface{}) map[string]interface{} {
	if payload["position_name"] != "" {
		payload["updated_at"] = time.Now()
		return payload
	}
	return nil
}

func (m *MasterPosition) ModelSoftDel(payload map[string]interface{}) map[string]interface{} {
	payload["delete_at"] = time.Now()
	return payload
}

func (m *MasterPosition) TbName() string {
	return "MasterPosition"
}
