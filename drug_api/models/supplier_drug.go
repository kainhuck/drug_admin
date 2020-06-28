package models

type SupplierDrug struct {
	SupplierDrugID int `gorm:"primary_key" json:"supplier_drug_id"`
	SupplierID     int `json:"supplier_id"`
	DrugID         int `json:"drug_id"`
	SalePrice      int `json:"sale_price"`
}

type SupplierDrugWithName struct {
	SupplierDrugID int    `gorm:"primary_key" json:"supplier_drug_id"`
	SupplierID     int    `json:"supplier_id"`
	DrugID         int    `json:"drug_id"`
	SalePrice      int    `json:"sale_price"`
	Cname           string `json:"cname"`
}

type DrugWithSalePrice struct {
	Drug      Drug
	SalePrice int `json:"sale_price"`
}

func GetDrugsBuySupplierID(sid, pageNum, pageSize int) ([]*DrugWithSalePrice, error) {
	drugWithSalePrices := make([]*DrugWithSalePrice, 0)

	// 1. 先从supplier_drug表里查找所有记录
	var supplierDrugs []*SupplierDrug
	err := db.Where("supplier_id = ?", sid).Offset(pageNum).Limit(pageSize).Find(&supplierDrugs).Error
	if err != nil {
		return nil, err
	}

	// 2. 拿到drug_id后从drug表里查找drug
	for _, v := range supplierDrugs {
		var drug Drug
		err = db.Where("drug_id = ?", v.DrugID).First(&drug).Error
		if err != nil {
			return nil, err
		}

		// 3. drugWithSalePrices
		drugWithSalePrices = append(drugWithSalePrices, &DrugWithSalePrice{
			Drug:      drug,
			SalePrice: v.SalePrice,
		})

	}

	return drugWithSalePrices, nil
}

func GetSearchDrugsBuySupplierID(sid, pageNum, pageSize int, searchContent string) ([]*DrugWithSalePrice, error) {
	drugWithSalePrices := make([]*DrugWithSalePrice, 0)

	// 1. 先从supplier_drug_with_name表里查找所有记录
	var supplierDrugs []*SupplierDrugWithName
	err := db.Where("supplier_id = ? AND cname LIKE ?", sid, "%"+searchContent+"%").Offset(pageNum).Limit(pageSize).Find(&supplierDrugs).Error
	if err != nil {
		return nil, err
	}

	// 2. 拿到drug_id后从drug表里查找drug
	for _, v := range supplierDrugs {
		var drug Drug
		err = db.Where("drug_id = ?", v.DrugID).First(&drug).Error
		if err != nil {
			return nil, err
		}

		// 3. drugWithSalePrices
		drugWithSalePrices = append(drugWithSalePrices, &DrugWithSalePrice{
			Drug:      drug,
			SalePrice: v.SalePrice,
		})

	}

	return drugWithSalePrices, nil
}

func GetSupplierDrugByIDs(did, sid int) (*SupplierDrug, error) {
	var sd SupplierDrug
	err := db.Where("drug_id = ? and supplier_id = ?", did, sid).First(&sd).Error
	if err != nil {
		return nil, err
	}

	return &sd, nil
}
