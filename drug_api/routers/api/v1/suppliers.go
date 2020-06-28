package v1

import (
	"drug_api/pkg/app"
	"drug_api/pkg/e"
	"drug_api/pkg/setting"
	"drug_api/pkg/util"
	"drug_api/service/supplier_service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// GetAllSuppliers 获取所有supplier
func GetAllSuppliers(c *gin.Context) {
	appG := app.Gin{C: c}

	supplierService := supplier_service.Supplier{
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	suppliers, err := supplierService.GetAllSupplier()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_SUPPLIERS_FAILED, nil)
		return
	}

	count, err := supplierService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_SUPPLIERS_COUNT_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["suppliers"] = suppliers
	data["count"] = count

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetAllSearchSuppliers(c *gin.Context) {
	appG := app.Gin{C: c}
	searchContent := c.Query("search_content")

	supplierService := supplier_service.Supplier{
		PageNum:       util.GetPage(c),
		PageSize:      setting.AppSetting.PageSize,
		SearchContent: searchContent,
	}

	suppliers, err := supplierService.GetAllSearchSuppliers()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_SUPPLIERS_FAILED, nil)
		return
	}

	count, err := supplierService.CountSearch()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_SUPPLIERS_COUNT_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["suppliers"] = suppliers
	data["count"] = count

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetSupplierDetail(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("sid")).MustInt()

	supplierService := supplier_service.Supplier{
		SupplierID:   id,
		DrugPageSize: setting.AppSetting.PageSize,
		DrugPageNum:  util.GetPage(c),
	}

	flag, err := supplierService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_SUPPLIER_FAILED, nil)
		return
	}

	if !flag {
		appG.Response(http.StatusBadRequest, e.ERROR_SUPPLIER_NOT_EXIST, nil)
		return
	}

	supplier, err := supplierService.GetSupplierDetail()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_SUPPLIER_DETAIL_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["supplier"] = supplier

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetSearchSupplierDetail(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("sid")).MustInt()
	searchContent := c.Query("search_content")

	supplierService := supplier_service.Supplier{
		SupplierID:    id,
		DrugPageSize:  setting.AppSetting.PageSize,
		DrugPageNum:   util.GetPage(c),
		SearchContent: searchContent,
	}

	flag, err := supplierService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_SUPPLIER_FAILED, nil)
		return
	}

	if !flag {
		appG.Response(http.StatusBadRequest, e.ERROR_SUPPLIER_NOT_EXIST, nil)
		return
	}

	supplier, err := supplierService.GetSearchSupplierDetail()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_SUPPLIER_DETAIL_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["supplier"] = supplier

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
