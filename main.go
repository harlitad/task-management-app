package main

import (
	"log"

	"example.com/task-management-app/config"
	_ "example.com/task-management-app/docs"
	"example.com/task-management-app/handler"
	"example.com/task-management-app/repository"
	"example.com/task-management-app/router"
	"example.com/task-management-app/service"
	"example.com/task-management-app/usecase"
	"github.com/gin-gonic/gin"
)

// @title Task Management App APIs
// @version 1.0
// @description Task Management App Swagger APIs.
// @host localhost:8080
// @BasePath /task-management-app
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

	// Initiate Task dependencies
	// repository
	taskRepository := repository.NewTaskRepository(db)
	// service
	taskService := service.NewTaskService(taskRepository)
	// usecase
	taskUsecase := usecase.NewTaskUsecase(taskService)
	// handler
	taskHandler := handler.NewTaskHandler(taskUsecase)

	// Initiate User dependencies
	// repository
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)
	userHandler := handler.NewUserHandler(userUsecase)

	// setup router
	router := router.NewRouter(router.Routes{Config: *appConfig, TaskHandler: taskHandler, UserHandler: userHandler})
	// wip: swagger
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
