package format

import (
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

//common constant
const (
	EmptyString     = ""
	CommonLayout    = "2006-01-02 15:04:05"
	EmptyLen        = 0
	Zero            = 0
	MaxSize         = 2000
	DefaultPageNo   = 1
	DefaultPageSize = 20
	DefaultRetCode  = 0
	Success         = "success"
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

func GetTraceSpanID(ctx context.Context) string {
	var traceID string
	var spanID string
	spanCtx := trace.SpanContextFromContext(ctx)
	spanID = spanCtx.SpanID().String()
	if spanCtx.HasTraceID() {
		traceID = spanCtx.TraceID().String()
	}
	return fmt.Sprintf("%s-%s", traceID, spanID)
}

// Ptr returns a pointer to the provided value.
func Ptr[T any](v T) *T {
	return &v
}

func ToMap(tag string, in interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			// set key of map to value in struct field
			val := v.Field(i)
			zero := reflect.Zero(val.Type()).Interface()
			current := val.Interface()
			if reflect.DeepEqual(current, zero) {
				continue
			}
			out[tagv] = current
		}
	}
	return out
}

func ToList(tag string, in interface{}) [2]string {
	var out [2]string
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			// set key of map to value in struct field
			val := v.Field(i)
			zero := reflect.Zero(val.Type()).Interface()
			current := val.Interface()

			if reflect.DeepEqual(current, zero) {
				continue
			}

			if cur, ok := current.(string); ok {
				return [2]string{tagv, cur}
			}

			return out
		}
	}

	return out
}
