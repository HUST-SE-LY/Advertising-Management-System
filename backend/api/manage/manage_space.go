package manage

import (
	"backend/models/manage_model/request"
	"backend/utils/gin_ext"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type ManageSpaceApi struct {
}

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

func (m *ManageSpaceApi) AddSpace(c *gin.Context) (err error) {
	return nil
}
