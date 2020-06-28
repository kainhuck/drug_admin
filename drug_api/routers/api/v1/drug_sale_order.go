package v1

import (
	"drug_api/models"
	"drug_api/pkg/app"
	"drug_api/pkg/e"
	"drug_api/pkg/setting"
	"drug_api/pkg/util"
	"drug_api/service/drug_sale_order_service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"time"
)

func AddDrugSaleOrder(c *gin.Context){
	appG := app.Gin{C: c}
	cid := com.StrTo(c.PostForm("customer_id")).MustInt()
	eid := com.StrTo(c.PostForm("employee_id")).MustInt()
	// 获得 库存药品 id 和 对应数量的列表, 前台传一个json序列化字符串
	// 格式为  [{InventoryDrugID: 1212, Num: 2}, {InventoryDrugID: 231231, Num: 3}]
	var drugs []models.InvDrugWithNum
	drugStr := c.PostForm("drugs")
	err := json.Unmarshal([]byte(drugStr), &drugs)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	drugSaleOrderService := drug_sale_order_service.DrugSaleOrder{
		CustomerID: cid,
		EmployeeID: eid,
		SaleDate: time.Now(),
		Drugs: drugs,
	}

	err = drugSaleOrderService.AddDrugSaleOrder()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_DRUG_SALE_ORDER_FAILED, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// 获取某段时间内的销售订单
func GetPeriodSales(c *gin.Context){
	appG := app.Gin{C: c}
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")

	drugSaleOrderService := drug_sale_order_service.DrugSaleOrder{
		StartTime: startTime,
		EndTime: endTime,
		PeriodSalePageNum: util.GetPage(c),
		PeriodSalePageSize: setting.AppSetting.PageSize,
	}

	periodSales, err := drugSaleOrderService.GetPeriodSales()
	if err != nil {
		appG.Response(http.StatusInternalServerError,e.ERROR_GET_PERIOD_SALES_ORDER_FAILED ,nil)
		return
	}

	count, err := drugSaleOrderService.CountPeriodSales()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_PERIOD_SALES_ORDER_COUNT_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["period_sales"] = periodSales
	data["count"] = count

	appG.Response(http.StatusOK, e.SUCCESS, data)
}


func GetTotalSales(c *gin.Context){
	appG := app.Gin{C: c}
	drugSaleOrderService := drug_sale_order_service.DrugSaleOrder{}
	totalSales , err := drugSaleOrderService.GetTotalSales()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TOTAL_SALES_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["totalSales"] = totalSales

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetTotalProfit(c *gin.Context){
	appG := app.Gin{C: c}
	drugSaleOrderService := drug_sale_order_service.DrugSaleOrder{}
	totalProfit , err := drugSaleOrderService.GetTotalProfit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TOTAL_PROFIT_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["totalProfit"] = totalProfit

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetDetailSaleOrder(c *gin.Context){
	appG := app.Gin{C: c}
	// 获取订单的id
	sid := com.StrTo(c.Param("sid")).MustInt()

	drugSaleOrderService := drug_sale_order_service.DrugSaleOrder{
		DrugSaleOrderID: sid,
		DetailPageNum: util.GetPage(c),
		DetailPageSize: setting.AppSetting.PageSize,
	}

	order, err := drugSaleOrderService.GetDetailSaleOrder()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_DETAIL_SALE_ORDER_FAILED ,nil)
		return
	}

	totalPrice, err := drugSaleOrderService.GetSaleOrderTotalPrice()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_DETAIL_SALE_ORDER_FAILED ,nil)
		return
	}

	data := make(map[string]interface{})
	data["saleOrder"] = order
	data["totalPrice"] = totalPrice

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetCustomerDetailSaleOrder(c *gin.Context){
	appG := app.Gin{C: c}
	// 获取订单的id
	sid := com.StrTo(c.Param("sid")).MustInt()

	drugSaleOrderService := drug_sale_order_service.DrugSaleOrder{
		DrugSaleOrderID: sid,
		DetailPageNum: util.GetPage(c),
		DetailPageSize: setting.AppSetting.PageSize,
	}

	order, err := drugSaleOrderService.GetCustomerDetailSaleOrder()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_DETAIL_SALE_ORDER_FAILED ,nil)
		return
	}

	totalPrice, err := drugSaleOrderService.GetSaleOrderTotalPrice()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_DETAIL_SALE_ORDER_FAILED ,nil)
		return
	}

	data := make(map[string]interface{})
	data["saleOrder"] = order
	data["totalPrice"] = totalPrice

	appG.Response(http.StatusOK, e.SUCCESS, data)
}