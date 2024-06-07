package app

import (
	"avei-todos-fiber/internal/handler"
	"avei-todos-fiber/internal/router"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func NewServer(
	handlers handler.Handlers,
) *Server {
	return &Server{
		handlers: handlers,
	}
}

func (s *Server) InitRouteAndServe() {
	app := fiber.New()

	app.Mount("/", router.InitRouter(s.handlers))

	port := os.Getenv("PORT")

	if port == "" {
		port = "7009"
	}

	fmt.Printf("Server is running on port %s\n", port)
	app.Listen(":" + port)
}
