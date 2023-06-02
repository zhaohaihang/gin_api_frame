package v1

import (
	"encoding/json"
	"fmt"
	"gin_api_frame/internal/serializer"

	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			return serializer.Response{
				Status: http.StatusBadRequest,
				Msg:    fmt.Sprintf("%s%s", e.Field(), e.Tag()),
				Error:  fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Status: http.StatusBadRequest,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
