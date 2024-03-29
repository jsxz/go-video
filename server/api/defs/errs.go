package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}
type ErrorResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC: 400, Error: Err{Error: "request body not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrorResponse{HttpSC: 401, Error: Err{Error: "user authentication failed", ErrorCode: "002"}}
	ErrorDBError                = ErrorResponse{HttpSC: 500, Error: Err{Error: "db ops  failed", ErrorCode: "003"}}
	ErrorInternalFaults         = ErrorResponse{HttpSC: 500, Error: Err{Error: "error interanl service  failed", ErrorCode: "004"}}
)
