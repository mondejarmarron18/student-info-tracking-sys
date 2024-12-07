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

func (r *UserRoute) Register() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", r.userController.CreateUser)
	mux.HandleFunc("GET /", r.userController.GetUsers)
	mux.HandleFunc("GET /{id}", r.userController.GetUserById)
	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Del("Content-Type")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})

	return mux
}
