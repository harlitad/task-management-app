package router

import (
	"github.com/gin-gonic/gin"
	"github.com/harlitad/task-management-app/config"
	"github.com/harlitad/task-management-app/internal/handler"
	"github.com/harlitad/task-management-app/internal/middleware"
	"github.com/harlitad/task-management-app/swagger"
)

type Router struct {
	Config      config.Config
	TaskHandler handler.TaskHandler
	UserHandler handler.UserHandler
}

func NewRouter(r Router) *gin.Engine {
	app := gin.Default()
	baseRouter := app.Group(r.Config.BaseUrl)
	r.authRouter(baseRouter)
	r.taskRouter(baseRouter)
	swagger.NewSwaggerApi(baseRouter)
	return app
}

func (r *Router) taskRouter(base *gin.RouterGroup) {
	v1 := base.Group("/v1/task", middleware.AuthMiddleware())
	v1.GET("", r.TaskHandler.GetTasks)
	v1.POST("", r.TaskHandler.Create)
	v1.GET("/:id", r.TaskHandler.GetTaskById)
	v1.PATCH("/:id", r.TaskHandler.UpdateTask)
	v1.DELETE("/:id", r.TaskHandler.Delete)
	v1.GET("/today", r.TaskHandler.GetTasksToday)

	// // get list task due date today
	// v1.GET("/task/today", r.TaskHandler.ListTaskDueToday)

	// // Get a user task
	// v1.GET("/user/:userId/task/:id", r.TaskHandler.GetUserTask)
	// // list a user task
	// v1.GET("/user/:userId/task", r.TaskHandler.ListUserTask)
}

func (r *Router) authRouter(router *gin.RouterGroup) {
	router.POST("/register", r.UserHandler.Create)
	router.POST("/auth", r.UserHandler.Authentication)
}

// func userRoute(r *gin.Engine, handler handler.Handler) {
// 	r.GET("/api/user", handler.List)
// 	r.GET("/api/user/:id", handler.Get)
// 	r.POST("/api/user", handler.Create)
// 	r.DELETE("/api/user/:id", handler.Delete)
// }
