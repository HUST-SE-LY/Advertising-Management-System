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
	"time"
)

type CompanyAdvertisementService struct {
}

func (a *CompanyAdvertisementService) CompanyUploadAdvertisement(req request.CompanyUploadAdvtReq, account string) error {
	// TODO
	var company entity.Company
	filename := account + "_" + req.AdvtInfo.Title + "_" + time.Now().Format(time.RFC3339) + getFileExtension(req.FileData.Filename)
	if err := saveFile(req.FileData, filename); err != nil {
		return err
	}

	// Find CompanyID
	global.GVA_DB.Where("account = ?", account).First(&company)
	CompanyId := company.Id
	err := global.GVA_DB.Create(&entity2.Advertisement{
		CompanyId:         CompanyId,
		ImageUrl:          global.GVA_APP_SETTING.Address + "/" + *global.GVA_FILE_SETTING + filename,
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
	// ignore
	defer func(src multipart.File) {
		_ = src.Close()
	}(src)
	dst := *global.GVA_FILE_SETTING + "/" + fileName
	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	// ignore
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	_, err = io.Copy(out, src)
	return err
}

func getFileExtension(filename string) string {
	return filepath.Ext(filename)
}
