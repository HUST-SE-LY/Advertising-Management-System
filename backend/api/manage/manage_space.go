package manage

import (
	"backend/global"
	entity2 "backend/models/advertisement_model/entity"
	"backend/models/manage_model/request"
	"backend/models/manage_model/response"
	"backend/utils/gin_ext"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type ManageSpaceApi struct {
}

// GetSpace godoc
//
//	@Summary		Get Space
//	@Description	Get Space
//
//	@Tags			Manage
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.GetSpaceResp	"Space Info"
//	@Router			/manage/space/get [get]
func (m *ManageSpaceApi) GetSpace(c *gin.Context) {
	if err := entity2.GetFreshSpace(); err != nil {
		return
	}
	spaces, err := entity2.GetAdvertisementSpaces()
	if err != nil {
		return
	}
	resp := response.GetSpace{Space: spaces}
	jsonResp, _ := jsoniter.Marshal(resp)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

// SetSpacePrice godoc
//
//	@Summary		Set Space
//	@Description	Set Space
//
//	@Tags			Manage
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}		"Space price set"
//	@Router			/manage/space/set_price	[post]
func (m *ManageSpaceApi) SetSpacePrice(c *gin.Context) {
	var req request.SetSpacePriceReq
	if err := gin_ext.BindJSON(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	id := req.Id
	money := req.Price
	err := adminService.ManageSpaceService.SetSpacePrice(id, money)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}

	jsonResp, _ := jsoniter.Marshal(id)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

func (m *ManageSpaceApi) PutURL(c *gin.Context) {
	return
}

// AddSpace godoc
//
//	@Summary		add Space
//	@Description	add Space
//
//	@Tags			Manage
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}		"Space add"
//	@Router			/manage/space/add	[post]
func (m *ManageSpaceApi) AddSpace(c *gin.Context) {
	var req request.AddSpaceReq
	if err := gin_ext.BindJSON(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	var space = entity2.AdvertisingSpace{
		//Enddate: time.Now().Format("2023-06-01"),
		Enddate: "2023-06-01",
		Price:   req.Price,
	}
	if err := global.GVA_DB.Create(&space).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}

	jsonResp, _ := jsoniter.Marshal(space)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}
