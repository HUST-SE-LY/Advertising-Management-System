package manage

import (
	"backend/global"
	"backend/models/advertisement_model/entity"
	"backend/utils/functional"
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
