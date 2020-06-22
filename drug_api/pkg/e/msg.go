package e

var MsgFlags = map[int]string{
	SUCCESS:                                   "ok",
	ERROR:                                     "fail",
	INVALID_PARAMS:                            "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:               "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:            "Token已超时",
	ERROR_AUTH_TOKEN:                          "Token生成失败",
	ERROR_AUTH:                                "Token错误",
	ERROR_GET_SUPPLIERS_FAILED:                "获取供应商信息失败",
	ERROR_GET_SUPPLIERS_COUNT_FAILED:          "获取供应商数量失败",
	ERROR_GET_SUPPLIER_DETAIL_FAILED:          "获取供应商详情失败",
	ERROR_CHECK_EXIST_SUPPLIER_FAILED:         "检查供应商是否存在失败",
	ERROR_SUPPLIER_NOT_EXIST:                  "供应商不存在",
	ERROR_GET_INVDRUGS_FAILED:                 "获取库存药品失败",
	ERROR_GET_INVDRUGS_COUNT_FAILED:           "获取库存药品数量失败",
	ERROR_EDIT_INVDRUG_SALE_PRICE_FAILED:      "修改售价失败",
	ERROR_ADD_DRUG_SALE_ORDER_FAILED:          "生成销售订单失败",
	ERROR_GET_PERIOD_SALES_ORDER_FAILED:       "获取这段时间订单数据失败",
	ERROR_GET_PERIOD_SALES_ORDER_COUNT_FAILED: "获取这段时间内的订单总数失败",
	ERROR_GET_TOTAL_SALES_FAILED:              "获取总销售额失败",
	ERROR_GET_DETAIL_SALE_ORDER_FAILED:        "获取售出订单详情失败",
	ERROR_ADD_DRUG_BUY_ORDER_FAILED:           "进货订单生成失败",
	ERROR_GET_PERIOD_BUY_ORDER_FAILED:         "获取这段时间内的进货订单失败",
	ERROR_GET_PERIOD_BUY_ORDER_COUNT_FAILED:   "获取这段时间内的订单总数失败",
	ERROR_GET_TOTAL_BUY_FAILED:                "获取进货总额失败",
	ERROR_GET_DETAIL_BUY_ORDER_FAILED:         "获取进货订单详情失败",
	ERROR_GET_EMPLOYEE_SALES_FAILED:           "获取这段时间内员工销售订单失败",
	ERROR_GET_EMPLOYEE_SALES_COUNT_FAILED:     "获取员工销售订单数量失败",
	ERROR_GET_CUSTOMER_SALES_FAILED:           "获取这段时间内顾客销售订单失败",
	ERROR_GET_CUSTOMER_SALES_COUNT_FAILED:     "获取顾客销售订单数量失败",
	ERROR_DIFF_PASSWORD:                       "两次密码不一样",
	ERROR_EDIT_PASSWORD_FAILED:                "修改密码失败",
	ERROR_NEW_MANAGER_FAILED:                  "新建经理失败",
	ERROR_CHECK_EXIST_FAILED:                  "检查是否存在失败",
	ERROR_EXIST_USERNAME:                      "账户名已存在",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
