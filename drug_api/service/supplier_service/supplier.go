package supplier_service

import "drug_api/models"

type Supplier struct {
	SupplierID    int    `json:"supplier_id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Drugs         []*models.DrugWithSalePrice
	PageNum       int `json:"page_num"`
	PageSize      int `json:"page_size"`
	DrugPageNum   int
	DrugPageSize  int
	SearchContent string
}

func (s *Supplier) GetAllSupplier() ([]*models.Supplier, error) {
	suList, err := models.GetSuppliersInfo(s.PageNum, s.PageSize, map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	return suList, nil
}

func (s *Supplier)GetAllSearchSuppliers() ([]*models.Supplier, error){
	return models.GetAllSearchSuppliers(s.PageNum, s.PageSize, s.SearchContent)
}

func (s *Supplier) GetSupplierDetail() (*models.Supplier, error) {
	return models.GetSupplierInfo(s.SupplierID, s.DrugPageNum, s.DrugPageSize)
}

func (s *Supplier) GetSearchSupplierDetail() (*models.Supplier, error) {
	return models.GetSearchSupplierInfo(s.SupplierID, s.DrugPageNum, s.DrugPageSize, s.SearchContent)
}

func (s *Supplier) Count() (int, error) {
	return models.GetSupplierTotal(map[string]interface{}{})
}

func (s *Supplier) CountSearch() (int, error) {
	return models.GetSearchSupplierTotal(s.SearchContent)
}

func (s *Supplier) ExistByID() (bool, error) {
	return models.ExistSupplierByID(s.SupplierID)
}
