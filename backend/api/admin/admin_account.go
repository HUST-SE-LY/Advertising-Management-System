package admin

import (
	"backend/models/admin_model/request"
	"backend/models/admin_model/response"
	"backend/utils/gin_ext"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type AdminAccountApi struct {
}

// AdminLogin godoc
//
//	@Summary		Admin login
//	@Description	Admin Login by account and password
//
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			account		body		string	true	"account"
//	@Param			password	body		string	true	"password"
//	@Success		200			{object}	string	"TODO"
//	@Router			/admin/login [post]
func (a *AdminAccountApi) AdminLogin(c *gin.Context) {
	var adminLoginReq request.AdminLoginReq
	err := c.ShouldBindJSON(&adminLoginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	if err, admin, adminToken := adminService.AdminLogin(&adminLoginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
	} else {
		jsonResp, _ := jsoniter.Marshal(response.AdminLoginResp{Account: admin.Account, Token: adminToken.Token})
		c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
	}
}
