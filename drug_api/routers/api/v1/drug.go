package v1

import (
	"drug_api/models"
	"drug_api/pkg/app"
	"drug_api/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetDrugDetail (c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	isExist, err := models.ExistDrugByID(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_FAILED ,nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_DRUG, nil)
		return
	}

	drug, err := models.GetDrugBuyID(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_DRUG_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["drug"] = drug

	appG.Response(http.StatusOK, e.SUCCESS, data)
	return
}