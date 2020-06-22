package models

import "github.com/jinzhu/gorm"

type Manager struct {
	ManagerID int    `gorm:"primary_key" json:"manager_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

// CheckMAuth checks if authentication information exists
func CheckMAuth(username, password string) (bool, error) {
	var auth Manager
	err := db.Select("manager_id").Where(Manager{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ManagerID > 0 {
		return true, nil
	}

	return false, nil
}

func GetManagerByID(id int) (*Manager, error){
	var manager Manager
	err := db.Select("manager_id, username").Where("manager_id = ?", id).First(&manager).Error
	if err != nil {
		return nil, err
	}

	return &manager, nil
}

func EditManagerPassword(id int, newPassword string) error {
	return db.Model(Manager{}).Where("manager_id = ?",id).Update("password", newPassword).Error
}

func AddManager(username, password string)(int, error){
	manager := Manager{
		Username: username,
		Password: password,
	}
	err := db.Create(&manager).Error
	if err != nil {
		return -1, err
	}

	return manager.ManagerID, nil
}

func ExistManagerByUsername(username string)(bool, error){
	var man Manager
	err := db.Select("manager_id").Where("username = ?", username).First(&man).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if man.ManagerID > 0 {
		return true, nil
	}

	return false, nil
}