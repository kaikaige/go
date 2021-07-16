package form

import "github.com/gin-gonic/gin"

type Form interface {
	GetError() error
	GetData() interface{}
	GetMessage() string
	Run(c *gin.Context)
}
