package entity

import (
	"backend/models/advertisement_model"
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
	return *advertisement_model.NewAdvertisementInfo(
		adv.Id,
		adv.CompanyId,
		adv.ImageUrl,
		adv.Title,
		adv.Position,
		adv.JumpToUrl,
		adv.DisplayTime,
	)
}
