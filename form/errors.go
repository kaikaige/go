package form

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"strconv"
)

type HttpError struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func (e HttpError) Error() string { return e.Message }

func BadRequest (ms ...string) HttpError {
	message := "Bad Request"
	if len(ms) > 0 {
		message = ms[0]
	}
	return HttpError{
		Code: http.StatusBadRequest,
		Message: message,
	}
}

func NotFound (ms ...string) HttpError {
	message := "Not Found"
	if len(ms) > 0 {
		message = ms[0]
	}
	return HttpError{
		Code: http.StatusNotFound,
		Message: message,
	}
}

func ValidateError (ms ...string) HttpError {
	message := "Form Validate Error"
	if len(ms) > 0 {
		message = ms[0]
	}
	return HttpError{
		Code: http.StatusUnprocessableEntity,
		Message: message,
	}
}

func GetError(err error, ref interface{}) (code int, errMsg string) {
	code = 400
	switch e := err.(type) {
	case validator.ValidationErrors:
		if ref, found := reflect.ValueOf(ref).Type().FieldByName(e[0].StructField()); found {
			code = 422
			if errMsg = ref.Tag.Get(e[0].ActualTag() + "-msg"); errMsg == "" {
				var label string
				if label = ref.Tag.Get("label"); label == "" {
					label = e[0].StructField()
				}
				switch e[0].ActualTag() {
				case "required":
					errMsg = label + "不能为空"
				default:
					errMsg = label
				}
			}
		}
	case *strconv.NumError:
		code = 422
		errMsg = e.Num + "不是一个数字"
	case HttpError:
		code = e.Code
		errMsg = e.Message
	default:
		errMsg = fmt.Sprintf("未知异常：%T, msg: %s\n", err, err.Error())
	}
	return
}
