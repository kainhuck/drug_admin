package customer_service

import "drug_api/models"

type Customer struct {
	CustomerID         int    `json:"customer_id"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	Name               string `json:"name"`
	Phone              string `json:"phone"`
	SaleDetailPageNum  int
	SaleDetailPageSize int
	StartTime          string
	EndTime            string
	NewPassword        string
	ConfirmPassword    string
	CustomersPageNum   int
	CustomersPageSize  int
}

func (c *Customer) Check() (int, bool, error) {
	return models.CheckCAuth(c.Username, c.Password)
}

type PeriodSalesMap map[string]interface{}

// 查找某段时间内的顾客订单
func (c *Customer) GetCustomerSaleInfo() ([]PeriodSalesMap, error) {
	// 1. 先获得该顾客在这段时间内的该页的销售订单id列表
	// 2. 根据这个订单ID列表查找以下数据
	//    售出订单的id, 销售员信息, 售出时间, 顾客信息 这一单的总价 这一单的总件数
	var periodSales []PeriodSalesMap

	drugs, err := models.GetPeriodSalesByCustomerID(c.StartTime, c.EndTime, c.SaleDetailPageNum, c.SaleDetailPageSize, c.CustomerID)
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

func (c *Customer) CountCustomerPeriodSales() (int, error) {
	return models.GetCustomerPeriodSalesTotal(c.StartTime, c.EndTime, c.CustomerID)
}

func (c *Customer) EditCustomerPassword() error {
	return models.EditCustomerPassword(c.CustomerID, c.NewPassword)
}

func (c *Customer) AddCustomer() (int, error) {
	return models.AddCustomer(c.Username, c.Password, c.Name)
}

func (c *Customer) ExistByPhone() (bool, error) {
	return models.ExistByPhone(c.Phone)
}

func (c *Customer) GetCustomers() ([]*models.Customer, error) {
	return models.GetCustomers(c.CustomersPageNum, c.CustomersPageSize)
}

func (c *Customer) CountCustomers() (int, error) {
	return models.CountCustomers()
}

func (c *Customer) GetCustomerByID() (*models.Customer, error) {
	return models.GetCustomerByID(c.CustomerID)
}

func (c *Customer) GetCustomerPeriodSalesTotalPrice() (int ,error) {
	return models.GetCustomerPeriodSalesTotalPrice(c.StartTime, c.EndTime, c.CustomerID)
}