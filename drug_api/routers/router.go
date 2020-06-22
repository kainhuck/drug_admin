package routers

import (
	"drug_api/middleware/jwt"
	v1 "drug_api/routers/api/v1"

	"drug_api/routers/api"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/auth/m", api.GetMAuth)
	r.POST("/auth/c", api.GetCAuth)
	r.POST("/auth/e", api.GetEAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/suppliers", v1.GetAllSuppliers)
		apiv1.GET("/suppliers/:sid", v1.GetSupplierDetail)
		apiv1.GET("/invdrugs", v1.GetAllInvDrugs)
		apiv1.PUT("/invdrugs/:id", v1.EditInvDrugSalePrice)
		apiv1.GET("/cinvdrugs", v1.GetAllInvDrugsCustomer)

		apiv1.POST("/saleorder", v1.AddDrugSaleOrder)
		apiv1.POST("/buyorder", v1.AddDrugBuyOrder)

		apiv1.GET("/periodSales", v1.GetPeriodSales)
		apiv1.GET("/periodBuy", v1.GetPeriodBuy)

		apiv1.PUT("/manager/:id", v1.EditManagerPassword)
		apiv1.PUT("/employee/:id", v1.EditEmployeePassword)
		apiv1.PUT("/customer/:id", v1.EditCustomerPassword)

		apiv1.POST("/manager", v1.AddManager)
		apiv1.POST("/employee", v1.AddEmployee)
		apiv1.POST("/customer", v1.AddCustomer)

		apiv1.GET("/totalSales", v1.GetTotalSales)
		apiv1.GET("/totalBuy", v1.GetTotalBuy)

		apiv1.GET("/detailSaleOrder/:sid", v1.GetDetailSaleOrder)
		apiv1.GET("/cdetailSaleOrder/:sid", v1.GetCustomerDetailSaleOrder)
		apiv1.GET("/detailBuyOrder/:bid", v1.GetDetailBuyOrder)

		apiv1.GET("/employeeSale/:eid", v1.GetEmployeeSaleInfo)
		apiv1.GET("/customerSale/:cid", v1.GetCustomerSaleInfo)
	}

	return r
}
