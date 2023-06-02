package company

import (
	"backend/models/company_model/request"
	"backend/models/company_model/response"
	"backend/utils/gin_ext"
	"backend/utils/jwt"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type CompanyAdvertisementApi struct {
}

// CompanyUploadAdvertisement godoc
//
//	@Summary		Company upload advertisement
//	@Description	Company upload advertisement
//
//	@Tags			Company
//	@Accept			mpfd
//	@Produce		json
//	@Param			advertisementInfo	body		request.CompanyUploadAdvtReq.AdvtInfo	true	"name: advertisementInfo"
//	@Param			image				body		byte									true	"name: image"
//	@Success		200					{object}	nil
//	@Router			/company/advt/upload [put]
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

// GetRecord godoc
//
//	@Summary	Company records
//
//	@Tags		Company
//	@Produce	json
//	@Success	200	{object}	"All records"
//	@Router		/company/get-records    [get]
func (com *CompanyAdvertisementApi) GetRecord(c *gin.Context) {
	var getrecordReq request.CompanyGetRecordReq
	if err := c.Bind(&getrecordReq); err != nil {
		c.JSON(http.StatusBadRequest, gin_ext.Response(status.ParseJsonError, nil))
		return
	}
	token := c.Request.Header.Get("Authorization")
	claims, _ := jwt.ParseToken(token)
	account := claims.Username
	records, err := companyService.CompanyAdvertisementService.GetCompanyRecord(account)
	if err != nil {
		return
	}
	resp := response.CompanyGetRecordsResp{Company_records: records}
	jsonResp, _ := jsoniter.Marshal(resp)
	c.JSON(http.StatusOK, gin_ext.Response(nil, string(jsonResp)))

}
