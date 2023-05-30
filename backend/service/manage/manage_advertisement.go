package manage

import (
	"backend/global"
	"backend/models/advertisement_model/entity"
	"backend/models/manage_model/enum"
	"backend/utils/functional"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"os"
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

func DeleteFile(filename string) error {
	realfilename := filename[15:]
	if err := os.Remove(realfilename); err != nil {
		return err
	}
	return nil
}
