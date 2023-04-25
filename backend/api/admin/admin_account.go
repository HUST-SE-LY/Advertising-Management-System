package admin

import (
	"backend/models/admin_model/request"
	"backend/utils/gin_ext"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type AdminAccountApi struct {
}

func (a *AdminAccountApi) AdminLogin(c *gin.Context) {
	var adminLoginReq request.AdminLoginReq
	err := c.ShouldBindJSON(&adminLoginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	if err, admin, adminToken := adminService.AdminLogin(&adminLoginReq); err != nil {
		c.JSON(http.StatusUnauthorized, gin_ext.Response(err, nil))
	} else {
		jsonResp, _ := jsoniter.Marshal(admin)
		c.Header("Authorization", adminToken.Token)
		c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
	}
}
