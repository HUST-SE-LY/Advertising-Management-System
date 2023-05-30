package request

type AdminLoginReq struct {
	Account  string `json:"account" example:"fnfunfunc"`
	Password string `json:"password" example:"1234"`
}
