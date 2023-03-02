package number

import (
	"fmt"
	"strconv"
)

// 取对应类型的指针
func Uint32(in uint32) *uint32 {
	return &in
}

// 取对应类型的指针
func Uint64(in uint64) *uint64 {
	return &in
}

// 取对应类型的指针
func Int(in int) *int {
	return &in
}

// 取对应类型的指针
func Int32(in int32) *int32 {
	return &in
}

// 取对应类型的指针
func Int64(in int64) *int64 {
	return &in
}

// 取对应类型的指针
func Float64(in float64) *float64 {
	return &in
}

// 取对应类型的指针
func String(in string) *string {
	return &in
}

// 解析string类型，输出相应数字类型指针
func ParseUint32(in string) uint32 {
	ret, _ := strconv.ParseUint(in, 10, 32)
	return uint32(ret)
}

// 解析string类型，输出相应数字类型指针
func ParseUint64(in string) uint64 {
	ret, _ := strconv.ParseUint(in, 10, 64)
	return ret
}

// 解析string类型，输出相应数字类型指针
func ParseInt32(in string) int32 {
	ret, _ := strconv.ParseInt(in, 10, 32)
	return int32(ret)
}

// 解析string类型，输出相应数字类型指针
func ParseInt64(in string) int64 {
	ret, _ := strconv.ParseInt(in, 10, 64)
	return ret
}

// 解析string类型，输出相应数字类型指针
func ParseFloat64(in string) float64 {
	ret, _ := strconv.ParseFloat(in, 64)
	return ret
}

// 输出相应相应的string
func ToString(in interface{}) string {
	return fmt.Sprintf("%+v", in)
}
