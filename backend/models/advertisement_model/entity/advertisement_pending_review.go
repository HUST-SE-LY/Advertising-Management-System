package entity

import (
	"backend/global"
	"backend/models/advertisement_model"
	"backend/models/company_model/entity"
)

type AdvertisementPendingReview struct {
	Advertisement
}

func NewAdvertisementPendingReview(advertisement Advertisement) *AdvertisementPendingReview {
	return &AdvertisementPendingReview{
		Advertisement: advertisement,
	}
}

func (adv AdvertisementPendingReview) ToAdvertisement() Advertisement {
	return adv.Advertisement
}

func (adv AdvertisementPendingReview) GetInfo() advertisement_model.AdvertisementToBePreviewedInfo {
	var company entity.Company
	global.GVA_DB.Where("id=?", adv.CompanyId).Find(&company)
	return *advertisement_model.NewAdvertisementInfo(
		company.Account,
		adv.Id,
		adv.CompanyId,
		adv.ImageUrl,
		adv.Title,
		adv.Position,
		adv.JumpToUrl,
		adv.DisplayTime,
	)
}
