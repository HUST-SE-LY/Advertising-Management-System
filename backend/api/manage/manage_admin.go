package manage

import (
	"backend/models/admin_model"
	"backend/models/admin_model/request"
	"backend/models/admin_model/response"
	"backend/utils/gin_ext"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type ManageAdminApi struct {
}

func (m *ManageAdminApi) CreateAdmin(c *gin.Context) {
	var adminRegisterReq request.AdminRegisterReq
	err := c.ShouldBindJSON(&adminRegisterReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	admin := admin_model.Admin{
		Account:  adminRegisterReq.Account,
		Password: adminRegisterReq.Password,
	}

	if err = adminService.CreateAdmin(&admin); err != nil {
		if st, ok := err.(status.Status); ok {
			if st.Code == status.SAME_ACCOUNT_EXISTS {
				c.JSON(http.StatusBadRequest, gin_ext.Response(st, nil))
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}

	//jwtToken, err := jwt.GenerateToken(admin.Account, admin.Password, admin_model.ADMIN_ROLE)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
	//	return
	//}

	adminLoginResp := response.AdminRegisterResp{
		Account: admin.Account,
	}

	jsonResp, _ := jsoniter.Marshal(adminLoginResp)

	//c.Header("Authentication", jwtToken)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}

func (m *ManageAdminApi) AdminLogin(c *gin.Context) {
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
		c.Header("Authentication", adminToken.Token)
		c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
	}
}
