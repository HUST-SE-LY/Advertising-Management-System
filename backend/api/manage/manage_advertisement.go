package manage

import (
	"backend/global"
	"backend/models/advertisement_model/entity"
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

type ManageAdvertisementApi struct {
}

// GetAllAdvertisements godoc
//
//	@Summary	Get all advertisements
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.GetAdvertisementsResp	"All advertisements"
//	@Router		/manage/advertisement/list [post]
func (m *ManageAdvertisementApi) GetAllAdvertisements(c *gin.Context) {
	advertisements, err := adminService.ManageAdvertisementService.GetAllAdvertisements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}
	advertisementsInfo := functional.Map(advertisements, entity.Advertisement.GetInfo)
	resp := response.GetAdvertisementsResp{AdvertisementInfos: advertisementsInfo}
	jsonResp, _ := jsoniter.Marshal(resp)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))

}

// GetAllAdvertisementsToBeReviewed godoc
//
//	@Summary	Get all advertisements to be reviewed
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.GetAdvertisementsResp	"All advertisements"
//	@Router		/manage/advertisement/list [post]
func (m *ManageAdvertisementApi) GetAllAdvertisementsToBeReviewed(c *gin.Context) {
	advertisements, err := adminService.ManageAdvertisementService.GetAllAdvertisementsToBeReview()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}

	advertisementsInfo := functional.Map(advertisements, entity.AdvertisementPendingReview.GetInfo)
	resp := response.GetAdvertisementToBePreviewedResp{AdvertisementInfos: advertisementsInfo}
	jsonResp, _ := jsoniter.Marshal(resp)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

func (m *ManageAdvertisementApi) AllowAdvertisement(c *gin.Context) {
	var allowAdvertisementReq request.AllowAdvertisementReq
	if err := gin_ext.BindJSON(c, &allowAdvertisementReq); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	pendingnumber := allowAdvertisementReq.PendingNumbers
	allowAdvertisements, err := adminService.ManageAdvertisementService.AllowAdvertisement(pendingnumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
		return
	}
	err = global.GVA_DB.Where("id In ?", pendingnumber).Delete(&entity.AdvertisementPendingReview{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
		return
	}
	jsonResp, _ := jsoniter.Marshal(allowAdvertisements)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

func (m *ManageAdvertisementApi) GetAdvertisementsPendingReviewCount(c *gin.Context) {
	count, err := adminService.ManageAdvertisementService.GetAdvertisementsPendingReviewCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}
	resp := response.GetCompaniesCountResp{Count: count}
	jsonResp, _ := jsoniter.Marshal(resp)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}
func (m *ManageAdvertisementApi) DeleteAdvertisement(c *gin.Context) {
	var deleteAdvertisementReq request.DeleteAdvertisementReq
	if err := gin_ext.BindJSON(c, &deleteAdvertisementReq); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	deletenumbers := deleteAdvertisementReq.DeleteNumbers
	deleteAdvertisements, err := adminService.ManageAdvertisementService.DeleteAdvertisement(deletenumbers)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
		return
	}
	err = global.GVA_DB.Where("id In ?", deleteAdvertisements).Delete(&entity.Advertisement{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
		return
	}
	jsonResp, _ := jsoniter.Marshal(deleteAdvertisements)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

func (m *ManageAdvertisementApi) GetAdvertisementsByTerm(c *gin.Context) {
	term := c.Query("term")
	_type, err := strconv.Atoi(c.Query("type"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.InvalidParams, nil))
		return
	}
	termType := enum.AdvertisementType(_type)
	if advertisements, err := adminService.GetAdvertisementsByTerm(term, termType); err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	} else {
		advertisementsInfo := functional.Map(advertisements, entity.Advertisement.GetInfo)
		resp := response.GetAdvertisementToBePreviewedResp{AdvertisementInfos: advertisementsInfo}
		jsonResp, _ := jsoniter.Marshal(resp)
		c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
	}

}
