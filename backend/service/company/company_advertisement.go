package company

import (
	"backend/global"
	"backend/models/advertisement_model"
	entity2 "backend/models/advertisement_model/entity"
	"backend/models/company_model/entity"
	"backend/models/company_model/request"
	entity3 "backend/models/record_model/entity"
	jsoniter "github.com/json-iterator/go"
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
	var AdvtInfo advertisement_model.AdvertisementInfo
	err := jsoniter.Unmarshal([]byte(req.AdvtInfo), &AdvtInfo)
	if err != nil {
		return err
	}
	filename := account + "_" + AdvtInfo.Title + "_" + time.Now().Format(time.RFC3339) + getFileExtension(req.FileData.Filename)
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
		*entity2.NewAdvertisementWithoutId(CompanyId, ImageUrl, AdvtInfo),
	)
	Id := int64(AdvtInfo.Position)
	var space entity2.AdvertisingSpace
	if err := global.GVA_DB.Where("id=?", Id).Find(&space).Error; err != nil {
		return err
	}
	space.Enddate = AdvtInfo.EndDate
	global.GVA_DB.Model(&space).Updates(space)
	err = global.GVA_DB.Create(&advertisementToBeReviewed).Error
	if err != nil {
		return err
	}
	if err := SaveConsume(account, AdvtInfo); err != nil {
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

func SaveConsume(account string, AdvtInfo advertisement_model.AdvertisementInfo) error {
	format := "2006-01-02"
	var space entity2.AdvertisingSpace
	var company entity.Company

	date1, err := time.ParseInLocation(format, AdvtInfo.StartDate, time.Local)
	if err != nil {
		return err
	}
	// 将字符串转化为Time格式
	date2, err := time.ParseInLocation(format, AdvtInfo.EndDate, time.Local)
	if err != nil {
		return err
	}
	global.GVA_DB.Where("id =?", int64(AdvtInfo.Position)).Find(&space)
	var cost = space.Price * int(date2.Sub(date1).Hours()/24)
	currentTime := time.Now().Format(format)

	consumerecord := entity3.ConsumeRecord{
		Account:  account,
		Start:    AdvtInfo.DisplayTime.StartDate,
		End:      AdvtInfo.DisplayTime.EndDate,
		Position: AdvtInfo.Position,
		Cost:     cost,
		Status:   0,
		Date:     currentTime,
	}
	global.GVA_DB.Where("account=?", account).Find(&company)
	if err := global.GVA_DB.Create(&consumerecord).Error; err != nil {
		return err
	}
	company.Balance = company.Balance - cost
	global.GVA_DB.Model(&company).Updates(company)
	return nil
}

func (a *CompanyAdvertisementService) GetCompanyRecord(account string) (Records []entity3.ConsumeRecord, err error) {
	var records []entity3.ConsumeRecord
	if err = global.GVA_DB.Where("account=?", account).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}
