package errors

type BizError struct {
	Code    string
	Message string
}

func (e *BizError) Error() string {
	return e.Message
}

func NewBizErr(code, msg string) *BizError {
	return &BizError{Code: code, Message: msg}
}
