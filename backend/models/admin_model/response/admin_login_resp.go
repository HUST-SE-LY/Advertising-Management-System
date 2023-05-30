package response

type AdminLoginResp struct {
	Account string `json:"account"`
	Token   string `json:"token"`
}
