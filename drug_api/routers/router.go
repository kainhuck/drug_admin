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
	// 顾客注册 +
	r.POST("/auth/customer", v1.AddCustomer)

	apiv1 := r.Group("/api/v1")
	apiv1.GET("/drugs/:id", v1.GetDrugDetail)
	apiv1.Use(jwt.JWTManager())
	{
		// 获取全部供应商
		apiv1.GET("/suppliers", v1.GetAllSuppliers)
		apiv1.GET("/searchSuppliers", v1.GetAllSearchSuppliers)
		// 获取具体供应商
		apiv1.GET("/suppliers/:sid", v1.GetSupplierDetail)
		apiv1.GET("/searchSuppliers/:sid", v1.GetSearchSupplierDetail)
		// 获取全部库存药,管理员版
		apiv1.GET("/invdrugs", v1.GetAllInvDrugs)
		// 查找全部库存药,管理员版
		apiv1.GET("/searchInvDrug", v1.SearchAllInvDrugs)
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
		// 修改员工职称
		apiv1.PUT("/employeePosition/:id", v1.EditEmployeePosition)

		// 经理注册
		apiv1.POST("/manager", v1.AddManager)
		// 员工注册
		apiv1.POST("/employee", v1.AddEmployee)

		// 统计总销售额
		apiv1.GET("/totalSales", v1.GetTotalSales)
		// 统计总进货额
		apiv1.GET("/totalBuy", v1.GetTotalBuy)
		// 统计总进利润
		apiv1.GET("/totalProfit", v1.GetTotalProfit)

		// 获取某一个售出订单的详细信息
		apiv1.GET("/detailSaleOrder/:sid", v1.GetDetailSaleOrder)
		// 获取某一个进货订单的详细信息
		apiv1.GET("/detailBuyOrder/:bid", v1.GetDetailBuyOrder)

		// 获取某一个员工的销售信息信息
		apiv1.GET("/employeeSale/:eid", v1.GetEmployeeSaleInfo)
		// 获取某一个顾客的所有订单
		apiv1.GET("/customerSale/:cid", v1.GetCustomerSaleInfo)

		// 获取所有员工
		apiv1.GET("/employees", v1.GetEmployees)
		// 获取所有顾客
		apiv1.GET("/customers", v1.GetCustomers)

		// 获取经理详细信息
		apiv1.GET("/manager/:mid", v1.GetManager)
	}
	apiv1cus := r.Group("/api/v1/c")
	apiv1cus.Use(jwt.JWTCustomer())
	{
		// 获取某一个顾客的所有订单 顾客版
		apiv1cus.GET("/customerSale/:cid", v1.GetCustomerSaleInfo)
		// 获取所有库存药,顾客版 +
		apiv1cus.GET("/invdrugs", v1.GetAllInvDrugsCustomer)
		apiv1cus.GET("/searchInvDrug", v1.SearchAllInvDrugsCustomer)
		// 新增销售订单 +
		apiv1cus.POST("/saleorder", v1.AddDrugSaleOrder)
		// 修改顾客密码 +
		apiv1cus.PUT("/customer/:id", v1.EditCustomerPassword)
		// 获取某一个售出订单的详细信息 顾客版 +
		apiv1cus.GET("/detailSaleOrder/:sid", v1.GetCustomerDetailSaleOrder)
		apiv1cus.GET("/employees", v1.GetAllEmployees)
	}

	return r
}
