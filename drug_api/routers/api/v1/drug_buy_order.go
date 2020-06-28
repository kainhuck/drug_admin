package v1

import (
	"drug_api/models"
	"drug_api/pkg/app"
	"drug_api/pkg/e"
	"drug_api/pkg/setting"
	"drug_api/pkg/util"
	"drug_api/service/drug_buy_order_service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"time"
)

func AddDrugBuyOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	mid := com.StrTo(c.PostForm("manager_id")).MustInt()
	sid := com.StrTo(c.PostForm("supplier_id")).MustInt()
	drugStr := c.PostForm("drugs")
	var drugs []models.DrugWithNum
	err := json.Unmarshal([]byte(drugStr), &drugs)
	fmt.Println(err)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	drugBuyOrderService := drug_buy_order_service.DrugBuyOrder{
		ManagerID: mid,
		SupplierID: sid,
		Drugs: drugs,
		BuyDate: time.Now(),
	}
	err = drugBuyOrderService.AddDrugBuyOrder()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_DRUG_BUY_ORDER_FAILED, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetPeriodBuy(c *gin.Context){
	appG := app.Gin{C: c}
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")

	drugBuyOrderService := drug_buy_order_service.DrugBuyOrder{
		StartTime: startTime,
		EndTime: endTime,
		PeriodBuyPageNum: util.GetPage(c),
		PeriodBuyPageSize: setting.AppSetting.PageSize,
	}

	periodBuys , err := drugBuyOrderService.GetPeriodBuys()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_PERIOD_BUY_ORDER_FAILED, nil)
		return
	}

	count, err := drugBuyOrderService.GetPeriodBuysCount()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_PERIOD_BUY_ORDER_COUNT_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["periodBuys"] = periodBuys
	data["count"] = count

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetTotalBuy(c *gin.Context) {
	appG := app.Gin{C:c}
	drugBuyOrderService := drug_buy_order_service.DrugBuyOrder{}

	totalBuy, err := drugBuyOrderService.GetTotalBuy()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TOTAL_BUY_FAILED ,nil)
		return
	}

	data := make(map[string]interface{})
	data["totalBuy"] = totalBuy

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetDetailBuyOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	bid := com.StrTo(c.Param("bid")).MustInt()

	drugBuyOrderService := drug_buy_order_service.DrugBuyOrder{
		DrugBuyOrderID: bid,
		DetailPageNum: util.GetPage(c),
		DetailPageSize: setting.AppSetting.PageSize,
	}

	order, err := drugBuyOrderService.GetDetailBuyOrder()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_DETAIL_BUY_ORDER_FAILED ,nil)
		return
	}

	totalPrice, err := drugBuyOrderService.GetBuyOrderTotalPrice()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_DETAIL_BUY_ORDER_FAILED ,nil)
		return
	}

	data := make(map[string]interface{})
	data["buyOrder"] = order
	data["totalPrice"] = totalPrice

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
