package controller

import (
	"exp-wagger/model/web"

	"github.com/gin-gonic/gin"
)

//	@Tags			Ping		
//	@Summary		Do a ping
//	@Description	get ping
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Router			/ping [get]
func Ping(c *gin.Context)  {
	web.Response(c, 200, "pong")
}