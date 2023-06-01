package manage

import (
	"backend/global"
	"backend/models/advertisement_model/entity"
)

type ManageSpaceService struct {
}

func (m *ManageSpaceService) SetSpacePrice(spaceid int64, money int) (err error) {
	var space entity.AdvertisingSpace
	if err = global.GVA_DB.Where("id = ?", spaceid).Find(&space).Error; err != nil {
		return
	}
	space.Id = spaceid
	space.Price = money
	global.GVA_DB.Model(&space).Updates(space)
	return nil
}
