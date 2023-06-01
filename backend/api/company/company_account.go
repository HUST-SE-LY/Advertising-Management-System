package company

import (
	"backend/models/company_model/request"
	"backend/models/company_model/response"
	"backend/utils/gin_ext"
	"backend/utils/jwt"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type CompanyAccountApi struct {
}

// CompanyRegister godoc
//
//	@Summary		Company register
//	@Description	Company register
//
//	@Tags			Company
//	@Accept			json
//	@Produce		json
//	@Param			request_body	body		request.CompanyRegisterReq	true	"company register request"
//	@Success		200				{object}	nil							"response"
//	@Router			/company/register [post]
func (com *CompanyAccountApi) CompanyRegister(c *gin.Context) {
	var companyRegisterReq request.CompanyRegisterReq
	if err := gin_ext.BindJSON(c, &companyRegisterReq); err != nil {
		return
	}
	if err := companyService.CompanyRegister(&companyRegisterReq); err != nil {
		// Same account
		if st, ok := err.(status.Status); ok {
			c.JSON(http.StatusBadRequest, gin_ext.Response(st, nil))
			return
		}
	}

	c.JSON(http.StatusOK, gin_ext.Response(nil, nil))
}

// CompanyLogin godoc
//
//	@Summary		Company login
//	@Description	Company Login by account and password
//
//	@Tags			Company
//	@Accept			json
//	@Produce		json
//	@Param			request_body	body		request.CompanyLoginReq		true	"company login request"
//	@Success		200				{object}	response.CompanyLoginResp	"Company login response"
//	@Router			/company/login [post]
func (com *CompanyAccountApi) CompanyLogin(c *gin.Context) {
	var companyLoginReq request.CompanyLoginReq
	if err := gin_ext.BindJSON(c, &companyLoginReq); err != nil {
		return
	}
	if company, companyToken, err := companyService.CompanyLogin(companyLoginReq); err != nil {
		c.JSON(http.StatusUnauthorized, gin_ext.Response(err, nil))
	} else {
		jsonResp, _ := jsoniter.Marshal(response.CompanyLoginResp{company.Account, companyToken.Token})
		c.Header("Authorization", companyToken.Token)
		c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
	}
}

// CompanyLogout godoc
//
//	@Summary		Company logout
//	@Description	Company logout
//
//	@Tags			Company
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	nil	"response"
//	@Router			/company/logout [post]
func (com *CompanyAccountApi) CompanyLogout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	err := companyService.CompanyLogout(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.LogoutError, nil))
		return
	}
	c.JSON(http.StatusOK, gin_ext.Response(nil, nil))
}

// CompanyUpdateInfo godoc
//
//	@Summary		Company update info
//	@Description	Company update info
//
//	@Tags			Company
//	@Accept			json
//	@Produce		json
//	@Param			request_body	body		request.CompanyUpdateInfoReq	true	"company update info  request"
//	@Success		200				{object}	nil								"Company UpdateInfoReq response"
//	@Router			/company/update-info [post]
func (com *CompanyAccountApi) CompanyUpdateInfo(c *gin.Context) {
	var companyUpdateInfoReq request.CompanyUpdateInfoReq
	token := c.Request.Header.Get("Authorization")
	if err := gin_ext.ShouldBindJSON(c, &companyUpdateInfoReq); err != nil {
		return
	}
	if err := companyService.CompanyUpdateInfo(companyUpdateInfoReq, token); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
	} else {
		c.JSON(http.StatusOK, gin_ext.Response(nil, nil))
	}
}

// CompanyUpdatePwd godoc
//
//	@Summary		Company update password
//	@Description	Company update password
//
//	@Tags			Company
//	@Accept			json
//	@Produce		json
//	@Param			request_body	body		request.CompanyUpdatePwdReq	false	"company update password request"
//	@Success		200				{object}	nil							"Company update password response"
//	@Router			/company/update-pwd [post]
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

// CompanyCancel godoc
//
//	@Summary		Company Cancel
//	@Description	Company Cancel
//
//	@Tags			Company
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	nil	"nil"
//	@Router			/company/cancel [post]
func (com *CompanyAccountApi) CompanyCancel(c *gin.Context) {
	// TODO
	token := c.Request.Header.Get("Authorization")

	if err := companyService.CompanyCancel(token); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
	} else {
		c.JSON(http.StatusOK, gin_ext.Response(nil, nil))
	}
}

// CompanyGetInfo godoc
//
//	@Summary		Company get information
//	@Description	Company get information
//
//	@Tags			Company
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.CompanyGetInfoResp	"company info"
//	@Router			/company/get-info [get]
func (com *CompanyAccountApi) CompanyGetInfo(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	company, err := companyService.CompanyGetInfo(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
		return
	}
	companyInfo := company.GetInfo()
	resp := response.CompanyGetInfoResp{CompanyInfo: companyInfo}
	jsonResp, _ := jsoniter.Marshal(resp)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}
func (com *CompanyAccountApi) CompanyRecharge(c *gin.Context) {
	var req request.CompanyRechargeReq
	if err := gin_ext.BindJSON(c, &req); err != nil {
		return
	}
	token := c.Request.Header.Get("Authorization")
	claims, _ := jwt.ParseToken(token)
	account := claims.Username
	if err := companyService.CompanyRecharge(req, account); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
	} else {
		c.JSON(http.StatusOK, gin_ext.Response(nil, nil))
	}
}
