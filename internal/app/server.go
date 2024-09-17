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
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	app.Use(logger.New(
		logger.Config{
			Format:     "${time} ${status} - ${method} ${path} \nBody: ${body}\nParams: ${queryParams}\n",
			TimeFormat: "02-Jan-2024 13:04:05",
			TimeZone:   "Asia/Jakarta",
		},
	))

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
