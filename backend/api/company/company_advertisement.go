package company

import (
	"backend/models/company_model/request"
	"backend/utils/gin_ext"
	"backend/utils/jwt"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CompanyAdvertisementApi struct {
}

func (com *CompanyAdvertisementApi) CompanyUploadAdvertisement(c *gin.Context) {
	var companyUploadAdvtReq request.CompanyUploadAdvtReq

	if err := c.Bind(&companyUploadAdvtReq); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}

	token := c.Request.Header.Get("Authorization")
	claims, _ := jwt.ParseToken(token)
	account := claims.Username
	err := companyService.CompanyUploadAdvertisement(companyUploadAdvtReq, account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(err, nil))
		return
	}
	c.JSON(http.StatusOK, gin_ext.Response(nil, nil))
}