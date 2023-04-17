package company

import (
	"backend/models/company_model/request"
	"backend/utils/gin_ext"
	"backend/utils/jwt"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type CompanyAccountApi struct {
}

func (com *CompanyAccountApi) CompanyRegister(c *gin.Context) {
	var companyRegisterReq request.CompanyRegisterReq
	if err := gin_ext.BindJSON(c, &companyRegisterReq); err != nil {
		return
	}
	if err := companyService.RegisterCompany(&companyRegisterReq); err != nil {
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
	if err := gin_ext.BindJSON(c, &companyLoginReq); err != nil {
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
	if err := gin_ext.ShouldBindJSON(c, &companyUpdateInfoReq); err != nil {
		return
	}
	if err := companyService.CompanyUpdateInfo(companyUpdateInfoReq); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
	} else {
		c.JSON(http.StatusOK, gin_ext.Response(nil, nil))
	}
}

func (com *CompanyAccountApi) CompanyUpdatePwd(c *gin.Context) {
	var companyUpdatePwdReq request.CompanyUpdatePwdReq
	if err := gin_ext.ShouldBindJSON(c, &companyUpdatePwdReq); err != nil {
		return
	}
	token := c.Request.Header.Get("Authorization")
	claims, _ := jwt.ParseToken(token)
	account := claims.Username
	if err := companyService.CompanyUpdatePwd(companyUpdatePwdReq, account); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
	} else {
		c.JSON(http.StatusOK, gin_ext.Response(nil, nil))
	}
}

func (com *CompanyAccountApi) CompanyCancel(c *gin.Context) {
	// TODO
	token := c.Request.Header.Get("Authorization")

	if err := companyService.CompanyCancel(token); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
	} else {
		c.JSON(http.StatusOK, gin_ext.Response(nil, nil))
	}

}
