package v1

import (
	"drug_api/pkg/app"
	"drug_api/pkg/e"
	"drug_api/pkg/setting"
	"drug_api/pkg/util"
	"drug_api/service/customer_service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// 获取某个员工在某段时间内的销售信息
func GetCustomerSaleInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	cid := com.StrTo(c.Param("cid")).MustInt()
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")

	customerService := customer_service.Customer{
		CustomerID:         cid,
		StartTime:          startTime,
		EndTime:            endTime,
		SaleDetailPageNum:  util.GetPage(c),
		SaleDetailPageSize: setting.AppSetting.PageSize,
	}

	orders , err := customerService.GetCustomerSaleInfo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CUSTOMER_SALES_FAILED ,nil)
		return
	}

	count ,err := customerService.CountCustomerPeriodSales()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CUSTOMER_SALES_COUNT_FAILED ,nil)
		return
	}

	data := make(map[string]interface{})
	data["orders"] = orders
	data["count"] = count

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func EditCustomerPassword(c *gin.Context){
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	newPassword := c.PostForm("new_password")
	confirmPassword := c.PostForm("confirm_password")

	if newPassword != confirmPassword{
		appG.Response(http.StatusInternalServerError, e.ERROR_DIFF_PASSWORD ,nil)
		return
	}

	customerService := customer_service.Customer{
		CustomerID: id,
		NewPassword: newPassword,
		ConfirmPassword: confirmPassword,
	}

	err := customerService.EditCustomerPassword()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_PASSWORD_FAILED ,nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func AddCustomer(c *gin.Context){
	appG := app.Gin{C: c}
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")
	name := c.PostForm("name")

	if password != confirmPassword{
		appG.Response(http.StatusInternalServerError, e.ERROR_DIFF_PASSWORD ,nil)
		return
	}

	customerService := customer_service.Customer{
		Name: name,
		Username: phone,
		Password: password,
		Phone: phone,
	}

	// 检查号码是否重复
	flag, err := customerService.ExistByPhone()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_FAILED ,nil)
		return
	}

	if flag{
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_USERNAME ,nil)
		return
	}

	id ,err := customerService.AddCustomer()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NEW_MANAGER_FAILED,nil)
		return
	}

	data := make(map[string]interface{})
	data["id"] = id

	appG.Response(http.StatusOK, e.SUCCESS, data)
}