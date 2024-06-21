package api

import (
	"net/http"
	"reflect"
)

type ApiResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func ApiSuccess(data any) *ApiResponse {
	if isAllFieldsPrivate(data) {
		return &ApiResponse{}
	}
	return &ApiResponse{
		Code: 200,
		Data: data,
		Msg:  "success",
	}
}

func ApiError(err error) *ApiResponse {
	if err == nil {
		return ApiSuccess(nil)
	}
	return &ApiResponse{Code: http.StatusNotFound, Msg: err.Error()}
}

func isAllFieldsPrivate(v any) bool {
	typeOf := reflect.TypeOf(v)
	if typeOf == nil {
		return false
	}
	if typeOf.Kind() == reflect.Ptr {
		typeOf = typeOf.Elem()
	}
	if typeOf.Kind() != reflect.Struct {
		return false
	}
	num := typeOf.NumField()
	for i := 0; i < num; i++ {
		c := typeOf.Field(i).Name[0]
		if c >= 'A' && c <= 'Z' {
			return false
		}
	}
	return true
}
