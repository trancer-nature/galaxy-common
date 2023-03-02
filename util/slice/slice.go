package slice

import (
	"fmt"
	"reflect"
	"strings"
)

// 任意数组生成[]interface{}
func ToInterfaces(arr interface{}) []interface{} {
	av := reflect.Indirect(reflect.ValueOf(arr))
	// 传入为空
	if arr == nil || av.Kind() == reflect.Ptr && av.IsNil() {
		return nil
	}

	if reflect.Slice != av.Kind() && reflect.Array != av.Kind() {
		return []interface{}{arr}
	}

	ret := make([]interface{}, 0)
	for i := 0; i < av.Len(); i++ {
		item := av.Index(i)
		if item.IsValid() {
			ret = append(ret, item.Interface())
		}
	}

	return ret
}

// 任意数组转换成String数组
func ToStrings(arr interface{}) []string {
	av := reflect.Indirect(reflect.ValueOf(arr))
	// 传入为空
	if arr == nil || av.Kind() == reflect.Ptr && av.IsNil() {
		return nil
	}

	if reflect.Slice != av.Kind() {
		return []string{fmt.Sprintf("%+v", av.Interface())}
	}

	ret := make([]string, 0)
	for i := 0; i < av.Len(); i++ {
		item := av.Index(i)
		if !item.IsValid() {
			continue
		}

		if reflect.String == item.Kind() {
			ret = append(ret, item.String())
		} else {
			ret = append(ret, fmt.Sprintf("%+v", item.Interface()))
		}
	}

	return ret
}

// 数组连接
func Join(objs interface{}, splitter string) string {
	return strings.Join(ToStrings(objs), splitter)
}
