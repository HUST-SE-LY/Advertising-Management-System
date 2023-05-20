package manage

import (
	"backend/global"
	"backend/models/advertisement_model/entity"
)

type ManageAdvertisementService struct {
}

func (m *ManageAdvertisementService) GetAllAdvertisements() (advertisements []entity.Advertisement, err error) {
	if err = global.GVA_DB.Find(&advertisements).Error; err != nil {
		return nil, err
	}
	return advertisements, nil

}
func (m *ManageAdvertisementService) GetAllAdvertisementsToBeReview() (advertisements []entity.Advertisement) {
	//if err = global.GVA_DB.Find()
	return
}
