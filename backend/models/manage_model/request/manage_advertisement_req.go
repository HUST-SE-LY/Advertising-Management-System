package request

type AllowAdvertisementReq struct {
	PendingNumbers []int64 `json:"pending_numbers"`
}

type DeleteAdvertisementReq struct {
	DeleteNumbers []int64 `json:"delete_numbers"`
}
