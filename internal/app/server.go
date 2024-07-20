package app

import (
	"avei-todos-fiber/internal/handler"
	"avei-todos-fiber/internal/router"
	"avei-todos-fiber/pkg/dotenv"
	"avei-todos-fiber/pkg/migrations"
	"fmt"
	"os"

	_ "avei-todos-fiber/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
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

	app.Use(cors.New())

	// Run Migrations
	migrations.Migrate()

	app.Mount("/", router.InitRouter(&s.handlers))

	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault)

	port := os.Getenv("PORT")

	if port == "" {
		port = dotenv.GetEnv("PORT")
	}

	fmt.Printf("Server is running on port %s\n", port)
	app.Listen(":" + port)
}
