package manage

import (
	"backend/global"
	"backend/models/advertisement_model/entity"
	entity2 "backend/models/company_model/entity"
	"backend/models/manage_model/enum"

	"backend/utils/functional"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"os"
	"time"
)

type ManageAdvertisementService struct {
}

func (m *ManageAdvertisementService) GetAllAdvertisements() (advertisements []entity.Advertisement, err error) {
	if err = global.GVA_DB.Find(&advertisements).Error; err != nil {
		return nil, err
	}
	return advertisements, nil

}

func (m *ManageAdvertisementService) GetAllAdvertisementsToBeReview() (advertisementstobepreview []entity.AdvertisementPendingReview, err error) {
	if err = global.GVA_DB.Find(&advertisementstobepreview).Error; err != nil {
		return nil, err
	}

	return advertisementstobepreview, nil
}

func (m *ManageAdvertisementService) AllowAdvertisement(pendingnumbers []int64) (addedadvnumber []int64, err error) {
	var advertisementToBeReviewed []entity.AdvertisementPendingReview
	if err = global.GVA_DB.Where("id IN ?", pendingnumbers).Find(&advertisementToBeReviewed).Error; err != nil {
		return nil, err
	}
	for _, s := range advertisementToBeReviewed {
		costbalance(int64(s.Position), s.CompanyId, s.EndDate, s.StartDate)
	}
	advertisements := functional.Map(advertisementToBeReviewed, entity.AdvertisementPendingReview.ToAdvertisement)
	advertisementlen := len(advertisements)
	if err = global.GVA_DB.CreateInBatches(&advertisements, advertisementlen).Error; err != nil {
		return nil, err
	}
	return functional.Map(advertisements, func(adv entity.Advertisement) int64 {
		return adv.Id
	}), nil

}

func (m *ManageAdvertisementService) DeleteAdvertisement(deleteid []int64) (deleted_id []int64, err error) {
	var advertisements []entity.Advertisement
	if err = global.GVA_DB.Where("id IN ?", deleteid).Find(&advertisements).Error; err != nil {
		return nil, err
	}
	for _, ad := range advertisements {
		DeleteFile(ad.ImageUrl)
	}
	return deleteid, nil
}

func (m *ManageAdvertisementService) GetAdvertisementsByTerm(term string, termType enum.AdvertisementType) (advertisements []entity.Advertisement, err error) {
	whereCondition := fmt.Sprintf("%s LIKE ?", termType.ToString())
	term = fmt.Sprintf("%%%s%%", term)
	if err = global.GVA_DB.Where(whereCondition, term).Find(&advertisements).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Advertisement{}, nil
		}
		return nil, err
	}

	return advertisements, nil
}
func (m *ManageAdvertisementService) GetAdvertisementsPendingReviewCount() (int, error) {
	var ads []entity.AdvertisementPendingReview
	if err := global.GVA_DB.Find(&ads).Error; err != nil {
		return 0, err
	}
	return len(ads), nil
}

func DeleteFile(filename string) error {
	realfilename := filename[15:]
	if err := os.Remove(realfilename); err != nil {
		return err
	}
	return nil
}

func GetDaysBetween2Date(format, date1Str, date2Str string) (int, error) {
	date1, err := time.ParseInLocation(format, date1Str, time.Local)
	if err != nil {
		return 0, err
	}
	// 将字符串转化为Time格式
	date2, err := time.ParseInLocation(format, date2Str, time.Local)
	if err != nil {
		return 0, err
	}
	//计算相差天数
	return int(date1.Sub(date2).Hours() / 24), nil
}
func costbalance(spaceid int64, com int64, startdate, enddate string) error {
	var company entity2.Company
	var space entity.AdvertisingSpace
	if err := global.GVA_DB.Where("id=?", com).Find(&company).Error; err != nil {
		return err
	}
	if err := global.GVA_DB.Where("id=?", spaceid).Find(&space).Error; err != nil {
		return err
	}
	format := "2006-01-02"
	e, err := GetDaysBetween2Date(format, enddate, startdate)
	if err != nil {
		return err
	}
	company.Balance = company.Balance + (space.Price * e)
	global.GVA_DB.Model(&company).Updates(company)

	return nil
}
