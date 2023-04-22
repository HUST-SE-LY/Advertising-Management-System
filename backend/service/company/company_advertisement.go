package company

import (
	"backend/global"
	entity2 "backend/models/advertisement_model/entity"
	"backend/models/company_model/entity"
	"backend/models/company_model/request"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type CompanyAdvertisementService struct {
}

func (a *CompanyAdvertisementService) CompanyUploadAdvertisement(req request.CompanyUploadAdvtReq, account string) error {
	// TODO
	var company entity.Company
	saveFile(req.FileData, req.FileData.Filename)
	// Find CompanyID
	global.GVA_DB.Where("account = ?", account).First(&company)
	CompanyId := company.Id
	err := global.GVA_DB.Create(&entity2.Advertisement{
		CompanyId:         CompanyId,
		AdvertisementInfo: req.AdvtInfo,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func saveFile(file *multipart.FileHeader, fileName string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	dst := *global.GVA_FILE_SETTING + "/" + fileName
	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
