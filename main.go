package main

import "avei-todos-fiber/internal/app"

func main() {
	server := app.InitHttp()

	server.InitRouteAndServe()
}
