package manage

import (
	"backend/models/admin_model"
	"backend/models/admin_model/request"
	"backend/models/admin_model/response"
	"backend/service"
	"backend/utils/gin_ext"
	"backend/utils/jwt"
	"backend/utils/status"
	jsoniter "github.com/json-iterator/go"

	"github.com/gin-gonic/gin"
	"net/http"
)

type ManageAdminApi struct {
}

var adminService = service.ServiceGroupApp.ManageServiceGroup

func (m *ManageAdminApi) CreateAdmin(c *gin.Context) {
	var adminLoginReq request.AdminLoginReq
	err := c.ShouldBindJSON(&adminLoginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	admin := admin_model.Admin{
		Account:  adminLoginReq.Account,
		Password: adminLoginReq.Password,
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

	jwtToken, err := jwt.GenerateToken(admin.Account, admin.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin_ext.Response(err, nil))
		return
	}

	adminLoginResp := response.AdminLoginResp{
		Account: admin.Account,
	}

	jsonResp, _ := jsoniter.Marshal(adminLoginResp)

	c.Header("Authentication", jwtToken)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}
