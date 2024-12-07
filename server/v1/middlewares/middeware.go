package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

func Apply(route http.Handler, middlewares []Middleware) *http.ServeMux {
	mux := http.NewServeMux()

	handler := route

	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	mux.Handle("/", handler)

	return mux
}
