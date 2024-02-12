package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewSwaggerApi(router *gin.RouterGroup) {
	router.GET("/apidocs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// router.Static("/apidocs", "./contents/swagger-ui/")
}
