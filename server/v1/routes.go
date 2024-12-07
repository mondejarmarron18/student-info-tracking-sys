package v1

import (
	"net/http"
	"server/v1/features/users/routes"
)

type Routes struct {
	userRoute *routes.UserRoute
}

func NewRoutes() *Routes {
	return &Routes{
		userRoute: routes.NewUserRoute(),
	}
}

func (r *Routes) Register() *http.ServeMux {
	mux := http.NewServeMux()

	//Register routes
	r.userRoute.Register(mux)

	newMux := http.NewServeMux()
	newMux.Handle("/v1/", http.StripPrefix("/v1", mux))

	return newMux
}
