package app

import (
	"avei-todos-fiber/internal/handler"
	projectHandler "avei-todos-fiber/internal/handler/project"
	todoHandler "avei-todos-fiber/internal/handler/todo"
	projectRepository "avei-todos-fiber/internal/repository/database/project"
	todoRepository "avei-todos-fiber/internal/repository/database/todo"
	"avei-todos-fiber/internal/repository/neon"
	projectService "avei-todos-fiber/internal/service/project"
	todoService "avei-todos-fiber/internal/service/todo"
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

	// Service
	todoService := todoService.NewService(todoRepository)
	projectService := projectService.NewService(projectRepository)

	// Handler
	todoHandler := todoHandler.NewHandler(todoService)
	projectHandler := projectHandler.NewHandler(projectService)

	serverHandler := handler.NewHandlers(todoHandler, projectHandler)
	server := NewServer(serverHandler)
	return server
}
