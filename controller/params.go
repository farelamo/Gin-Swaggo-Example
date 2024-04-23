package controller

import (
	"exp-wagger/constants"
	"exp-wagger/model/web"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

//	@Tags			Params
//	@Summary		Experience with params
//	@Description	Try with params here
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		400	{object}	web.response
//	@Param			id	path		int	true	"ID"
//	@Router			/params/{id} [get]
func WithParams(c *gin.Context) {
	param  := new(web.RequestParam)
	if err := c.ShouldBindUri(param); err != nil {
		log.Println("masukk")
		web.ResponseError(c, 400, "Terjadi Kesalahan Saat Mengurai Data")
		return
	}

	log.Println("paramnya", param.ID)

	if !slices.Contains(constants.ParamIDs, strconv.Itoa(param.ID)) {
		web.ResponseError(c, 404, "Data dengan Id " + strconv.Itoa(param.ID) + " tidak ditemukan")
		return
	}

	result := "Ini adalah id anda " + strconv.Itoa(param.ID)
	web.Response(c, 200, result)
}