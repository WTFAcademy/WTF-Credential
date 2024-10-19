package errors

const (
	NormalErrorCode  = 10001
	EntityErrorCode  = 10002
	UnknownErrorCode = 10003
)

type ResultCode int

type Result struct {
	Code ResultCode  `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Unknown(err error) *Result {
	return &Result{
		Data: nil,
		Code: UnknownErrorCode,
		Msg:  err.Error(),
	}
}

func Entity(msg string) *Result {
	return &Result{
		Data: nil,
		Code: EntityErrorCode,
		Msg:  msg,
	}
}

func Normal(msg string) *Result {
	return &Result{
		Data: nil,
		Code: NormalErrorCode,
		Msg:  msg,
	}
}
