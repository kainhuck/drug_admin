package models

import "github.com/jinzhu/gorm"

type Drug struct {
	DrugID           int        `gorm:"primary_key" json:"drug_id"`
	Cname            string     `json:"cname"`
	Ename            string     `json:"ename"`
	Introduction     string     `json:"introduction"`
	Component        string     `json:"component"`
	Property         string     `json:"property"`
	Indication       string     `json:"indication"`
	MedicFormat      string     `json:"medic_format"`
	Taboo            string     `json:"taboo"`
	Ytime            string     `json:"ytime"`
	Mstandard        string     `json:"mstandard"`
	Dosage           string     `json:"dosage"`
	AdverseReactions string     `json:"adverseReactions"`
	Interactions     string     `json:"interactions"`
	Notice           string     `json:"notice"`
	DrugType         string     `json:"drug_type"`
	DrugHealthType   string     `json:"drug_health_type"`
	DrugRecipeType   string     `json:"drug_recipe_type"`
}

type DrugWithNum map[string]int
// ExistDrugByID checks
func ExistDrugByID(id int) (bool, error) {
	var drug Drug
	err := db.Select("durg_id").Where("durg_id = ?", id).First(&drug).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if drug.DrugID > 0 {
		return true, nil
	}

	return false, nil
}

// GetDrugTotal gets
func GetDrugTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Drug{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetDrugs gets
func GetDrugs(pageNum int, pageSize int, maps interface{}) ([]*Drug, error) {
	var drugs []*Drug
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&drugs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return drugs, nil
}

// GetDrug ...
func GetDrugBuyID(id int) (*Drug, error) {
	var drug Drug
	err := db.Where("drug_id = ?", id).First(&drug).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &drug, nil
}
