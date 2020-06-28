package models

import "github.com/jinzhu/gorm"

type Employee struct {
	EmployeeID int    `gorm:"primary_key" json:"employee_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Position   string `json:"position"`
}

// CheckEAuth checks if authentication information exists
func CheckEAuth(username, password string) (int, bool, error) {
	var auth Employee
	err := db.Select("employee_id").Where(Employee{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, false, err
	}

	if auth.EmployeeID > 0 {
		return auth.EmployeeID, true, nil
	}

	return 0, false, nil
}

func GetEmployeeByID(id int) (*Employee, error){
	var employee Employee
	err := db.Select("employee_id, name, position").Where("employee_id = ?", id).First(&employee).Error
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func EditEmployeePassword(id int, newPassword string) error {
	return db.Model(Employee{}).Where("employee_id = ?",id).Update("password", newPassword).Error
}

func EditEmployeePosition(id int, newPosition string) error {
	return db.Model(Employee{}).Where("employee_id = ?",id).Update("position", newPosition).Error
}

func AddEmployee(username, password, name, positon string)(int, error){
	employee := Employee{
		Name: name,
		Position: positon,
		Username: username,
		Password: password,
	}

	err := db.Create(&employee).Error
	if err != nil {
		return -1, err
	}

	return employee.EmployeeID, nil
}

func ExistByUsername(username string)(bool, error){
	var emp Employee
	err := db.Select("employee_id").Where("username = ?", username).First(&emp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if emp.EmployeeID > 0 {
		return true, nil
	}

	return false, nil
}

func GetEmployees(pageNum int, pageSize int)([]*Employee, error){
	var employees []*Employee
	err := db.Offset(pageNum).Limit(pageSize).Find(&employees).Error
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func GetAllEmployees()([]*Employee, error){
	var employees []*Employee
	err := db.Find(&employees).Error
	if err != nil {
		return nil, err
	}

	for _,v:=range employees{
		v.Username = ""
		v.Password = ""
	}

	return employees, nil
}

func GetTotalEmployees() (int, error) {
	var count int
	err := db.Model(Employee{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, err
}