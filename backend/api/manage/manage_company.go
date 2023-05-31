package manage

import (
	"backend/global"
	"backend/models/company_model/entity"
	"backend/models/manage_model/enum"
	"backend/models/manage_model/request"
	"backend/models/manage_model/response"
	"backend/utils/functional"
	"backend/utils/gin_ext"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"strconv"
)

type ManageCompanyApi struct {
}

// GetPendingReviewCompaniesCount godoc
//
//	@Summary	Get info pending review companies count
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.GetCompaniesCountResp
//	@Router		/manage/company/review_count [get]
func (m *ManageCompanyApi) GetPendingReviewCompaniesCount(c *gin.Context) {
	count, err := adminService.ManageCompanyService.GetPendingReviewCompaniesCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}
	resp := response.GetCompaniesCountResp{Count: count}
	jsonResp, _ := jsoniter.Marshal(resp)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

// GetInfoPendingReviewCompaniesCount godoc
//
//	@Summary	Get info pending review companies count
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.GetCompaniesCountResp
//	@Router		/manage/company/info_review_count [get]
func (m *ManageCompanyApi) GetInfoPendingReviewCompaniesCount(c *gin.Context) {
	count, err := adminService.ManageCompanyService.GetInfoPendingReviewCompaniesCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}
	resp := response.GetCompaniesCountResp{Count: count}
	jsonResp, _ := jsoniter.Marshal(resp)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

// GetAllCompanies godoc
//
//	@Summary	Get all companies
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.GetCompaniesResp	"Get companies response"
//	@Router		/manage/company/list [get]
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

// GetAllCompaniesToBeReviewed godoc
//
//	@Summary	Get all companies to be reviewed
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.GetCompaniesToBeReviewedResp	"Get companies to be reviewed response"
//	@Router		/manage/company/review [get]
func (m *ManageCompanyApi) GetAllCompaniesToBeReviewed(c *gin.Context) {
	companies, err := adminService.ManageCompanyService.GetAllCompaniesToBeReviewed()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}
	companiesInfo := functional.Map(companies, entity.CompanyPendingReview.GetInfo)
	resp := response.GetCompaniesToBeReviewedResp{CompanyInfos: companiesInfo}
	jsonResp, _ := jsoniter.Marshal(resp)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

// GetCompaniesByTerm godoc
//
//	@Summary		Get Companies by term
//	@Description	Get Companies by term
//
//	@Tags			Manage
//	@Accept			json
//	@Produce		json
//	@Param			term	query		string						true	"term"
//	@Param			type	query		int							true	"Enum: 0 -> Account, 1 -> Name, 2 -> Address, 3 -> ManagerName, 4 -> MangerTel, 5 -> BusinessLicenseNumber"
//	@Success		200		{object}	response.GetCompaniesResp	"Company Info"
//	@Router			/company/search [get]
func (m *ManageCompanyApi) GetCompaniesByTerm(c *gin.Context) {
	term := c.Query("term")
	_type, err := strconv.Atoi(c.Query("type"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.InvalidParams, nil))
		return
	}
	termType := enum.CompanySearchType(_type)
	if companies, err := adminService.GetCompaniesByTerm(term, termType); err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	} else {
		companiesInfo := functional.Map(companies, entity.Company.GetInfo)
		resp := response.GetCompaniesResp{CompanyInfos: companiesInfo}
		jsonResp, _ := jsoniter.Marshal(resp)
		c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
	}
}

// AllowRegistrationForCompanies godoc
//
//	@Summary	Allow Registration For Companies
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Param		request_body	body		request.AllowCompaniesRegisterReq	true	"allow companies register request"
//	@Success	200				{object}	[]string							"allow companies response"
//	@Router		/manage/company/register [post]
func (m *ManageCompanyApi) AllowRegistrationForCompanies(c *gin.Context) {
	var allowCompaniesRegisterReq request.AllowCompaniesRegisterReq
	if err := gin_ext.BindJSON(c, &allowCompaniesRegisterReq); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	companyAccount := allowCompaniesRegisterReq.CompanyAccounts
	allowAccounts, err := adminService.ManageCompanyService.AllowRegistrationForCompanies(companyAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}
	// Delete company in pending reviews
	err = global.GVA_DB.Where("account LIKE ?", companyAccount).Delete(&entity.CompanyPendingReview{}).Error
	jsonResp, _ := jsoniter.Marshal(allowAccounts)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

// AllowUpdateForCompanies godoc
//
//	@Summary	Allow Update For Companies
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Param		request_body	body		request.AllowCompaniesUpdateReq	true	"allow companies update request"
//	@Success	200				{object}	[]string						"allow update for companies response"
//	@Router		/manage/company/info [post]
func (m *ManageCompanyApi) AllowUpdateForCompanies(c *gin.Context) {
	var allowCompaniesUpdateReq request.AllowCompaniesUpdateReq
	if err := gin_ext.BindJSON(c, &allowCompaniesUpdateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	allowAccounts, err := adminService.ManageCompanyService.AllowCompaniesInfoUpdate(allowCompaniesUpdateReq.CompanyAccounts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}
	jsonResp, _ := jsoniter.Marshal(allowAccounts)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}
