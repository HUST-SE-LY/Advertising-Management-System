package response

import (
	"backend/models/advertisement_model"
)

type GetAdvertisementsResp struct {
	AdvertisementInfos []advertisement_model.AdvertisementInfo `json:"advertisement_infos"`
}

type GetAdvertisementToBePreviewedResp struct {
	AdvertisementInfos []advertisement_model.AdvertisementpreviewedInfo `json:"advertisement_infos"`
}
