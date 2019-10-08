package codes

import "fmt"

// 框架状态码
const (
	OK = 0
	ConfigErrorCode = 101
	FrameworkInitErrorCode = 102
	ServerAddressErrorCode = 103
	ServerEncodeErrorCode = 104
	ServerDecodeErrorCode = 105
)

type ErrorCode uint8

// 错误码类型
const (
	FrameworkError = 1
	BusinuessError = 2
)


// 框架错误码
var (
	ConfigError = NewFrameworkError(ConfigErrorCode,"config error")
	FrameworkInitError = NewFrameworkError(FrameworkInitErrorCode, "framework init error")
	ServerAddressError = NewFrameworkError(ServerAddressErrorCode, "framework init error")
	ServerEncodeError = NewFrameworkError(ServerEncodeErrorCode, "server encode error")
	ServerDecodeError = NewFrameworkError(ServerDecodeErrorCode, "server decode error")
)


// 业务错误码
const (

)

type Error struct {
	Code int
	Type int32
	Message string
}

const (
	Success = "success"
)

func (e *Error) Error() string {
	if e == nil {
		return Success
	}
	if e.Type == FrameworkError {
		return fmt.Sprintf("type : framework, code : %d, msg : %s",e.Code, e.Message)
	}
	return fmt.Sprintf("type : business, code : %d, msg : %s",e.Code, e.Message)
}

func NewFrameworkError(code int, msg string) *Error{
	return &Error{
		Type : FrameworkError,
		Code : code,
		Message : msg,
	}
}

// 方便业务使用，默认是业务类型错误
func New(code int, msg string) *Error{
	return &Error{
		Type : BusinuessError,
		Code : code,
		Message : msg,
	}
}
