package format

import "encoding/json"

//common constant
const (
	EmptyString     = ""
	CommonLayout    = "2006-01-02 15:04:05"
	EmptyLen        = 0
	Zero            = 0
	MaxSize         = 2000
	DefaultPageNo   = 1
	DefaultPageSize = 20
)

// ToJsonString
func ToJsonString(v interface{}) string {
	if ret, err := Marshal(v); err != nil {
		return err.Error()
	} else {
		return string(ret)
	}
}

// FromJsonString
func FromJsonString(j string, v interface{}) error {
	return Unmarshal([]byte(j), v)
}

// Unmarshal
// 序列化，包装原生json
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// Marshal
// 序列化，包装原生json
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
