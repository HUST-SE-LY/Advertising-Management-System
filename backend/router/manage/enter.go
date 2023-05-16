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
	manageRouterWithoutJWT := router.Group("manage")
	var manageAdminApi = api.ApiGroupApp.ManageApiGroup.ManageAdminApi
	var manageCompanyApi = api.ApiGroupApp.ManageApiGroup.ManageCompanyApi
	var manageAdvertisement = api.ApiGroupApp.ManageApiGroup.ManageAdvertisementApi
	{
		manageRouterWithoutJWT.POST("admin/create", manageAdminApi.CreateAdmin)
		manageRouter.GET("company/list", manageCompanyApi.GetAllCompanies)
		manageRouter.GET("company/search", manageCompanyApi.GetCompaniesByTerm)
		// TODO(Determine the name of the route)
		manageRouter.GET("company/review", manageCompanyApi.GetAllCompaniesToBeReviewed)
		manageRouter.POST("company/register", manageCompanyApi.AllowRegistrationForCompanies)
		manageRouter.POST("company/info", manageCompanyApi.AllowUpdateForCompanies)

		manageRouter.GET("advertisement/list", manageAdvertisement.GetAllAdvertisements)
		manageRouter.GET("advertisement/review", manageAdvertisement.GetAllAdvertisementsToBeReviewed)
		manageRouter.GET("advertisement/allow", manageAdvertisement.AllowAdvertisement)
	}
}
