package v1

import (
	"drug_api/pkg/app"
	"drug_api/pkg/e"
	"drug_api/service/manager_service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func EditManagerPassword(c *gin.Context){
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	newPassword := c.PostForm("new_password")
	confirmPassword := c.PostForm("confirm_password")

	if newPassword != confirmPassword{
		appG.Response(http.StatusInternalServerError, e.ERROR_DIFF_PASSWORD ,nil)
		return
	}

	managerService := manager_service.Manager{
		ManagerID: id,
		NewPassword: newPassword,
	}

	err := managerService.EditManagerPassword()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_PASSWORD_FAILED ,nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func AddManager(c *gin.Context) {
	appG := app.Gin{C: c}
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")

	if password != confirmPassword{
		appG.Response(http.StatusInternalServerError, e.ERROR_DIFF_PASSWORD ,nil)
		return
	}

	managerService := manager_service.Manager{
		Username: username,
		Password: password,
	}

	flag, err := managerService.ExistByUsername()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_FAILED ,nil)
		return
	}

	if flag{
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_USERNAME ,nil)
		return
	}

	id ,err := managerService.AddManager()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NEW_MANAGER_FAILED,nil)
		return
	}

	data := make(map[string]interface{})
	data["id"] = id

	appG.Response(http.StatusOK, e.SUCCESS, data)
}