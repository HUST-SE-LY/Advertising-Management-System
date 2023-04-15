package request

type AdminLoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
