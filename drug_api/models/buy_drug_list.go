package models

type BuyDrugList struct {
	BuyDrugListID  int `gorm:"primary_key" json:"buy_drug_list_id"`
	DrugID         int `json:"drug_id"`
	Num            int `json:"num"`
	DrugBuyOrderID int `json:"drug_buy_order_id"`
}

func GetTotalPriceByDrugBuyOrderID(id int, sid int)(int, error){
	var drugs []BuyDrugList
	totalPrice := 0
	// 1.通过这个id获取所有drug_id
	err := db.Where("drug_buy_order_id = ?", id).Find(&drugs).Error
	if err != nil {
		return 0, err
	}

	for _, each := range drugs{
		// 2. 通过这个药品id和sid查找售价
		sDrug, err := GetSupplierDrugByIDs(each.DrugID, sid)
		if err != nil {
			return 0, err
		}
		totalPrice += each.Num * sDrug.SalePrice
	}
	return totalPrice, nil
}

func GetBuyDrugListsByDrugBuyOrderID(pageNum, pageSize, id int)([]*BuyDrugList, error){
	var drugs []*BuyDrugList
	err := db.Where("drug_buy_order_id = ?", id).Offset(pageNum).Limit(pageSize).Find(&drugs).Error
	if err != nil {
		return nil, err
	}

	return drugs, nil
}

func GetBuyDrugListCountByDrugBuyOrderID(id int)(int, error){
	var count int
	err := db.Model(BuyDrugList{}).Where("drug_buy_order_id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}