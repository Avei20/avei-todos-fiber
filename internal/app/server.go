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
	"github.com/gofiber/swagger"
)

func NewServer(
	handlers handler.Handlers,
) *Server {
	return &Server{
		handlers: handlers,
	}
}

//	@title			Fiber Example API
//	@version		1.0
//	@description	This is a sample swagger for Fiber
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.email	fiber@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

func (s *Server) InitRouteAndServe() {
	app := fiber.New()

	// Run Migrations
	migrations.Migrate()

	app.Mount("/", router.InitRouter(&s.handlers))

	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault) // default

	app.Get("/docs/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	port := os.Getenv("PORT")

	if port == "" {
		port = dotenv.GetEnv("PORT")
	}

	fmt.Printf("Server is running on port %s\n", port)
	app.Listen(":" + port)
}
