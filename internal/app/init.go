package app

import (
	"avei-todos-fiber/internal/handler"
	projectHandler "avei-todos-fiber/internal/handler/project"
	todoHandler "avei-todos-fiber/internal/handler/todo"
	userHandler "avei-todos-fiber/internal/handler/user"
	
	projectRepository "avei-todos-fiber/internal/repository/database/project"
	todoRepository "avei-todos-fiber/internal/repository/database/todo"
	userRepository "avei-todos-fiber/internal/repository/database/user"
	
	projectService "avei-todos-fiber/internal/service/project"
	todoService "avei-todos-fiber/internal/service/todo"
	userService "avei-todos-fiber/internal/service/user"
	
	"avei-todos-fiber/internal/repository/neon"
	"avei-todos-fiber/pkg/dotenv"
	"context"
	"os"
)

func InitHttp() *Server {
	var (
		db               neon.DB
		ctx              context.Context
		connectionString string
	)

	ctx = context.Background()
	connectionString = os.Getenv("DATABASE_URL")

	if connectionString == "" {
		connectionString = dotenv.GetEnv("DATABASE_URL")
	}

	// DB Connection
	db = neon.NewDB(ctx, connectionString)

	// Repository
	todoRepository := todoRepository.NewRepository(db)
	projectRepository := projectRepository.NewRepository(db)
	userRepository := userRepository.NewRepository(db)

	// Service
	todoService := todoService.NewService(todoRepository)
	projectService := projectService.NewService(projectRepository)
	userService := userService.NewService(userRepository)

	// Handler
	todoHandler := todoHandler.NewHandler(todoService)
	projectHandler := projectHandler.NewHandler(projectService)
	userHanler := userHandler.NewHandler(userService)

	serverHandler := handler.NewHandlers(todoHandler, projectHandler, userHanler)
	server := NewServer(serverHandler)
	return server
}
