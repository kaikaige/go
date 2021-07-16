package form

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Middleware(errFunc func(ctx *gin.Context, err error), sucFunc func(ctx *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if formData, found := c.Get("form"); found {
			form := formData.(Form)
			form.Run(c)
			if form.GetError() != nil {
				if errFunc == nil {
					code, msg := GetError(form.GetError(), form)
					c.JSON(code, msg)
				} else {
					errFunc(c, form.GetError())
				}

			} else {
				if sucFunc == nil {
					c.JSON(http.StatusOK, form.GetData())
				} else {
					sucFunc(c)
				}
			}
		}
	}
}
