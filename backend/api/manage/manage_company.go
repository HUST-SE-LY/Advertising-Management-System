package manage

import (
	"backend/utils/gin_ext"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type ManageCompanyApi struct {
}

func (m *ManageCompanyApi) ReadAllCompaniesToBeReviewed(c *gin.Context) {
	companies, err := adminService.ManageCompanyService.ReadAllCompaniesToBeReviewed()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}
	jsonResp, _ := jsoniter.Marshal(companies)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

func (m *ManageCompanyApi) AllowRegistrationForCompanies(c *gin.Context) {
	var companyAccounts []string
	err := c.ShouldBindJSON(&companyAccounts)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	allowCompanies, err := adminService.ManageCompanyService.AllowRegistrationForCompanies(&companyAccounts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}
	jsonResp, _ := jsoniter.Marshal(allowCompanies)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}
