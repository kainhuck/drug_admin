package inventory_drug_service

import "drug_api/models"

type InventoryDrug struct {
	InventoryDrugID int
	DrugID          int
	PurchasePrice   int
	SalePrice       int
	SupplierID      int
	InventoryNum    int
	Drug            *models.Drug
	PageNum         int
	PageSize        int
	NewSalePrice    int
}

func (i *InventoryDrug) GetAllInvDrugs() ([]*models.InventoryDrug, error) {
	return models.GetInvDrugs(i.PageNum, i.PageSize, map[string]interface{}{})
}

func (i *InventoryDrug) GetAllInvDrugsCustomer() ([]*models.InventoryDrug, error) {
	return models.GetInvDrugsCustomer(i.PageNum, i.PageSize, map[string]interface{}{})
}

func (i *InventoryDrug) Count() (int, error) {
	return models.GetInvDrugTotal(map[string]interface{}{})
}

func (i *InventoryDrug)EditInvDrugSalePrice()error{
	return models.EditInvDrugSalePrice(i.InventoryDrugID, i.NewSalePrice)
}