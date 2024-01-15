package router

import (
	"example.com/task-management-app/handler"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	TaskHandler handler.TaskHandler
	RouteBase   *gin.RouterGroup
}

func SetupRouter(routes Routes) *gin.Engine {
	r := gin.Default()
	routes.RouteBase = r.Group("api")
	routes.taskRoute()
	return r
}

func (r *Routes) taskRoute() {
	r.RouteBase.Group("/v1/task").
		GET("/", r.TaskHandler.List).
		GET("/:id", r.TaskHandler.Get)

	// api.GET("/task/:id", handler.Get)
	// api.POST("/task", handler.Create)
	// api.DELETE("/task/:id", handler.Delete)
}

// func userRoute(r *gin.Engine, handler handler.Handler) {
// 	r.GET("/api/user", handler.List)
// 	r.GET("/api/user/:id", handler.Get)
// 	r.POST("/api/user", handler.Create)
// 	r.DELETE("/api/user/:id", handler.Delete)
// }
