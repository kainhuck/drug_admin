package models

import "github.com/jinzhu/gorm"

type InventoryDrug struct {
	InventoryDrugID int `gorm:"primary_key" json:"inventory_drug_id"`
	DrugID          int `json:"drug_id"`
	PurchasePrice   int `json:"purchase_price"`
	SalePrice       int `json:"sale_price"`
	SupplierID      int `json:"supplier_id"`
	InventoryNum    int `json:"inventory_num"`
	Drug            *Drug
	Supplier        *Supplier
}

type InventoryDrugWithName struct {
	InventoryDrugID int    `gorm:"primary_key" json:"inventory_drug_id"`
	DrugID          int    `json:"drug_id"`
	PurchasePrice   int    `json:"purchase_price"`
	SalePrice       int    `json:"sale_price"`
	SupplierID      int    `json:"supplier_id"`
	InventoryNum    int    `json:"inventory_num"`
	Cname           string `json:"cname"`
	Drug            *Drug
	Supplier        *Supplier
}

// ExistInvDrugByID checks
func ExistInvDrugByID(id int) (bool, error) {
	var drug InventoryDrug
	err := db.Select("inventory_drug_id").Where("inventory_drug_id = ?", id).First(&drug).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if drug.InventoryDrugID > 0 {
		return true, nil
	}

	return false, nil
}

// GetInvDrugTotal gets
func GetInvDrugTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&InventoryDrug{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetSearchInvDrugTotal(searchContent string) (int, error) {
	var count int
	if err := db.Model(&InventoryDrugWithName{}).Where("cname LIKE ?", "%"+searchContent+"%").Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// 修改药品价格
func EditInvDrugSalePrice(id int, newSalePrice int) error {
	return db.Model(InventoryDrug{}).Where("inventory_drug_id = ?", id).Update("sale_price", newSalePrice).Error
}

func GetInvDrugs(pageNum int, pageSize int, maps interface{}) ([]*InventoryDrug, error) {
	var drugs []*InventoryDrug
	err := db.Where(maps).Order("inventory_num").Offset(pageNum).Limit(pageSize).Find(&drugs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	for _, each := range drugs {
		drug, err := GetDrugBuyID(each.DrugID)
		if err != nil {
			return nil, err
		}
		each.Drug = drug

		supplier, err := GetSupplierBriefInfo(each.SupplierID)
		if err != nil {
			return nil, err
		}
		each.Supplier = supplier
	}

	return drugs, nil
}

func SearchAllInvDrugs(pageNum int, pageSize int, searchContent string) ([]*InventoryDrugWithName, error) {
	var drugs []*InventoryDrugWithName
	err := db.Debug().Where("cname LIKE ?", "%"+searchContent+"%").Order("inventory_num").Offset(pageNum).Limit(pageSize).Find(&drugs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	for _, each := range drugs {
		drug, err := GetDrugBuyID(each.DrugID)
		if err != nil {
			return nil, err
		}
		each.Drug = drug

		supplier, err := GetSupplierBriefInfo(each.SupplierID)
		if err != nil {
			return nil, err
		}
		each.Supplier = supplier
	}

	return drugs, nil
}

func GetInvDrugsCustomer(pageNum int, pageSize int, maps interface{}) ([]*InventoryDrug, error) {
	var drugs []*InventoryDrug
	err := db.Where(maps).Order("inventory_num").Offset(pageNum).Limit(pageSize).Find(&drugs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	for _, each := range drugs {
		each.PurchasePrice = -1
		drug, err := GetDrugBuyID(each.DrugID)
		if err != nil {
			return nil, err
		}
		each.Drug = drug

		supplier, err := GetSupplierBriefInfo(each.SupplierID)
		if err != nil {
			return nil, err
		}
		each.Supplier = supplier
	}

	return drugs, nil
}

func GetInvDrugByID(id int) (*InventoryDrug, error) {
	var invDrug InventoryDrug
	err := db.Where("inventory_drug_id = ?", id).First(&invDrug).Error
	if err != nil {
		return nil, err
	}

	return &invDrug, nil
}
