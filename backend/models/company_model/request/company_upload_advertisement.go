package request

import (
	"mime/multipart"
)

type CompanyUploadAdvtReq struct {
	AdvtInfo string                `form:"advertisementInfo"`
	FileData *multipart.FileHeader `form:"image" swaggerignore:"true"`
}
