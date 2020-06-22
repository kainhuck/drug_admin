package models

type SaleDrugList struct {
	SaleDrugListID  int `gorm:"primary_key" json:"sale_drug_list_id"`
	InventoryDrugID int `json:"inventory_drug_id"`
	Num             int `json:"num"`
	DrugSaleOrderID int `json:"drug_sale_order_id"`
}

type InvDrugWithNum map[string]int

func GetTotalSalePriceAndProfitByDrugSaleOrderID(dsoID int) (int, int, error) {
	var drugs []SaleDrugList
	err := db.Where("drug_sale_order_id = ?", dsoID).Find(&drugs).Error
	if err != nil {
		return 0, 0, err
	}
	totalSale := 0
	totalProfit := 0

	for _, drug := range drugs {
		// 查找当前库存药品id对应的详细信息
		invDrug, err := GetInvDrugByID(drug.InventoryDrugID)
		if err != nil {
			return 0, 0, err
		}
		totalSale += invDrug.SalePrice * drug.Num
		totalProfit += (invDrug.SalePrice - invDrug.PurchasePrice) * drug.Num
	}

	return totalSale, totalProfit, nil
}

func GetTotalNumByDrugSaleOrderID(id int) (int, error) {
	var sdls []SaleDrugList
	count := 0
	err := db.Where("drug_sale_order_id = ?", id).Find(&sdls).Error
	if err != nil {
		return 0, err
	}

	for _, v := range sdls {
		count += v.Num
	}

	return count, nil
}

func GetSaleDrugListsByDrugSaleOrderID(pageNum, pageSize, id int) ([]*SaleDrugList, error) {
	var sdls []*SaleDrugList

	err := db.Where("drug_sale_order_id = ?", id).Offset(pageNum).Limit(pageSize).Find(&sdls).Error
	if err != nil {
		return nil, err
	}

	return sdls, nil
}

func GetSaleDrugListCountByDrugSaleOrderID(id int)(int ,error){
	var count int
	err := db.Model(SaleDrugList{}).Where("drug_sale_order_id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}