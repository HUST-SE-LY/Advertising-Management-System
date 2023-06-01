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
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			request_body	body		request.AdminLoginReq	true	"Admin login request"
//	@Success		200				{object}	response.AdminLoginResp	"Admin login response"
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

// AdminLogout godoc
//
//	@Summary		Admin logout
//	@Description	Company logout
//
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	nil	"response"
//	@Router			/admin/logout [post]
func (a *AdminAccountApi) AdminLogout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	err := adminService.AdminLogout(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.LogoutError, nil))
		return
	}
	c.JSON(http.StatusOK, gin_ext.Response(nil, nil))
}

// dummy
func (a *AdminAccountApi) CheckLogin(c *gin.Context) {

}
