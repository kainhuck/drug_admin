package employee_service

import "drug_api/models"

type Employee struct {
	EmployeeID         int    `json:"employee_id"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	Name               string `json:"name"`
	Position           string `json:"position"`
	SaleDetailPageNum  int
	SaleDetailPageSize int
	StartTime          string
	EndTime            string
	NewPassword        string
	ConfirmPassword    string
	EmployeesPageNum   int
	EmployeesPageSize  int
	NewPosition        string
}

func (e *Employee) Check() (int, bool, error) {
	return models.CheckEAuth(e.Username, e.Password)
}

type PeriodSalesMap map[string]interface{}

// 查找某段时间内的员工业绩
func (e *Employee) GetEmployeeSaleInfo() ([]PeriodSalesMap, error) {
	// 1. 先获得该员工在这段时间内的该页的销售订单id列表
	// 2. 根据这个订单ID列表查找以下数据
	//    售出订单的id, 销售员信息, 售出时间, 顾客信息 这一单的总价 这一单的总利润, 这一单的总件数
	var periodSales []PeriodSalesMap

	drugs, err := models.GetPeriodSalesByEmployeeID(e.StartTime, e.EndTime, e.SaleDetailPageNum, e.SaleDetailPageSize, e.EmployeeID)
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

func (e *Employee) CountEmployeePeriodSales() (int, error) {
	return models.GetEmployeePeriodSalesTotal(e.StartTime, e.EndTime, e.EmployeeID)
}

func (e *Employee) GetEmployeePeriodSalesTotalPrice() (int, error) {
	return models.GetEmployeePeriodSalesTotalPrice(e.StartTime, e.EndTime, e.EmployeeID)
}

func (e *Employee) EditEmployeePassword() error {
	return models.EditEmployeePassword(e.EmployeeID, e.NewPassword)
}

func (e *Employee) EditEmployeePosition() error {
	return models.EditEmployeePosition(e.EmployeeID, e.NewPosition)
}

func (e *Employee) AddEmployee() (int, error) {
	return models.AddEmployee(e.Username, e.Password, e.Name, e.Position)
}

func (e *Employee) ExistByUsername() (bool, error) {
	return models.ExistByUsername(e.Username)
}

func (e *Employee) GetEmployees() ([]*models.Employee, error) {
	return models.GetEmployees(e.EmployeesPageNum, e.EmployeesPageSize)
}

func (e *Employee) GetAllEmployees () ([]*models.Employee, error) {
	return models.GetAllEmployees()
}

func (e *Employee) CountEmployees() (int, error) {
	return models.GetTotalEmployees()
}

func (e *Employee) GetEmployeeByID() (*models.Employee, error) {
	return models.GetEmployeeByID(e.EmployeeID)
}