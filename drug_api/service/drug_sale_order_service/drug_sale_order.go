package drug_sale_order_service

import (
	"drug_api/models"
	"time"
)

type DrugSaleOrder struct {
	DrugSaleOrderID    int
	EmployeeID         int
	SaleDate           time.Time
	CustomerID         int
	Drugs              []models.InvDrugWithNum
	StartTime          string
	EndTime            string
	PeriodSalePageNum  int
	PeriodSalePageSize int
	DetailPageNum      int
	DetailPageSize     int
}

func (dso *DrugSaleOrder) AddDrugSaleOrder() error {
	return models.AddDrugSaleOrder(dso.EmployeeID, dso.CustomerID, dso.SaleDate, dso.Drugs)
}

type PeriodSalesMap map[string]interface{}

func (dso *DrugSaleOrder) GetPeriodSales() ([]PeriodSalesMap, error) {
	var periodSales []PeriodSalesMap

	drugs, err := models.GetPeriodSales(dso.StartTime, dso.EndTime, dso.PeriodSalePageNum, dso.PeriodSalePageSize)
	if err != nil {
		return nil, err
	}
	for _, each := range drugs {
		periodSalesMap := make(PeriodSalesMap)
		// 1. 销售订单的id
		periodSalesMap["drug_sale_order_id"] = each.DrugSaleOrderID
		// 2. 销售员信息
		employee, err := models.GetEmployeeByID(each.EmployeeID)
		if err != nil {
			return nil, err
		}
		periodSalesMap["employee"] = employee
		// 3. 售出时间
		periodSalesMap["sale_date"] = each.SaleDate
		// 4. 顾客信息
		customer, err := models.GetCustomerByID(each.CustomerID)
		if err != nil {
			return nil, err
		}
		periodSalesMap["customer"] = customer

		totalSale, totalProfit, err := models.GetTotalSalePriceAndProfitByDrugSaleOrderID(each.DrugSaleOrderID)
		if err != nil {
			return nil, err
		}
		// 5. 这单总价
		periodSalesMap["totalSale"] = totalSale
		// 6. 这单总利润
		periodSalesMap["totalProfit"] = totalProfit

		totalNum, err := models.GetTotalNumByDrugSaleOrderID(each.DrugSaleOrderID)
		if err != nil {
			return nil, err
		}

		periodSalesMap["totalNum"] = totalNum

		periodSales = append(periodSales, periodSalesMap)
	}

	return periodSales, nil
}

func (dso *DrugSaleOrder) CountPeriodSales() (int, error) {
	return models.GetPeriodSalesTotal(dso.StartTime, dso.EndTime)
}

func (dso *DrugSaleOrder) GetTotalSales() (int, error) {
	return models.GetTotalSales()
}

func (dso *DrugSaleOrder) GetDetailSaleOrder() (interface{}, error) {
	ret := make(map[string]interface{})

	// 获取员工信息,客户信息和售出时间
	saleOrder, err := models.GetDrugSaleOrderByID(dso.DrugSaleOrderID)
	if err != nil {
		return nil, err
	}
	ret["sale_date"] = saleOrder.SaleDate
	employee, err := models.GetEmployeeByID(saleOrder.EmployeeID)
	if err != nil {
		return nil, err
	}
	ret["employee"] = employee
	customer, err := models.GetCustomerByID(saleOrder.CustomerID)
	if err != nil {
		return nil, err
	}
	ret["customer"] = customer

	// 获取该订单的库存id和数量
	drugs, err := models.GetSaleDrugListsByDrugSaleOrderID(dso.DetailPageNum, dso.DetailPageSize, dso.DrugSaleOrderID)
	if err != nil {
		return nil, err
	}

	// 获得药品个数
	totalDrugs, err := models.GetSaleDrugListCountByDrugSaleOrderID(dso.DrugSaleOrderID)
	if err != nil {
		return nil, err
	}
	ret["totalDrugs"] = totalDrugs

	// 构造一个药品列表, 包含的信息如下
	// drug => drug 对象
	// sale_price => 售价
	// supplier => 供应商
	// num => 数量
	drugList := make([]map[string]interface{}, 0)
	for _, v := range drugs {
		drug := make(map[string]interface{})
		drug["num"] = v.Num
		// 通过inventory_drug_id获取inventory_drug
		invDrug, err := models.GetInvDrugByID(v.InventoryDrugID)
		if err != nil {
			return nil, err
		}
		drug["sale_price"] = invDrug.SalePrice
		drug["buy_price"] = invDrug.PurchasePrice
		supplier, err := models.GetSupplierByID(invDrug.SupplierID)
		if err != nil {
			return nil, err
		}
		drug["supplier"] = supplier
		d, err := models.GetDrugBuyID(invDrug.DrugID)
		if err != nil {
			return nil, err
		}
		drug["drug"] = d

		drugList = append(drugList, drug)
	}

	ret["drug_list"] = drugList

	return ret, nil
}

func (dso *DrugSaleOrder) GetCustomerDetailSaleOrder() (interface{}, error) {
	ret := make(map[string]interface{})

	// 获取员工信息,客户信息和售出时间
	saleOrder, err := models.GetDrugSaleOrderByID(dso.DrugSaleOrderID)
	if err != nil {
		return nil, err
	}
	ret["sale_date"] = saleOrder.SaleDate
	employee, err := models.GetEmployeeByID(saleOrder.EmployeeID)
	if err != nil {
		return nil, err
	}
	ret["employee"] = employee
	customer, err := models.GetCustomerByID(saleOrder.CustomerID)
	if err != nil {
		return nil, err
	}
	ret["customer"] = customer

	// 获取该订单的库存id和数量
	drugs, err := models.GetSaleDrugListsByDrugSaleOrderID(dso.DetailPageNum, dso.DetailPageSize, dso.DrugSaleOrderID)
	if err != nil {
		return nil, err
	}

	// 获得药品个数
	totalDrugs, err := models.GetSaleDrugListCountByDrugSaleOrderID(dso.DrugSaleOrderID)
	if err != nil {
		return nil, err
	}
	ret["totalDrugs"] = totalDrugs

	// 构造一个药品列表, 包含的信息如下
	// drug => drug 对象
	// sale_price => 售价
	// supplier => 供应商
	// num => 数量
	drugList := make([]map[string]interface{}, 0)
	for _, v := range drugs {
		drug := make(map[string]interface{})
		drug["num"] = v.Num
		// 通过inventory_drug_id获取inventory_drug
		invDrug, err := models.GetInvDrugByID(v.InventoryDrugID)
		if err != nil {
			return nil, err
		}
		drug["sale_price"] = invDrug.SalePrice
		supplier, err := models.GetSupplierByID(invDrug.SupplierID)
		if err != nil {
			return nil, err
		}
		drug["supplier"] = supplier
		d, err := models.GetDrugBuyID(invDrug.DrugID)
		if err != nil {
			return nil, err
		}
		drug["drug"] = d

		drugList = append(drugList, drug)
	}

	ret["drug_list"] = drugList

	return ret, nil
}