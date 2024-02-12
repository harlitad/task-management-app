package router

import (
	"example.com/task-management-app/config"
	"example.com/task-management-app/handler"
	"example.com/task-management-app/middleware"
	"example.com/task-management-app/swagger"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Config      config.Config
	TaskHandler handler.TaskHandler
	UserHandler handler.UserHandler
}

func NewRouter(routes Routes) *gin.Engine {
	r := gin.Default()

	baseRouter := r.Group(routes.Config.BaseUrl)
	routes.authRoute(baseRouter)

	// base router with auth
	baseRouterAuth := baseRouter.Group("/", middleware.AuthMiddleware())
	routes.taskRoute(baseRouterAuth)

	swagger.NewSwaggerApi(baseRouter)

	return r
}

func (r *Routes) taskRoute(router *gin.RouterGroup) {
	taskRouterV1 := router.Group("/v1/task")

	taskRouterV1.GET("/", r.TaskHandler.List)
	taskRouterV1.GET("/:id", r.TaskHandler.Get)
	taskRouterV1.POST("/", r.TaskHandler.Create)
	taskRouterV1.DELETE("/:id", r.TaskHandler.Delete)
}

func (r *Routes) userRoute(router *gin.RouterGroup) {
	// userRouterV1 := router.Group("/v1/user")

	// userRouterV1.POST("/", r.UserHandler.Create)
	// userRouterV1.POST("/auth", r.UserHandler.Authentication)
}

func (r *Routes) authRoute(router *gin.RouterGroup) {
	router.POST("/register", r.UserHandler.Create)
	router.POST("/auth", r.UserHandler.Authentication)
}

// func userRoute(r *gin.Engine, handler handler.Handler) {
// 	r.GET("/api/user", handler.List)
// 	r.GET("/api/user/:id", handler.Get)
// 	r.POST("/api/user", handler.Create)
// 	r.DELETE("/api/user/:id", handler.Delete)
// }
