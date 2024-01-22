package xlog

import (
	"bytes"
	"github.com/trancer-nature/galaxy-common/util/format"
	"net/http"
	"time"
)

type LogWriter struct {
	http.ResponseWriter
	Body *bytes.Buffer
}

// OptLog 定义了操作日志的结构
type OptLog struct {
	User       string `json:"user,omitempty"`
	OpTime     string `json:"op_time,omitempty"`
	OpType     string `json:"op_type,omitempty"`
	IP         string `json:"ip,omitempty"`
	Company    string `json:"company,omitempty"`
	Permission string `json:"permission,omitempty"`
	Method     string `json:"method,omitempty"`
	URL        string `json:"url,omitempty"`
	Module     string `json:"module,omitempty"`
	Param      string `json:"param,omitempty"`
	Trace      string `json:"trace,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	Result     Result `json:"result,omitempty"`
}

type Result struct {
	Code int32       `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
}

// NewOptLog 创建一个新的 OptLog 实例，并应用默认值
func NewOptLog(module, trace string) *OptLog {
	return &OptLog{
		Module:    module,
		Trace:     trace,
		OpTime:    time.Now().Format(format.CommonLayout),
		CreatedAt: time.Now().Format(format.CommonLayout),
		// 设置其他需要的默认值
	}
}

// WithOpTime 为 OptLog 设置操作时间
func (opt *OptLog) WithOpTime(time string) *OptLog {
	opt.OpTime = time
	return opt
}

func (opt *OptLog) WithOpType(ty string) *OptLog {
	opt.OpType = ty
	return opt
}

func (opt *OptLog) WithMethod(method string) *OptLog {
	opt.Method = method
	return opt
}
func (opt *OptLog) WithIp(ip string) *OptLog {
	opt.IP = ip

	return opt
}

func (opt *OptLog) WithUrl(url string) *OptLog {
	opt.URL = url
	return opt

}

func (opt *OptLog) WithParam(param string) *OptLog {
	opt.Param = param
	return opt
}

func (opt *OptLog) WithRsp(ret Result) *OptLog {
	opt.Result = ret
	return opt
}

func (opt *OptLog) WithUser(user string) *OptLog {
	opt.User = user
	return opt
}

func (opt *OptLog) WithCompany(company string) *OptLog {

	opt.Company = company
	return opt
}

func (opt *OptLog) WithPermission(permission string) *OptLog {
	opt.Permission = permission
	return opt
}
