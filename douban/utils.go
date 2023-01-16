package douban

import (
	"strings"

	"github.com/tidwall/gjson"
)

func GJsonGetDefault(getValue gjson.Result, defaultValue interface{}) string {
	if !getValue.Exists() {
		switch value := defaultValue.(type) {
		case string:
			return value
		case gjson.Result:
			return value.String()
		}
	}
	return getValue.String()
}

func GJsonArrayToString(result gjson.Result, sep string) string {
	stringSlice := make([]string, 0)
	for _, item := range result.Array() {
		stringSlice = append(stringSlice, item.String())
	}
	return strings.Join(stringSlice, sep)
}

func GJsonArrayToStringExcept(result gjson.Result, exp, sep string) string {
	stringSlice := make([]string, 0)
	for _, item := range result.Array() {
		if item.String() != exp {
			stringSlice = append(stringSlice, item.String())
		}
	}
	return strings.Join(stringSlice, sep)
}
