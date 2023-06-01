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
	"strings"
	"time"
)

type CompanyAdvertisementService struct {
}

func (a *CompanyAdvertisementService) CompanyUploadAdvertisement(req request.CompanyUploadAdvtReq, account string) error {
	// TODO
	var company entity.Company
	filename := account + "_" + req.AdvtInfo.Title + "_" + time.Now().Format(time.RFC3339) + getFileExtension(req.FileData.Filename)
	r := strings.NewReplacer(" ", "_", ":", "_")
	filename = r.Replace(filename)
	if err := saveFile(req.FileData, filename); err != nil {
		return err
	}

	// Find CompanyID
	global.GVA_DB.Where("account = ?", account).First(&company)
	CompanyId := company.Id
	ImageUrl := global.GVA_APP_SETTING.Domain + "/" + *global.GVA_FILE_SETTING + "/" + filename
	advertisementToBeReviewed := entity2.NewAdvertisementPendingReview(
		*entity2.NewAdvertisementWithoutId(CompanyId, ImageUrl, req.AdvtInfo),
	)
	err := global.GVA_DB.Create(&advertisementToBeReviewed).Error
	if err != nil {
		return err
	}
	//if err := SaveConsume(account, req); err != nil {
	//	return err
	//}
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

//func SaveConsume(account string, req request.CompanyUploadAdvtReq) error {
//	format := "2006-01-02"
//	var space entity2.AdvertisingSpace
//	var company entity.Company
//
//	date1, err := time.ParseInLocation(format, req.AdvtInfo.StartDate, time.Local)
//	if err != nil {
//		return err
//	}
//	// 将字符串转化为Time格式
//	date2, err := time.ParseInLocation(format, req.AdvtInfo.EndDate, time.Local)
//	if err != nil {
//		return err
//	}
//	global.GVA_DB.Where("id =?", int64(req.AdvtInfo.Position)).Find(&space)
//	var cost = space.Price * int(date1.Sub(date2).Hours()/24)
//	consumerecord := entity3.ConsumeRecord{
//		Account:   account,
//		Startdate: req.AdvtInfo.StartDate,
//		Enddate:   req.AdvtInfo.EndDate,
//		Cost:      cost,
//		Status:    enum.UnderReview,
//	}
//	global.GVA_DB.Where("account=?", account).Find(&company)
//	if err := global.GVA_DB.Create(&consumerecord).Error; err != nil {
//		return err
//	}
//	company.Balance = company.Balance - cost
//	global.GVA_DB.Model(&company).Updates(company)
//	return nil
//}
