package response

import entity2 "backend/models/advertisement_model/entity"

type GetSpace struct {
	Space []entity2.AdvertisingSpace `json:"space"`
}
