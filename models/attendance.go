package models

import (
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	Attendance_id string    `json:"attendance_id"`
	Employee_id   string    `json:"employee_id"`
	Location_id   string    `json:"location_id"`
	Absent_in     time.Time `json:"absent_in"`
	Absent_out    time.Time `json:"absent_out"`
	Created_at    time.Time `json:"created_at"`
	Created_by    string    `json:"created_by"`
	Updated_at    time.Time `json:"updated_at"`
	Updated_by    string    `json:"updated_by"`
	Deleted_at    time.Time `json:"deleted_at"`
}

func (m *Attendance) Model() interface{} {
	return &Attendance{}
}

func (m *Attendance) ModelCreate(payload map[string]interface{}) map[string]interface{} {
	payload["attendance_id"] = uuid.New().String()
	payload["created_at"] = time.Now()
	payload["updated_at"] = time.Now()
	return payload
}

func (m *Attendance) ModelUpdate(payload map[string]interface{}) map[string]interface{} {
	if payload["absent_in"] != "" {
		payload["updated_at"] = time.Now()
		return payload
	}
	return nil
}
func (m *Attendance) ModelSoftDel(payload map[string]interface{}) map[string]interface{} {
	payload["delete_at"] = time.Now()
	return payload
}

func (m *Attendance) TbName() string {
	return "Attendance"
}
