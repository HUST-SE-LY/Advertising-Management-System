package manage

import (
	"backend/api"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

type ManageRouterGroup struct {
}

func (a *ManageRouterGroup) Init(router *gin.RouterGroup) {
	manageRouter := router.Group("manage").Use(middleware.AdminJWTAuth())
	var manageAdminApi = api.ApiGroupApp.ManageApiGroup.ManageAdminApi
	var manageCompanyApi = api.ApiGroupApp.ManageApiGroup.ManageCompanyApi
	var manageAdvertisement = api.ApiGroupApp.ManageApiGroup.ManageAdvertisementApi
	var manageSpace = api.ApiGroupApp.ManageApiGroup.ManageSpaceApi
	{
		manageRouter.POST("admin/create", manageAdminApi.CreateAdmin)
		manageRouter.GET("company/review_count", manageCompanyApi.GetPendingReviewCompaniesCount)
		manageRouter.GET("company/info_review_count", manageCompanyApi.GetInfoPendingReviewCompaniesCount)
		manageRouter.GET("company/list", manageCompanyApi.GetAllCompanies)
		manageRouter.GET("company/search", manageCompanyApi.GetCompaniesByTerm)
		// TODO(Determine the name of the route)
		manageRouter.GET("company/review", manageCompanyApi.GetAllCompaniesToBeReviewed)
		manageRouter.POST("company/register", manageCompanyApi.AllowRegistrationForCompanies)
		manageRouter.POST("company/reject-registration", manageCompanyApi.RejectRegistrationForCompanies)
		manageRouter.POST("company/info", manageCompanyApi.AllowUpdateForCompanies)
		manageRouter.POST("company/reject-info", manageCompanyApi.RejectUpdateForCompanies)
		manageRouter.GET("company/update-info/list", manageCompanyApi.GetCompaniesInfoUpdateList)
		manageRouter.GET("advertisement/count", manageAdvertisement.GetAdvertisementsPendingReviewCount)
		manageRouter.GET("advertisement/list", manageAdvertisement.GetAllAdvertisements)
		manageRouter.GET("advertisement/review", manageAdvertisement.GetAllAdvertisementsToBeReviewed)
		manageRouter.POST("advertisement/allow", manageAdvertisement.AllowAdvertisement)
		manageRouter.POST("advertisement/reject", manageAdvertisement.RejectAdvertisement)
		manageRouter.POST("advertisement/delete", manageAdvertisement.DeleteAdvertisement)
		manageRouter.POST("advertisement/stop", manageAdvertisement.StopAdvertisement)
		manageRouter.GET("advertisement/search", manageAdvertisement.GetAdvertisementsByTerm)
		manageRouter.POST("space/set_price", manageSpace.SetSpacePrice)
		manageRouter.GET("space/get", manageSpace.GetSpace)
		manageRouter.POST("space/add", manageSpace.AddSpace)
	}

}
