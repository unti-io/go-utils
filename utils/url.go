package utils

import (
	"fmt"
	"reflect"
	"strings"
)

var Url = &UrlStruct{}

type UrlStruct struct {
}

// Encoded - 将 map 编码为 URL 查询字符串 - x-www-form-urlencoded
func (this *UrlStruct) Encoded(params map[string]any) string {

	var parts []string

	for key, value := range params {
		parts = append(parts, this.build(key, value)...)
	}

	return strings.Join(parts, "&")
}

// build - 构建 URL 查询字符串
func (this *UrlStruct) build(key string, value any) []string {

	var parts []string

	switch item := value.(type) {
	case string:
		parts = append(parts, fmt.Sprintf("%s=%v", key, item))
	case []string:
		for _, sv := range item {
			parts = append(parts, fmt.Sprintf("%s[]=%v", key, sv))
		}
	case []int:
		for _, iv := range item {
			parts = append(parts, fmt.Sprintf("%s[]=%d", key, iv))
		}
	case map[string]any:
		for k, sub := range item {
			parts = append(parts, this.build(fmt.Sprintf("%s[%s]", key, k), sub)...)
		}
	case []any:
		for i, sub := range item {
			parts = append(parts, this.build(fmt.Sprintf("%s[%d]", key, i), sub)...)
		}
	default:
		parts = append(parts, fmt.Sprintf("%s=%v", key, this.stringify(value)))
	}

	return parts
}

// stringify - 将任意类型转换为字符串
func (this *UrlStruct) stringify(value any) string {

	item := reflect.ValueOf(value)

	switch item.Kind() {
	case reflect.Array, reflect.Slice:
		var strVals []string
		for i := 0; i < item.Len(); i++ {
			strVals = append(strVals, fmt.Sprintf("%v", item.Index(i)))
		}
		return strings.Join(strVals, ",")
	default:
		return fmt.Sprintf("%v", value)
	}
}
