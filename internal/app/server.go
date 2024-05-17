package app

import (
	"avei-todos-fiber/internal/handler"
	"avei-todos-fiber/internal/router"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func NewServer(handler *handler.Handler) *Server {
	return &Server{
		handler: handler,
	}
}

func (s *Server) InitRouteAndServe() {
	r := fiber.New()

	r.Mount("/", router.InitRouter(s.handler))

	port := os.Getenv("PORT")

	if port == "" {
		port = "7009"
	}

	fmt.Printf("Server is running on port %s\n", port)
	r.Listen(":" + port)
}
