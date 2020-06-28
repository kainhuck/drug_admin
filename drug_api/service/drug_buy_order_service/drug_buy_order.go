package drug_buy_order_service

import (
	"drug_api/models"
	"time"
)

type DrugBuyOrder struct {
	DrugBuyOrderID    int
	ManagerID         int
	BuyDate           time.Time
	SupplierID        int
	Drugs             []models.DrugWithNum
	StartTime         string
	EndTime           string
	PeriodBuyPageNum  int
	PeriodBuyPageSize int
	DetailPageNum     int
	DetailPageSize    int
}

func (d *DrugBuyOrder) AddDrugBuyOrder() error {
	return models.AddDrugBuyOrder(d.ManagerID, d.SupplierID, d.BuyDate, d.Drugs)
}

type PeriodBuyMap map[string]interface{}

func (d *DrugBuyOrder) GetPeriodBuys() ([]PeriodBuyMap, error) {
	var periodbuy []PeriodBuyMap

	drugs, err := models.GetPeriodBuy(d.StartTime, d.EndTime, d.PeriodBuyPageNum, d.PeriodBuyPageSize)
	if err != nil {
		return nil, err
	}
	for _, each := range drugs {
		periodBuyMap := make(PeriodBuyMap)
		// 1. 销售订单的id
		periodBuyMap["drug_buy_order_id"] = each.DrugBuyOrderID
		// 2.经理信息
		manager, err := models.GetManagerByID(each.ManagerID)
		if err != nil {
			return nil, err
		}
		periodBuyMap["manager"] = manager
		// 3. 售出时间
		periodBuyMap["buy_date"] = each.BuyDate
		// 4. 供应商信息
		supplier, err := models.GetSupplierByID(each.SupplierID)
		if err != nil {
			return nil, err
		}
		periodBuyMap["supplier"] = supplier

		totalPrice, err := models.GetTotalPriceByDrugBuyOrderID(each.DrugBuyOrderID, each.SupplierID)
		if err != nil {
			return nil, err
		}
		// 5. 这单总价
		periodBuyMap["totalPrice"] = totalPrice

		periodbuy = append(periodbuy, periodBuyMap)
	}

	return periodbuy, nil
}

func (d *DrugBuyOrder) GetPeriodBuysCount() (int, error) {
	return models.GetPeriodBuyTotal(d.StartTime, d.EndTime)
}

func (d *DrugBuyOrder) GetTotalBuy() (int, error) {
	return models.GetTotalBuy()
}

func (d *DrugBuyOrder) GetDetailBuyOrder() (interface{}, error) {
	ret := make(map[string]interface{})

	// 获取经理信息,供应商信息和进货时间
	buyOrder, err := models.GetDrugBuyOrderByID(d.DrugBuyOrderID)
	if err != nil {
		return nil, err
	}
	ret["buy_date"] = buyOrder.BuyDate
	manager, err := models.GetManagerByID(buyOrder.ManagerID)
	if err != nil {
		return nil, err
	}
	ret["manager"] = manager
	supplier, err := models.GetSupplierByID(buyOrder.SupplierID)
	if err != nil {
		return nil, err
	}
	ret["supplier"] = supplier

	// 获取该订单的所有药品id和数量
	drugs, err := models.GetBuyDrugListsByDrugBuyOrderID(d.DetailPageNum, d.DetailPageSize, d.DrugBuyOrderID)
	if err != nil {
		return nil, err
	}

	totalDrugs, err := models.GetBuyDrugListCountByDrugBuyOrderID(d.DrugBuyOrderID)
	if err != nil {
		return nil, err
	}
	ret["totalDrugs"] = totalDrugs

	// 构造一个药品列表, 包含的信息如下
	// drug => drug 对象
	// buy_price => 进价
	// num => 数量
	drugList := make([]map[string]interface{}, 0)
	for _, v := range drugs {
		drug := make(map[string]interface{})
		drug["num"] = v.Num
		// 通过drug_id和supplier_id获取supplier_drug
		sDrug, err := models.GetSupplierDrugByIDs(v.DrugID, buyOrder.SupplierID)
		if err != nil {
			return nil, err
		}
		// 进价
		drug["buy_price"] = sDrug.SalePrice
		d, err := models.GetDrugBuyID(v.DrugID)
		if err != nil {
			return nil, err
		}
		drug["drug"] = d

		drugList = append(drugList, drug)
	}

	ret["drug_list"] = drugList

	return ret, nil
}

func (d *DrugBuyOrder) GetBuyOrderTotalPrice() (int, error) {
	return models.GetBuyOrderTotalPrice(d.DrugBuyOrderID)
}
