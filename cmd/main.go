package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/harlitad/task-management-app/config"
	_ "github.com/harlitad/task-management-app/docs"
	"github.com/harlitad/task-management-app/internal/handler"
	"github.com/harlitad/task-management-app/internal/repository"
	"github.com/harlitad/task-management-app/internal/router"
	"github.com/harlitad/task-management-app/internal/service"
	"github.com/harlitad/task-management-app/internal/usecase"
	"github.com/harlitad/task-management-app/pkg/logger"
)

// @title Task Management Service APIs
// @version 1.0
// @description Task Management App Swagger APIs.
// @host localhost:8080
// @BasePath /task-management-service
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// parsing envar
	appConfig := config.ParseConfig()

	app := initiateApp(appConfig)
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("http listen failed!")
	}
}

func initiateApp(appConfig *config.Config) *gin.Engine {

	// create new client of postgreSql
	db, err := repository.NewPostgreClient(*appConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewLogrusLogger(*appConfig)

	// Task
	taskRepository := repository.NewTaskRepository(logger, db)
	taskService := service.NewTaskService(logger, taskRepository)
	taskUsecase := usecase.NewTaskUsecase(logger, taskService)
	taskHandler := handler.NewTaskHandler(logger, taskUsecase)

	// User
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)
	userHandler := handler.NewUserHandler(userUsecase)

	// setup router
	router := router.NewRouter(router.Router{Config: *appConfig, TaskHandler: taskHandler, UserHandler: userHandler})
	// wip: swagger
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
