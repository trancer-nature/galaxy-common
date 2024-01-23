package xlog

import (
	"bytes"
	"github.com/trancer-nature/galaxy-common/util/format"
	"io/ioutil"
	"net/http"
	"time"
)

type LogWriter struct {
	http.ResponseWriter
	Body *bytes.Buffer
}

// OptLog 定义了操作日志的结构
type OptLog struct {
	User         string `json:"user,omitempty"`
	OpTime       string `json:"op_time,omitempty"`
	OpType       string `json:"op_type,omitempty"`
	IP           string `json:"ip,omitempty"`
	Company      string `json:"company,omitempty"`
	Permission   string `json:"permission,omitempty"`
	Method       string `json:"method,omitempty"`
	URL          string `json:"url,omitempty"`
	Module       string `json:"module,omitempty"`
	Param        string `json:"param,omitempty"`
	Trace        string `json:"trace,omitempty"`
	HttpCode     int32  `json:"http_code,omitempty"`
	ResponseData string `json:"response_data,omitempty"`
	Msg          string `json:"msg,omitempty"`
}

// NewOptLog 创建一个新的 OptLog 实例，并应用默认值
func NewOptLog(module, trace string) *OptLog {
	return &OptLog{
		Module: module,
		Trace:  trace,
		OpTime: time.Now().Format(format.CommonLayout),
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

// ExtractParam 从 http.Request 中提取参数
func ExtractParam(r *http.Request) string {
	if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
		if r.Body != nil {
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(r.Body)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(buf.Bytes()))
			return buf.String()
		}
	}
	if r.Method == http.MethodGet {
		return r.URL.Query().Encode()
	}
	return ""
}

type xLogResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func NewLoggingResponseWriter(w http.ResponseWriter) *xLogResponseWriter {
	return &xLogResponseWriter{ResponseWriter: w, statusCode: http.StatusOK, body: &bytes.Buffer{}}
}

func (rw *xLogResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *xLogResponseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *xLogResponseWriter) GetCode() int {
	return rw.statusCode
}

func (rw *xLogResponseWriter) GetBody() *bytes.Buffer {
	return rw.body
}
