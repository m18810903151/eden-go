package serializer

import (
	"errors"
)

// AppError 应用错误，实现了error接口
type AppError struct {
	Code     int
	Msg      string
	RawError error
}

// NewError 返回新的错误对象
func NewError(code int, msg string, err error) AppError {
	return AppError{
		Code:     code,
		Msg:      msg,
		RawError: err,
	}
}

// WithError 将应用error携带标准库中的error
func (err *AppError) WithError(raw error) AppError {
	err.RawError = raw
	return *err
}

// Error 返回业务代码确定的可读错误信息
func (err AppError) Error() string {
	return err.Msg
}

// DBErr 数据库操作失败
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "Database operation failed."
	}
	return Err(CodeDBError, msg, err)
}

// ParamErr 各种参数错误
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "Invalid parameters."
	}
	return Err(CodeParamErr, msg, err)
}

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	// 底层错误是AppError，则尝试从AppError中获取详细信息
	var appError AppError
	if errors.As(err, &appError) {
		errCode = appError.Code
		err = appError.RawError
		msg = appError.Msg
	}

	return Response{
		Code: errCode,
		Msg:  msg,
	}
}

const (
	// IdNotExists id不存在
	IdNotExists    = 10001
	CodeCheckLogin = 401
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
)
