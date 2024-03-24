package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	docs "github.com/NiltonMorais/gopportunities/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	// @Summary Health check
	// @Description Check if the service is up and running
	// @Produce json
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group(basePath)
	{
		v1.GET("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET Opening",
			})
		})
	}

}
