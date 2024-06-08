package main

import "avei-todos-fiber/internal/app"

// @title						Avei Todos API with Fiber and Clean Architecture
// @version					1.0
// @description				Uh, i create this because i want to autogenerated my todos to pdf.
// @host						localhost:7009
// @BasePath					/
func main() {
	server := app.InitHttp()
	server.InitRouteAndServe()
}
