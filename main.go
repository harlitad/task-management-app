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

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"
)

// @title Task Management API
func main() {
	initiateApp()
}

func initiateApp() {

	// parsing envar
	appConfig := config.ParseConfig()

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

	// setup router
	router := router.SetupRouter(router.Routes{TaskHandler: taskHandler})

	// wip: swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// run app
	router.Run()

}
