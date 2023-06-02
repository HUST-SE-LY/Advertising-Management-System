package manage

import (
	"backend/global"
	"backend/models/advertisement_model/entity"
	entity2 "backend/models/company_model/entity"
	"backend/models/manage_model/enum"
	entity3 "backend/models/record_model/entity"
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
	var record entity3.ConsumeRecord
	if err = global.GVA_DB.Where("id IN ?", pendingnumbers).Find(&advertisementToBeReviewed).Error; err != nil {
		return nil, err
	}
	for _, s := range advertisementToBeReviewed {
		if err = global.GVA_DB.Where("position =?", s.Position).Where("end = ?", s.EndDate).Find(&record).Error; err != nil {
			return
		}
		record.Status = 1
		global.GVA_DB.Model(&record).Updates(record)
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
func (m *ManageAdvertisementService) RejectAdvertisement(reject_numbers []int64) (err error) {
	var advertisementToBeReviewed []entity.AdvertisementPendingReview
	var record entity3.ConsumeRecord
	var company entity2.Company
	if err = global.GVA_DB.Where("id IN ?", reject_numbers).Find(&advertisementToBeReviewed).Error; err != nil {
		return err
	}

	for _, s := range advertisementToBeReviewed {
		if err = global.GVA_DB.Where("position =?", s.Position).Where("end = ?", s.EndDate).Find(&record).Error; err != nil {
			return
		}
		record.Status = 2
		global.GVA_DB.Model(&record).Updates(record)
		global.GVA_DB.Where("account=?", record.Account).Find(&company)
		company.Balance = company.Balance + record.Cost
		global.GVA_DB.Model(&company).Updates(company)
	}
	if err = global.GVA_DB.Delete(&advertisementToBeReviewed).Error; err != nil {
		return err
	}

	return nil
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

func (m *ManageAdvertisementService) StopAdvertisement(column int64) error {
	var space entity.AdvertisingSpace
	var record entity3.ConsumeRecord
	var advertisement entity.Advertisement
	var com entity2.Company
	format := "2006-01-02"
	if err := global.GVA_DB.Where("id=?", column).Find(&space).Error; err != nil {
		return err
	}
	global.GVA_DB.Where("id=?", space.AdvID).Find(&advertisement)

	if err := global.GVA_DB.Where("id=?", advertisement.CompanyId).Find(&com).Error; err != nil {
		return err
	}

	if err := global.GVA_DB.Where("position =?", int(space.Id)).Where("end = ?", advertisement.EndDate).Find(&record).Error; err != nil {
		return err
	}
	record.Status = 4
	global.GVA_DB.Model(&record).Updates(record)
	data1, err := time.ParseInLocation(format, "2023-06-02", time.Local)
	if err != nil {
		return err
	}
	data2, _ := time.ParseInLocation(format, advertisement.EndDate, time.Local)
	data3, _ := time.ParseInLocation(format, record.Start, time.Local)
	data4, _ := time.ParseInLocation(format, record.End, time.Local)
	var cost = float32(int(data2.Sub(data1).Hours()/24)) / float32(int(data4.Sub(data3).Hours()/24))
	cost = float32(record.Cost) * cost
	com.Balance = com.Balance + int(cost)
	if err = global.GVA_DB.Model(&com).Updates(com).Error; err != nil {
		return err
	}
	space.AdvID = 1
	global.GVA_DB.Model(&record).Updates(record)
	if err = global.GVA_DB.Model(&space).Updates(space).Error; err != nil {
		return err
	}
	return nil
}
func (m *ManageAdvertisementService) GetAdvertisementsPendingReviewCount() (int64, error) {
	var ads []entity.AdvertisementPendingReview
	if err := global.GVA_DB.Find(&ads).Error; err != nil {
		return 0, err
	}
	return int64(len(ads)), nil
}

func DeleteFile(filename string) error {
	realfilename := filename[15:]
	if err := os.Remove(realfilename); err != nil {
		return err
	}
	return nil
}

//func GetDaysBetween2Date(format, date1Str, date2Str string) (int, error) {
//	date1, err := time.ParseInLocation(format, date1Str, time.Local)
//	if err != nil {
//		return 0, err
//	}
//	// 将字符串转化为Time格式
//	date2, err := time.ParseInLocation(format, date2Str, time.Local)
//	if err != nil {
//		return 0, err
//	}
//	//计算相差天数
//	return int(date1.Sub(date2).Hours() / 24), nil
//}
//func costbalance(spaceid int64, com int64, startdate, enddate string) error {
//	var company entity2.Company
//	var space entity.AdvertisingSpace
//	if err := global.GVA_DB.Where("id=?", com).Find(&company).Error; err != nil {
//		return err
//	}
//	if err := global.GVA_DB.Where("id=?", spaceid).Find(&space).Error; err != nil {
//		return err
//	}
//	format := "2006-01-02"
//	e, err := GetDaysBetween2Date(format, enddate, startdate)
//	if err != nil {
//		return err
//	}
//	company.Balance = company.Balance + (space.Price * e)
//	global.GVA_DB.Model(&company).Updates(company)
//
//	return nil
//}
