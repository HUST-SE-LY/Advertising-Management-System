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
	"sort"
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
//	@Router			/manage/company/search [get]
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
//	@Param		request_body	body		request.CompaniesRegisterReq	true	"array of companies' accounts"
//	@Success	200				{object}	[]string						"allow companies response"
//	@Router		/manage/company/register [post]
func (m *ManageCompanyApi) AllowRegistrationForCompanies(c *gin.Context) {
	var allowCompaniesRegisterReq request.CompaniesRegisterReq
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

	jsonResp, _ := jsoniter.Marshal(allowAccounts)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

// RejectRegistrationForCompanies godoc
//
//	@Summary	Reject Registration For Companies
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Param		request_body	body		request.CompaniesRegisterReq	true	"reject companies register request"
//	@Success	200				{object}	nil								"reject companies response"
//	@Router		/manage/company/reject-registration [post]
func (m *ManageCompanyApi) RejectRegistrationForCompanies(c *gin.Context) {
	var RejectCompaniesRegisterReq request.CompaniesRegisterReq
	if err := gin_ext.BindJSON(c, &RejectCompaniesRegisterReq); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	companyAccount := RejectCompaniesRegisterReq.CompanyAccounts
	err := adminService.ManageCompanyService.RejectRegistrationForCompanies(companyAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}
	c.JSON(http.StatusOK, gin_ext.Response(nil, nil))
}

// AllowUpdateForCompanies godoc
//
//	@Summary	Allow Update For Companies
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Param		request_body	body		request.CompaniesUpdateReq	true	"account不能修改，Body不能出现account，根据token来判断account，其余每一项都是可选的，但是必须要有一项"
//	@Success	200				{object}	[]string					"allow update for companies response"
//	@Router		/manage/company/info [post]
func (m *ManageCompanyApi) AllowUpdateForCompanies(c *gin.Context) {
	var allowCompaniesUpdateReq request.CompaniesUpdateReq
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

// RejectUpdateForCompanies godoc
//
//	@Summary	Reject Update For Companies
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Param		request_body	body		request.CompaniesUpdateReq	true	"account不能修改，Body不能出现account，根据token来判断account，其余每一项都是可选的，但是必须要有一项"
//	@Success	200				{object}	[]string					"reject update for companies response"
//	@Router		/manage/company/reject-info [post]
func (m *ManageCompanyApi) RejectUpdateForCompanies(c *gin.Context) {
	var rejectCompaniesUpdateReq request.CompaniesUpdateReq
	if err := gin_ext.BindJSON(c, &rejectCompaniesUpdateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	rejectAccounts, err := adminService.ManageCompanyService.RejectCompaniesInfoUpdate(rejectCompaniesUpdateReq.CompanyAccounts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}
	jsonResp, _ := jsoniter.Marshal(rejectAccounts)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

// GetCompaniesInfoUpdateList godoc
//
//	@Summary	Get companies info update list
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.GetCompaniesUpdateInfoResp	""
//	@Router		/manage/company/update-info/list [get]
func (m *ManageCompanyApi) GetCompaniesInfoUpdateList(c *gin.Context) {
	var GetCompaniesInfoUpdateListResp response.GetCompaniesUpdateInfoResp
	var companies []entity.Company
	companiesInfoToBeReviewed, err := adminService.ManageCompanyService.GetAllCompaniesInfoToBeReviewed()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
		return
	}

	companiesNumber := len(companiesInfoToBeReviewed)
	companiesAccounts := functional.Map(companiesInfoToBeReviewed, func(com entity.CompanyInfoPendingReview) string {
		return com.Account
	})

	if err := global.GVA_DB.Where("account IN ?", companiesAccounts).Find(&companies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
	}

	sort.Slice(companies, func(i, j int) bool {
		return companies[i].Id < companies[j].Id
	})
	sort.Slice(companiesInfoToBeReviewed, func(i, j int) bool {
		return companiesInfoToBeReviewed[i].Id < companiesInfoToBeReviewed[j].Id
	})
	GetCompaniesInfoUpdateListResp.CompaniesUpdateInfoList = make([]response.CompaniesUpdateInfo, companiesNumber)
	for i, _ := range companiesInfoToBeReviewed {
		GetCompaniesInfoUpdateListResp.CompaniesUpdateInfoList[i] = response.CompaniesUpdateInfo{
			Account:                       companies[i].Account,
			PreviousName:                  companies[i].Name,
			PreviousAddress:               companies[i].Address,
			PreviousManagerName:           companies[i].ManagerName,
			PreviousManagerTel:            companies[i].ManagerTel,
			PreviousBusinessLicenseNumber: companies[i].BusinessLicenseNumber,
			NewName:                       companiesInfoToBeReviewed[i].Name,
			NewAddress:                    companiesInfoToBeReviewed[i].Address,
			NewManagerName:                companiesInfoToBeReviewed[i].ManagerName,
			NewManagerTel:                 companiesInfoToBeReviewed[i].ManagerTel,
			NewBusinessLicenseNumber:      companiesInfoToBeReviewed[i].BusinessLicenseNumber,
		}
	}
	jsonResp, err := jsoniter.Marshal(GetCompaniesInfoUpdateListResp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
		return
	}
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}
