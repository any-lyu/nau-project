package apierror

import (
	"encoding/json"
	"fmt"

	"family-joint-school/common/errors/errorcode"
)

// Error 表示 api 级别的错误.
type Error interface {
	// ErrorCode 返回错误码
	ErrorCode() int
	// ErrorMessage 返回错误描述
	ErrorMessage() string
}

// New 返回一个实现了 Error 的 error.
//  如果 code == errorcode.Code_OK, 返回 nil.
func New(code errorcode.Code, msg string) error {
	if code == errorcode.Code_OK {
		return nil
	}
	return &apiError{
		Code:    code,
		Message: msg,
	}
}

type apiError struct {
	Code    errorcode.Code `json:"code"`
	Message string         `json:"msg"`
}

var _ error = (*apiError)(nil)

func (e *apiError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("(code:%d,msg:%s)", e.Code, e.Message)
	}
	return string(b)
}

var _ Error = (*apiError)(nil)

func (e *apiError) ErrorCode() int       { return int(e.Code) }
func (e *apiError) ErrorMessage() string { return e.Message }

func (e *apiError) Is(target error) bool {
	apiError, ok := target.(*apiError)
	if !ok || apiError == nil {
		return false
	}
	return e.ErrorCode() == apiError.ErrorCode()
}

// CodeIs 判断 err 是否是 New 返回的并且 code==target
func CodeIs(err error, target errorcode.Code) bool {
	apiError, ok := err.(*apiError)
	if !ok || apiError == nil {
		return false
	}
	return apiError.ErrorCode() == int(target)
}

// TryGetErrCode 尝试获取 error code, 如果 err 不是 *apiError 或者 err 是 nil, 则返回 0.
func TryGetErrCode(err error) errorcode.Code {
	apiError, ok := err.(*apiError)
	if ok && apiError != nil {
		return apiError.Code
	}
	return 0
}
