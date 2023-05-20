package manage

import (
	"backend/models/advertisement_model/entity"
	"backend/models/manage_model/response"
	"backend/utils/functional"
	"backend/utils/gin_ext"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type ManageAdvertisementApi struct {
}

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

func (m *ManageAdvertisementApi) GetAllAdvertisementsToBeReviewed(c *gin.Context) {
	// TODO
}

func (m *ManageAdvertisementApi) AllowAdvertisement(c *gin.Context) {
	//var allowAdvertisementReq request.AllowAdvertisementReq
	//if err := gin_ext.BindJSON(c, &allowAdvertisementReq); err != nil {
	//	c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
	//	return
	//}

}
