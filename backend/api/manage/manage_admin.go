package manage

import (
	"backend/models/admin_model"
	"backend/models/manage_model/request"
	"backend/models/manage_model/response"
	"backend/utils/gin_ext"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type ManageAdminApi struct {
}

func (m *ManageAdminApi) CreateAdmin(c *gin.Context) {
	var adminRegisterReq request.AdminCreateReq
	err := c.ShouldBindJSON(&adminRegisterReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	var admin *admin_model.Admin
	if admin, err = adminService.CreateAdmin(&adminRegisterReq); err != nil {
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

	adminLoginResp := response.AdminCreateResp{
		Account: admin.Account,
	}

	jsonResp, _ := jsoniter.Marshal(adminLoginResp)

	//c.Header("Authentication", jwtToken)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))
}
