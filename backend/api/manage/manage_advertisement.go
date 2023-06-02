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

// AllowAdvertisement godoc
//
//	@Summary	Get all advertisements to be allowed
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]string	"All advertisements"
//	@Router		/manage/advertisement/allow [post]
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

// RejectAdvertisement godoc
//
//	@Summary	Get all advertisements to be rejected
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]int64	"All advertisements"
//	@Router		/manage/advertisement/reject [post]
func (m *ManageAdvertisementApi) RejectAdvertisement(c *gin.Context) {
	var rejectAdvertisementReq request.RejectAdvertisementReq
	if err := gin_ext.BindJSON(c, &rejectAdvertisementReq); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	rejectnumbers := rejectAdvertisementReq.RejectNumbers
	err := adminService.ManageAdvertisementService.RejectAdvertisement(rejectnumbers)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
		return
	}
	adminService.ManageAdvertisementService.DeleteAdvertisement(rejectnumbers)

	jsonResp, _ := jsoniter.Marshal(rejectnumbers)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

// GetAdvertisementsPendingReviewCount godoc
//
//	@Summary	Get all advertisements to be reviewed
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	int	"All advertisements"
//	@Router		/manage/advertisement/count [get]
func (m *ManageAdvertisementApi) GetAdvertisementsPendingReviewCount(c *gin.Context) {
	count, err := adminService.ManageAdvertisementService.GetAdvertisementsPendingReviewCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}

	resp := response.GetAdvertisementsCountResp{Count: count}
	jsonResp, _ := jsoniter.Marshal(resp)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

// DeleteAdvertisement godoc
//
//	@Summary	Get all advertisements to be reviewed
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]int64	"All advertisements"
//	@Router		/manage/advertisement/delete [post]
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

// GetAdvertisementsByTerm godoc
//
//	@Summary	Get all advertisements by term
//
//	@Tags		Manage
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.GetAdvertisementToBePreviewedResp	"All advertisements"
//	@Router		/manage/advertisement/search [get]
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

func (m *ManageAdvertisementApi) StopAdvertisement(c *gin.Context) {
	var req request.StopAdvertisementReq
	var numbers []int64
	if err := gin_ext.BindJSON(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	deletenumber := req.StopNumber

	number, err := adminService.ManageAdvertisementService.StopAdvertisement(deletenumber)
	numbers[0] = number
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}
	adminService.DeleteAdvertisement(numbers)
	jsonResp, _ := jsoniter.Marshal(deletenumber)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}
