package request

type SetSpacePriceReq struct {
	Id    int64 `json:"id"`
	Price int   `json:"price"`
}
type AddSpaceReq struct {
	Price int `json:"price"`
}
