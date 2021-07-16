package form

import (
	"github.com/gin-gonic/gin"
)

type BaseForm struct {
	Err     error
	Data    interface{}
	Message string
}

func (form *BaseForm) GetError() error {
	return form.Err
}

func (form *BaseForm) GetData() interface{} {
	return form.Data
}

func (form *BaseForm) GetMessage() string {
	return form.Message
}

func (form *BaseForm) Run(*gin.Context) {

}
