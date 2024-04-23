package web

import "github.com/gin-gonic/gin"

type response struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, code int, data interface{}) {
	c.JSON(code, response{
		Code:    code,
		Status:  true,
		Message: "Success",
		Data:    data,
	})
}

func ResponseError(c *gin.Context, code int, message string) {
	c.JSON(code, response{
		Code:    code,
		Status:  false,
		Message: message,
		Data:    nil,
	})
}