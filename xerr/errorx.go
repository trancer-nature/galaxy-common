package xerr

import (
	"fmt"
	"go.opentelemetry.io/otel/trace"
)

/**
常用通用固定错误
*/
const DefaultCode uint32 = 1

var message map[uint32]string

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

type CodeError struct {
	errCode uint32
	errMsg  string
}

//返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

//返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}
func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}

func GetSpanID(span trace.SpanContext) string {
	if !span.IsValid() {
		return ""
	}

	if !span.HasSpanID() {
		return ""
	}

	return span.SpanID().String()
}

func GetTraceID(span trace.SpanContext) string {
	if !span.IsValid() {
		return ""
	}

	if !span.HasTraceID() {
		return ""
	}

	return span.TraceID().String()
}
