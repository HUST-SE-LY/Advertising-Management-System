package status

type StatusCode int

type Status struct {
	Code StatusCode
	Msg  string
}

func NewStatus(code StatusCode, msg string) Status {
	return Status{Code: code, Msg: msg}
}

func ErrorToStatus(err error) Status {
	return NewStatus(ERROR, err.Error())
}

const (
	SUCCESS StatusCode = 200

	INVALID_PARAMS StatusCode = 400

	ERROR StatusCode = 500

	PARSE_JSON_ERROR StatusCode = 600

	SAME_ACCOUNT_EXISTS StatusCode = 700

	ERROR_AUTH_CHECK_TOKEN_FAIL    StatusCode = 800
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT StatusCode = 801
)

// Default Status
var (
	Success       = Status{Code: SUCCESS, Msg: "OK"}
	InvalidParams = Status{Code: INVALID_PARAMS, Msg: "Illegal request parameters"}
	Error         = Status{Code: ERROR, Msg: "Fail"}

	ParseJsonError = Status{Code: PARSE_JSON_ERROR, Msg: "Error while parsing json"}

	SameAccountExists = Status{Code: SAME_ACCOUNT_EXISTS, Msg: "Same account exists"}

	ErrorAuthCheckTokenFail    = Status{Code: ERROR_AUTH_CHECK_TOKEN_FAIL, Msg: "Token authentication failed"}
	ErrorAuthCheckTokenTimeout = Status{Code: ERROR_AUTH_CHECK_TOKEN_TIMEOUT, Msg: "Token has expired"}
)

func (st Status) Error() string {
	return st.Msg
}
