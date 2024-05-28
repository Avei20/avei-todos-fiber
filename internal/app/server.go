package app

import (
	"avei-todos-fiber/internal/handler/todo"
	"avei-todos-fiber/internal/router"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func NewServer(
	todoHandler todo.Handler,
) *Server {
	return &Server{
		todoHandler: todoHandler,
	}
}

func (s *Server) InitRouteAndServe() {
	app := fiber.New()

	app.Mount("/", router.InitRouter(s.handler))

	port := os.Getenv("PORT")

	if port == "" {
		port = "7009"
	}

	fmt.Printf("Server is running on port %s\n", port)
	app.Listen(":" + port)
}
