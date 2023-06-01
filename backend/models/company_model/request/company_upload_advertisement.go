package request

import (
	"backend/models/advertisement_model"
	"mime/multipart"
)

type CompanyUploadAdvtReq struct {
	AdvtInfo advertisement_model.AdvertisementInfo `form:"advertisementInfo"`
	FileData *multipart.FileHeader                 `form:"image" swaggerignore:"true"`
}
