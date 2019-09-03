package utils

import (
	"encoding/json"
	"fmt"
	"go-fast-ws-mq/serializer"
	"net/http"

	"gopkg.in/go-playground/validator.v8"
)

func BuildErrorResponse(err error) *serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := fmt.Sprintf("Field.%s", e.Field)
			tag := fmt.Sprintf("Tag.Valid.%s", e.Tag)
			return &serializer.Response{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("%s%s", field, tag),
				Error:   fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return &serializer.Response{
			Code:    http.StatusBadRequest,
			Message: "JSON类型不匹配",
			Error:   fmt.Sprint(err),
		}
	}

	return &serializer.Response{
		Code:    http.StatusBadRequest,
		Message: "参数错误",
		Error:   fmt.Sprint(err),
	}
}
