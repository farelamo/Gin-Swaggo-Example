package main

import (
	"exp-wagger/controller"
	"exp-wagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@contact.name	Fariz Fahmi Faturachmad
//	@contact.url	https://github.com/farelamo
//	@contact.email	farizfahmi674@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	r := gin.Default()

	docs.SwaggerInfo.Title 			= "API Documentation"
	docs.SwaggerInfo.Description 	= "This is a sample description of documentation"
	docs.SwaggerInfo.Version 		= "1.0"
	docs.SwaggerInfo.Host 	  		= "localhost:3000"
	docs.SwaggerInfo.BasePath 		= "/api/v1"
	docs.SwaggerInfo.Schemes  		= []string{"http", "https"}

	api := r.Group("/api/v1")
	{
		api.GET("/ping", controller.Ping)
		api.GET("/params/:id", controller.WithParams)
		api.POST("/body", controller.WithBody)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":3000")
}
