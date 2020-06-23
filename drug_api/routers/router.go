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
	apiv1.Use(jwt.JWTManager())
	{
		// 获取全部供应商
		apiv1.GET("/suppliers", v1.GetAllSuppliers)
		// 获取具体供应商
		apiv1.GET("/suppliers/:sid", v1.GetSupplierDetail)
		// 获取全部库存药,管理员版
		apiv1.GET("/invdrugs", v1.GetAllInvDrugs)
		// 修改库存要售价
		apiv1.PUT("/invdrugs/:id", v1.EditInvDrugSalePrice)

		// 新增进货订单
		apiv1.POST("/buyorder", v1.AddDrugBuyOrder)

		// 获取某段时间内的销售订单
		apiv1.GET("/periodSales", v1.GetPeriodSales)
		// 获取某段时间内的进货订单
		apiv1.GET("/periodBuy", v1.GetPeriodBuy)

		// 修改经理密码
		apiv1.PUT("/manager/:id", v1.EditManagerPassword)
		// 修改员工密码
		apiv1.PUT("/employee/:id", v1.EditEmployeePassword)

		// 经理注册
		apiv1.POST("/manager", v1.AddManager)
		// 员工注册
		apiv1.POST("/employee", v1.AddEmployee)

		// 统计总销售额
		apiv1.GET("/totalSales", v1.GetTotalSales)
		// 统计总进货额
		apiv1.GET("/totalBuy", v1.GetTotalBuy)

		// 获取某一个售出订单的详细信息
		apiv1.GET("/detailSaleOrder/:sid", v1.GetDetailSaleOrder)
		// 获取某一个售出订单的详细信息 顾客版 +
		apiv1.GET("/cdetailSaleOrder/:sid", v1.GetCustomerDetailSaleOrder)
		// 获取某一个进货订单的详细信息
		apiv1.GET("/detailBuyOrder/:bid", v1.GetDetailBuyOrder)

		// 获取某一个员工的信息信息
		apiv1.GET("/employeeSale/:eid", v1.GetEmployeeSaleInfo)
	}
	apiv1cus := r.Group("/api/v1/c")
	apiv1cus.Use(jwt.JWTCustomer())
	{
		// 获取所有库存药,顾客版 +
		apiv1cus.GET("/invdrugs", v1.GetAllInvDrugsCustomer)
		// 新增销售订单 +
		apiv1cus.POST("/saleorder", v1.AddDrugSaleOrder)
		// 修改顾客密码 +
		apiv1cus.PUT("/customer/:id", v1.EditCustomerPassword)
		// 顾客注册 +
		apiv1cus.POST("/customer", v1.AddCustomer)
		// 获取某一个顾客的详细信息 +
		apiv1cus.GET("/customerSale/:cid", v1.GetCustomerSaleInfo)
	}

	return r
}
