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
	SearchContent   string
}

func (i *InventoryDrug) GetAllInvDrugs() ([]*models.InventoryDrug, error) {
	return models.GetInvDrugs(i.PageNum, i.PageSize, map[string]interface{}{})
}

func (i *InventoryDrug) SearchAllInvDrugs () ([]*models.InventoryDrugWithName, error) {
	return models.SearchAllInvDrugs(i.PageNum, i.PageSize, i.SearchContent)
}

func (i *InventoryDrug) SearchAllInvDrugsCustomer () ([]*models.InventoryDrugWithName, error) {
	drugs ,err := models.SearchAllInvDrugs(i.PageNum, i.PageSize, i.SearchContent)
	if err != nil {
		return nil, err
	}

	for _, v:= range drugs{
		v.PurchasePrice = -1
	}

	return drugs, nil
}

func (i *InventoryDrug) GetAllInvDrugsCustomer() ([]*models.InventoryDrug, error) {
	return models.GetInvDrugsCustomer(i.PageNum, i.PageSize, map[string]interface{}{})
}

func (i *InventoryDrug) Count() (int, error) {
	return models.GetInvDrugTotal(map[string]interface{}{})
}

func (i *InventoryDrug) CountSearch() (int, error) {
	return models.GetSearchInvDrugTotal(i.SearchContent)
}

func (i *InventoryDrug) EditInvDrugSalePrice() error {
	return models.EditInvDrugSalePrice(i.InventoryDrugID, i.NewSalePrice)
}
