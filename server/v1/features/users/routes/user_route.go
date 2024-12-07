package routes

import (
	"net/http"
	"server/v1/features/users/controllers"
)

type UserRoute struct {
	userController *controllers.UserController
}

func NewUserRoute() *UserRoute {
	return &UserRoute{
		userController: controllers.NewUserController(),
	}
}

func (r *UserRoute) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /users", r.userController.GetUsers)
	mux.HandleFunc("GET /users/{id}", r.userController.GetUserById)
	mux.HandleFunc("POST /users", r.userController.CreateUser)
}
