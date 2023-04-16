package manage

import (
	"backend/models/company_model/entity"
	"backend/models/manage_model/response"
	"backend/utils/functional"
	"backend/utils/gin_ext"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type ManageCompanyApi struct {
}

func (m *ManageCompanyApi) GetAllCompanies(c *gin.Context) {
	companies, err := adminService.ManageCompanyService.GetAllCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}
	companiesInfo := functional.Map(companies, entity.Company.GetInfo)
	resp := response.GetCompaniesResp{CompanyInfos: companiesInfo}
	jsonResp, _ := jsoniter.Marshal(resp)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

func (m *ManageCompanyApi) GetAllCompaniesToBeReviewed(c *gin.Context) {
	companies, err := adminService.ManageCompanyService.GetAllCompaniesToBeReviewed()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}
	companiesInfo := functional.Map(companies, entity.CompanyToBeReviewed.GetInfo)
	resp := response.GetCompaniesToBeReviewedResp{CompanyInfos: companiesInfo}
	jsonResp, _ := jsoniter.Marshal(resp)
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
	allowAccounts := functional.Map(*allowCompanies, func(com entity.Company) string {
		return com.Account
	})
	jsonResp, _ := jsoniter.Marshal(allowAccounts)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}
