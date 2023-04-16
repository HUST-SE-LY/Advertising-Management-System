package company

import (
	"backend/models/company_model/request"
	"backend/utils/gin_ext"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type CompanyAccountApi struct {
}

func (com *CompanyAccountApi) RegisterCompany(c *gin.Context) {
	var companyRegisterReq request.CompanyRegisterReq
	err := c.BindJSON(&companyRegisterReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	if err = companyService.RegisterCompany(&companyRegisterReq); err != nil {
		// Same account
		if st, ok := err.(status.Status); ok {
			c.JSON(http.StatusBadRequest, gin_ext.Response(st, nil))
			return
		}
	}

	c.JSON(http.StatusOK, gin_ext.Response(nil, nil))
}

func (com *CompanyAccountApi) CompanyLogin(c *gin.Context) {
	var companyLoginReq request.CompanyLoginReq
	err := c.BindJSON(&companyLoginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	if company, companyToken, err := companyService.CompanyLogin(companyLoginReq); err != nil {
		c.JSON(http.StatusUnauthorized, gin_ext.Response(err, nil))
	} else {
		jsonResp, _ := jsoniter.Marshal(company)
		c.Header("Authorization", companyToken.Token)
		c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
	}
}

func (com *CompanyAccountApi) CompanyLogout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	err := companyService.CompanyLogout(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.LogoutError, nil))
		return
	}
	c.JSON(http.StatusOK, gin_ext.Response(nil, nil))
}

func (com *CompanyAccountApi) CompanyUpdateInfo(c *gin.Context) {
	var companyUpdateInfoReq request.CompanyUpdateInfoReq
	err := c.BindJSON(&companyUpdateInfoReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.LogoutError, nil))
		return
	}
	if err = companyService.CompanyUpdateInfo(companyUpdateInfoReq); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
	} else {
		c.JSON(http.StatusOK, gin_ext.Response(nil, nil))
	}
}
