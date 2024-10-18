package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Response struct {
	Code int `json:"code"`
	Message interface{} `json:"message"`
	Data interface{} `json:"data"`
}

func RenderSuccess(c *gin.Context, data interface{}, code int, message interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Message: message,
		Data: data,
	})
}

func RenderError(c *gin.Context, code int, message string) {
	c.JSON(code, &Response{
		Code: code,
		Message: message,
	})
}