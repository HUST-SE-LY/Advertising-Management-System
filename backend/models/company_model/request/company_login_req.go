package request

type CompanyLoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
