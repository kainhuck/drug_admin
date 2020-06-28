package v1

import (
	"drug_api/pkg/app"
	"drug_api/pkg/e"
	"drug_api/pkg/setting"
	"drug_api/pkg/util"
	"drug_api/service/inventory_drug_service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetAllInvDrugs(c *gin.Context) {
	appG := app.Gin{C: c}

	inventoryDrugService := inventory_drug_service.InventoryDrug{
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	invDrugs, err := inventoryDrugService.GetAllInvDrugs()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_INVDRUGS_FAILED, nil)
		return
	}

	count, err := inventoryDrugService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_INVDRUGS_COUNT_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["count"] = count
	data["invDrugs"] = invDrugs

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func SearchAllInvDrugs(c *gin.Context) {
	appG := app.Gin{C: c}
	searchContent := c.Query("search_content")

	inventoryDrugService := inventory_drug_service.InventoryDrug{
		PageNum:       util.GetPage(c),
		PageSize:      setting.AppSetting.PageSize,
		SearchContent: searchContent,
	}

	invDrugs, err := inventoryDrugService.SearchAllInvDrugs()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_INVDRUGS_FAILED, nil)
		return
	}

	count, err := inventoryDrugService.CountSearch()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_INVDRUGS_COUNT_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["count"] = count
	data["invDrugs"] = invDrugs

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func SearchAllInvDrugsCustomer(c *gin.Context) {
	appG := app.Gin{C: c}
	searchContent := c.Query("search_content")

	inventoryDrugService := inventory_drug_service.InventoryDrug{
		PageNum:       util.GetPage(c),
		PageSize:      setting.AppSetting.PageSize,
		SearchContent: searchContent,
	}

	invDrugs, err := inventoryDrugService.SearchAllInvDrugsCustomer()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_INVDRUGS_FAILED, nil)
		return
	}

	count, err := inventoryDrugService.CountSearch()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_INVDRUGS_COUNT_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["count"] = count
	data["invDrugs"] = invDrugs

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetAllInvDrugsCustomer(c *gin.Context) {
	appG := app.Gin{C: c}

	inventoryDrugService := inventory_drug_service.InventoryDrug{
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	invDrugs, err := inventoryDrugService.GetAllInvDrugsCustomer()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_INVDRUGS_FAILED, nil)
		return
	}

	count, err := inventoryDrugService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_INVDRUGS_COUNT_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["count"] = count
	data["invDrugs"] = invDrugs

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func EditInvDrugSalePrice(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	newSalePrice := com.StrTo(c.PostForm("new_sale_price")).MustInt()
	inventoryDrugService := inventory_drug_service.InventoryDrug{
		InventoryDrugID: id,
		NewSalePrice:    newSalePrice,
	}

	err := inventoryDrugService.EditInvDrugSalePrice()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_INVDRUG_SALE_PRICE_FAILED, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
