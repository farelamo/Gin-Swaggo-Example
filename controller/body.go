package controller

import (
	"exp-wagger/model/web"
	"log"

	"github.com/gin-gonic/gin"
)

//	@Tags			Body
//	@Summary		Experience with body
//	@Description	Try with body here
//	@Accept			json
//	@Produce		json
//  @Param			category	body		web.RequestBody	true	"Request Body"
//	@Success		200	{object}	web.response
//	@Failure		400	{object}	web.response
//	@Router			/body/ [post]
func WithBody(c *gin.Context) {
	req  := new(web.RequestBody)
	if err := c.ShouldBind(req); err != nil {
		log.Println("masukk")
		web.ResponseError(c, 400, "Terjadi Kesalahan Saat Mengurai Data")
		return
	}

	web.Response(c, 200, req)
}