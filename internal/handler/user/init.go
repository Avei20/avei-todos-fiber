package user

import "avei-todos-fiber/internal/service/user"

func NewHandler(userService user.Service) Handler {
	return &HandlerImpl {
		userService: userService,
	}
}