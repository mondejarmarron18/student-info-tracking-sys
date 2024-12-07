package v1

import (
	"net/http"
	"server/v1/features/users/routes"
)

type Route struct {
	userRoute *routes.UserRoute
}

func NewRoute() *Route {
	return &Route{
		userRoute: routes.NewUserRoute(),
	}
}

func (r *Route) Register() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/users/", http.StripPrefix("/users", r.userRoute.Register()))

	return mux
}
