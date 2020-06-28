package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DrugBuyOrder struct {
	DrugBuyOrderID int       `gorm:"primary_key" json:"drug_buy_order_id"`
	ManagerID      int       `json:"manager_id"`
	BuyDate        time.Time `gorm:"TYPE:DATETIME" json:"buy_date"`
	SupplierID     int       `json:"supplier_id"`
}

type BuyDrugTotalPrice struct {
	DrugBuyOrderID int
	TotalPrice     int
}

// AddDrugBuyOrder 增加进货订单
func AddDrugBuyOrder(mid int, sid int, buyDate time.Time, drugs []DrugWithNum) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 1. 生成销售订单,返回id
	dbo := DrugBuyOrder{
		ManagerID:  mid,
		SupplierID: sid,
		BuyDate:    buyDate,
	}
	err := tx.Create(&dbo).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 2. 根据订单id,插入药品清单
	for _, drug := range drugs {
		bDrug := BuyDrugList{
			DrugID:         drug["DrugID"],
			Num:            drug["Num"],
			DrugBuyOrderID: dbo.DrugBuyOrderID,
		}

		err := tx.Create(&bDrug).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		// 3. 插入库存药
		//invDrug := InventoryDrug{
		//	DrugID: drug["DrugID"],
		//	PurchasePrice: drug[""],
		//	SalePrice: drug[""],
		//	SupplierID: sid,
		//	InventoryNum: drug["Num"],
		//}
		var invDrug InventoryDrug
		err = tx.Where("drug_id = ? and supplier_id = ?", drug["DrugID"], sid).First(&invDrug).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			tx.Rollback()
			return err
		}

		if invDrug.InventoryDrugID > 0 {
			// 3.1 如果存在就更新
			// 增加数量
			err = tx.Model(&invDrug).Where("inventory_drug_id = ?", invDrug.InventoryDrugID).Update("inventory_num", invDrug.InventoryNum+drug["Num"]).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			// 3.2 不存在 新建
			// 获取厂家卖的价格
			var sd SupplierDrug
			err = tx.Where("drug_id = ? and supplier_id = ?", drug["DrugID"], sid).First(&sd).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			invDrug = InventoryDrug{
				DrugID:        drug["DrugID"],
				PurchasePrice: sd.SalePrice,
				SalePrice:     drug["SalePrice"],
				SupplierID:    sid,
				InventoryNum:  drug["Num"],
			}
			err = tx.Create(&invDrug).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

func GetPeriodBuy(startTime, endTime string, pageNum, pageSize int) ([]DrugBuyOrder, error) {
	var drugS []DrugBuyOrder
	err := db.Where("buy_date BETWEEN ? AND DATE_ADD(?,INTERVAL 1 DAY)", startTime, endTime).Offset(pageNum).Limit(pageSize).Find(&drugS).Error
	if err != nil {
		return nil, err
	}
	return drugS, nil
}

// 获取某时间段的订单总数
func GetPeriodBuyTotal(startTime, endTime string) (int, error) {
	var count int
	err := db.Model(DrugBuyOrder{}).Where("buy_date BETWEEN ? AND DATE_ADD(?,INTERVAL 1 DAY)", startTime, endTime).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetTotalBuy() (int, error) {
	var orderList []BuyDrugTotalPrice
	totalPrice := 0
	err := db.Model(BuyDrugTotalPrice{}).Find(&orderList).Error
	if err != nil {
		return 0, err
	}
	for _, v := range orderList {
		totalPrice += v.TotalPrice
	}

	return totalPrice, nil
}

func GetDrugBuyOrderByID(id int) (*DrugBuyOrder, error) {
	var order DrugBuyOrder
	err := db.Where("drug_buy_order_id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func GetBuyOrderTotalPrice(id int) (int, error) {
	var price struct {
		TotalPrice int `json:"total_price"`
	}
	err := db.Table("buy_drug_total_price").Where("drug_buy_order_id = ?", id).First(&price).Error
	if err != nil {
		return 0, err
	}

	return price.TotalPrice, nil
}
