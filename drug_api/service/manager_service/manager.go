package manager_service

import "drug_api/models"

type Manager struct {
	ManagerID       int    `json:"manager_id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	NewPassword     string
}

func (m *Manager) Check() (bool, error) {
	return models.CheckMAuth(m.Username, m.Password)
}

func (m *Manager)EditManagerPassword()error{
	return models.EditManagerPassword(m.ManagerID, m.NewPassword)
}

func (m *Manager)AddManager()(int, error){
	return models.AddManager(m.Username, m.Password)
}

func (m *Manager)ExistByUsername()(bool, error){
	return models.ExistManagerByUsername(m.Username)
}