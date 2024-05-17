package app

import "avei-todos-fiber/internal/handler"

func InitHttp() *Server {
	serverHandler := handler.NewHandler()
	server := NewServer(serverHandler)
	return server
}
