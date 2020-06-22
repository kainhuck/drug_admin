package models

import (
	"fmt"
	"time"
)

type DrugSaleOrder struct {
	DrugSaleOrderID int       `gorm:"primary_key" json:"drug_sale_order_id"`
	EmployeeID      int       `json:"employee_id"`
	SaleDate        time.Time `gorm:"TYPE:DATETIME" json:"sale_date"`
	CustomerID      int       `json:"customer_id"`
}

type SaleDrugTotalPrice struct {
	DrugSaleOrderID int
	TotalPrice      int
}

// AddDrugSaleOrder 增加销售订单
func AddDrugSaleOrder(eid int, cid int, saleDate time.Time, drugs []InvDrugWithNum) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 1. 生成销售订单,返回id
	dso := DrugSaleOrder{
		EmployeeID: eid,
		CustomerID: cid,
		SaleDate:   saleDate,
	}
	err := tx.Create(&dso).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println("生成销售订单: OK")
	// 2. 根据订单id,插入药品清单
	for _, drug := range drugs {
		sDrug := SaleDrugList{
			InventoryDrugID: drug["InventoryDrugID"],
			Num:             drug["Num"],
			DrugSaleOrderID: dso.DrugSaleOrderID,
		}

		err := tx.Create(&sDrug).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		fmt.Println("生成清单订单: OK", drug["InventoryDrugID"])
	}

	return tx.Commit().Error
}

func GetPeriodSales(startTime, endTime string, pageNum, pageSize int) ([]DrugSaleOrder, error) {
	var drugS []DrugSaleOrder
	err := db.Where("sale_date BETWEEN ? AND ?", startTime, endTime).Offset(pageNum).Limit(pageSize).Find(&drugS).Error
	if err != nil {
		return nil, err
	}
	return drugS, nil
}

func GetPeriodSalesByEmployeeID(startTime, endTime string, pageNum, pageSize, eid int) ([]DrugSaleOrder, error) {
	var drugS []DrugSaleOrder
	err := db.Where("employee_id = ? AND sale_date BETWEEN ? AND ?", eid, startTime, endTime).Offset(pageNum).Limit(pageSize).Find(&drugS).Error
	if err != nil {
		return nil, err
	}
	return drugS, nil
}

func GetPeriodSalesByCustomerID(startTime, endTime string, pageNum, pageSize, cid int) ([]DrugSaleOrder, error) {
	var drugS []DrugSaleOrder
	err := db.Where("customer_id = ? AND sale_date BETWEEN ? AND ?", cid, startTime, endTime).Offset(pageNum).Limit(pageSize).Find(&drugS).Error
	if err != nil {
		return nil, err
	}
	return drugS, nil
}

// 获取某时间段的订单总数
func GetPeriodSalesTotal(startTime, endTime string) (int, error) {
	var count int
	err := db.Model(DrugSaleOrder{}).Where("sale_date BETWEEN ? AND ?", startTime, endTime).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 获取某时间段的订单总数
func GetEmployeePeriodSalesTotal(startTime, endTime string, eid int) (int, error) {
	var count int
	err := db.Model(DrugSaleOrder{}).Where("employee_id = ? AND sale_date BETWEEN ? AND ?",eid, startTime, endTime).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 获取某时间段的订单总数
func GetCustomerPeriodSalesTotal(startTime, endTime string, cid int) (int, error) {
	var count int
	err := db.Model(DrugSaleOrder{}).Where("customer_id = ? AND sale_date BETWEEN ? AND ?",cid, startTime, endTime).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetTotalSales() (int, error) {
	var orderList []SaleDrugTotalPrice
	totalSales := 0
	err := db.Model(SaleDrugTotalPrice{}).Find(&orderList).Error
	if err != nil {
		return 0, err
	}

	for _, drug := range orderList{
		totalSales += drug.TotalPrice
	}

	return totalSales, nil
}

func GetDrugSaleOrderByID(id int)(*DrugSaleOrder, error){
	var order DrugSaleOrder
	err := db.Where("drug_sale_order_id = ? ", id).First(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}