package request

type AllowAdvertisementReq struct {
	PendingNumbers []int64 `json:"pending_numbers"`
}
