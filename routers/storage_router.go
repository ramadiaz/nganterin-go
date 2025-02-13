package routers

import (
	"nganterin-go/api/storages/controllers"

	"github.com/gin-gonic/gin"
)

func StorageRoutes(r *gin.RouterGroup, orderControllers controllers.CompControllers) {
	filesGroup := r.Group("/files")
	{
		filesGroup.POST("/upload", orderControllers.Create)
	}
}
