package models

import (
	"github.com/google/uuid"
)

type User struct {
	User_id  string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m *User) Model() interface{} {
	return &MasterPosition{}
}

func (m *User) ModelCreate(payload map[string]interface{}) map[string]interface{} {
	payload["user_id"] = uuid.New().String()
	return payload
}

func (m *User) ModelUpdate(payload map[string]interface{}) map[string]interface{} {

	return nil
}
func (m *User) ModelSoftDel(payload map[string]interface{}) map[string]interface{} {

	return nil
}
func (m *User) TbName() string {
	return "users"
}
