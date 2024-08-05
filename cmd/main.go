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
	files "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
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
	// Parsing environment variables
	appConfig := config.ParseConfig()

	// Initialize the application
	app := initiateApp(appConfig)

	// Set up Swagger
	app.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	// Start the application
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("http listen failed: %v", err)
	}
}

func initiateApp(appConfig *config.Config) *gin.Engine {

	logger := logger.NewLogrusLogger(*appConfig)

	appConfig.UseMongo = true

	var taskRepository repository.ITaskRepository
	var userRepository repository.IUserRepository

	if appConfig.UseMongo {
		logger.Info("Database uses MongoDB")

		// Initialize MongoDB client
		mongoClient, err := repository.NewMongoClient(*appConfig)
		if err != nil {
			logger.Fatalf("Failed to connect mongo db, %s", err.Error())
		}

		// Initialize MongoDB repositories
		taskRepository = repository.NewMongoTaskRepository(logger, mongoClient)
		userRepository = repository.NewUserMongoRepository(mongoClient)
	} else {
		// Initialize PostgreSQL client
		postgresDB, err := repository.NewPostgreClient(*appConfig)
		if err != nil {
			logger.Fatalf("Failed to connect postgres db, %s", err.Error())
		}

		// Initialize PostgreSQL repositories
		taskRepository = repository.NewPostgreTaskRepository(logger, postgresDB)
		userRepository = repository.NewUserPostgresRepository(postgresDB)
	}

	// Task
	taskService := service.NewTaskService(logger, taskRepository)
	taskUsecase := usecase.NewTaskUsecase(logger, taskService)
	taskHandler := handler.NewTaskHandler(logger, taskUsecase)

	// User
	userService := service.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)
	userHandler := handler.NewUserHandler(userUsecase)

	// Setup router
	r := router.NewRouter(router.Router{Config: *appConfig, TaskHandler: taskHandler, UserHandler: userHandler})

	return r
}
