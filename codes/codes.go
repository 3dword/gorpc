package codes

// 框架状态码
const (
	OK = 0
)

type ErrorCode uint8

// 错误码类型
const (
	FrameworkError = 1
	BusinuessError = 2
)

// 框架错误码
const (
	CONFIGERROR = 201
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
		return e.Message
	}
}

func NewFrameError() {

}
