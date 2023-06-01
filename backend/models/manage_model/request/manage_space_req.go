package request

type SetSpacePriceReq struct {
	Id    int64 `json:"id"`
	Price int   `json:"price"`
}
