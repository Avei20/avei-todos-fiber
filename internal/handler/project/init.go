package project

import "avei-todos-fiber/internal/service/project"

func NewHandler(projectService project.Service) Handler {
	return &HandlerImpl{
		projectService: projectService,
	}
}
