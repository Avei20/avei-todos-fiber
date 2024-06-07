package app

import (
	"avei-todos-fiber/internal/handler"
	todoHandler "avei-todos-fiber/internal/handler/todo"
	todoRepository "avei-todos-fiber/internal/repository/database/todo"
	"avei-todos-fiber/internal/repository/neon"
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

	// Service
	todoService := todoService.NewService(todoRepository)

	// Handler
	todoHandler := todoHandler.NewHandler(todoService)

	serverHandler := handler.NewHandlers(todoHandler)
	server := NewServer(serverHandler)
	return server
}
