package util

import (
	"net/http"
	"reflect"
	"runtime"
)

//GetFunctionName return function name from pointer
func GetFunctionName(function http.HandlerFunc) string {
	return runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
}
