package models

import (
	"github.com/jinzhu/gorm"
)

type Customer struct {
	CustomerID int    `gorm:"primary_key" json:"customer_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
}

// CheckCAuth checks if authentication information exists
func CheckCAuth(username, password string) (int, bool, error) {
	var auth Customer
	err := db.Select("customer_id").Where(Customer{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, false, err
	}

	if auth.CustomerID > 0 {
		return auth.CustomerID, true, nil
	}

	return 0, false, nil
}

func GetCustomerByID(id int)(*Customer, error){
	var customer Customer
	err := db.Select("customer_id, name, phone").Where("customer_id = ?", id).First(&customer).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func EditCustomerPassword(id int, newPassword string) error {
	return db.Model(Customer{}).Where("customer_id = ?",id).Update("password", newPassword).Error
}

func AddCustomer(username, password, name string)(int, error){
	customer := Customer{
		Name: name,
		Username: username,
		Password: password,
		Phone: username,
	}

	err := db.Create(&customer).Error
	if err != nil {
		return -1, err
	}

	return customer.CustomerID, nil
}

func ExistByPhone(phone string)(bool, error){
	var cus Customer
	err := db.Select("customer_id").Where("phone = ?", phone).First(&cus).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if cus.CustomerID > 0 {
		return true, nil
	}

	return false, nil
}

func GetCustomers(pageNum, pageSize int)([]*Customer, error){
	var customers []*Customer
	err := db.Offset(pageNum).Limit(pageSize).Find(&customers).Error
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func CountCustomers() (int, error) {
	var count int
	err := db.Model(Customer{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}