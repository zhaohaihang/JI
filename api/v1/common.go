package v1

import (
	"encoding/json"
	"fmt"
	"ji/consts"
	"ji/serializer"

	"github.com/go-playground/validator/v10"
)

func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			return serializer.Response{
				Status: consts.IlleageRequest,
				Msg:    fmt.Sprintf("%s%s", e.Field(), e.Tag()),
				Error:  fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: consts.IlleageRequest,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Status: consts.IlleageRequest,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
