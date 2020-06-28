package models

import "github.com/jinzhu/gorm"

type Supplier struct {
	SupplierID int    `gorm:"primary_key" json:"supplier_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Drugs      []*DrugWithSalePrice
	TotalDrugs int `json:"total_drugs"`
}

func ExistSupplierByID(id int) (bool, error) {
	var su Supplier
	err := db.Select("supplier_id").Where("supplier_id = ? ", id).First(&su).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if su.SupplierID > 0 {
		return true, nil
	}

	return false, nil
}

// GetSupplierTotal 获取供应商总数
func GetSupplierTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Supplier{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetSearchSupplierTotal(searchContent string) (int, error){
	var count int
	if err:=db.Debug().Model(&Supplier{}).Where("name LIKE ?", "%"+searchContent+"%").Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// 获取一页供应商信息,不含该供应商卖的药品
func GetSuppliersInfo(pageNum int, pageSize int, maps interface{}) ([]*Supplier, error) {
	var sus []*Supplier
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&sus).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return sus, nil
}

func GetAllSearchSuppliers(pageNum int, pageSize int, searchContent string) ([]*Supplier, error) {
	var sus []*Supplier
	err := db.Where("name LIKE ?", "%"+searchContent+"%").Offset(pageNum).Limit(pageSize).Find(&sus).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return sus, nil
}

// 获取指定供应商详细信息包括卖的药品种类
func GetSupplierInfo(id, pageNum, pageSize int) (*Supplier, error) {
	var su Supplier
	err := db.Where("supplier_id = ?", id).First(&su).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	drugs, err := GetDrugsBuySupplierID(id, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	su.Drugs = drugs

	var count int
	if err := db.Model(&SupplierDrug{}).Where("supplier_id = ?", id).Count(&count).Error; err != nil {
		return nil, err
	}

	su.TotalDrugs = count
	return &su, nil
}

func GetSearchSupplierInfo(id, pageNum, pageSize int, searchContent string)(*Supplier, error) {
	var su Supplier
	err := db.Where("supplier_id = ?", id).First(&su).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	drugs, err := GetSearchDrugsBuySupplierID(id, pageNum, pageSize, searchContent)
	if err != nil {
		return nil, err
	}
	su.Drugs = drugs

	var count int
	if err := db.Model(&SupplierDrugWithName{}).Where("supplier_id = ? AND cname LIKE ?", id, "%"+searchContent+"%").Count(&count).Error; err != nil {
		return nil, err
	}

	su.TotalDrugs = count
	return &su, nil
}

func GetSupplierBriefInfo(id int) (*Supplier, error) {
	var su Supplier
	err := db.Where("supplier_id = ?", id).First(&su).Error
	if err != nil {
		return nil, err
	}
	return &su, nil
}

func GetSupplierByID(id int) (*Supplier, error) {
	var supplier Supplier
	err := db.Where("Supplier_id = ?", id).First(&supplier).Error
	if err != nil {
		return nil, err
	}

	return &supplier, nil
}
