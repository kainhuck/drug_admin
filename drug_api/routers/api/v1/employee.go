package v1

import (
	"drug_api/pkg/app"
	"drug_api/pkg/e"
	"drug_api/pkg/setting"
	"drug_api/pkg/util"
	"drug_api/service/employee_service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// 获取某个员工在某段时间内的销售信息
func GetEmployeeSaleInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	eid := com.StrTo(c.Param("eid")).MustInt()
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")

	employeeService := employee_service.Employee{
		EmployeeID:         eid,
		StartTime:          startTime,
		EndTime:            endTime,
		SaleDetailPageNum:  util.GetPage(c),
		SaleDetailPageSize: setting.AppSetting.PageSize,
	}

	orders , err := employeeService.GetEmployeeSaleInfo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EMPLOYEE_SALES_FAILED ,nil)
		return
	}

	count ,err := employeeService.CountEmployeePeriodSales()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EMPLOYEE_SALES_COUNT_FAILED, nil)
		return
	}

	employee, err := employeeService.GetEmployeeByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EMLOYEE_FAILED, nil)
		return
	}

	totalPrice, err := employeeService.GetEmployeePeriodSalesTotalPrice()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EMPLOYEE_TOTAL_PRICE_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["orders"] = orders
	data["count"] = count
	data["employee"] = employee
	data["totalPrice"] = totalPrice

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// 修改员工职称
func EditEmployeePosition(c *gin.Context){
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	newPosition := c.PostForm("new_position")

	employeeService := employee_service.Employee{
		EmployeeID: id,
		NewPosition: newPosition,
	}

	err := employeeService.EditEmployeePosition()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_POSITION_FAILED ,nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func EditEmployeePassword(c *gin.Context){
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	newPassword := c.PostForm("new_password")
	confirmPassword := c.PostForm("confirm_password")

	if newPassword != confirmPassword{
		appG.Response(http.StatusInternalServerError, e.ERROR_DIFF_PASSWORD ,nil)
		return
	}

	employeeService := employee_service.Employee{
		EmployeeID: id,
		NewPassword: newPassword,
		ConfirmPassword: confirmPassword,
	}

	err := employeeService.EditEmployeePassword()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_PASSWORD_FAILED ,nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func AddEmployee(c *gin.Context) {
	appG := app.Gin{C: c}
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")
	name := c.PostForm("name")
	position := c.PostForm("position")

	if password != confirmPassword{
		appG.Response(http.StatusInternalServerError, e.ERROR_DIFF_PASSWORD ,nil)
		return
	}

	employeeService := employee_service.Employee{
		Name: name,
		Username: username,
		Password: password,
		Position: position,
	}

	flag, err := employeeService.ExistByUsername()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_FAILED ,nil)
		return
	}

	if flag{
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_USERNAME ,nil)
		return
	}

	id ,err := employeeService.AddEmployee()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NEW_MANAGER_FAILED,nil)
		return
	}

	data := make(map[string]interface{})
	data["id"] = id

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetEmployees(c *gin.Context) {
	appG := app.Gin{C: c}

	employeeService := employee_service.Employee{
		EmployeesPageNum: util.GetPage(c),
		EmployeesPageSize: setting.AppSetting.PageSize,
	}

	employees, err := employeeService.GetEmployees()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EMPLOYEES_FAILED, nil)
		return
	}

	count, err := employeeService.CountEmployees()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EMPLOYEES_COUNT_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["employees"] = employees
	data["count"] = count

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetAllEmployees(c *gin.Context) {
	appG := app.Gin{C: c}

	employeeService := employee_service.Employee{}

	employees, err := employeeService.GetAllEmployees()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EMPLOYEES_FAILED, nil)
		return
	}


	data := make(map[string]interface{})
	data["employees"] = employees

	appG.Response(http.StatusOK, e.SUCCESS, data)
}