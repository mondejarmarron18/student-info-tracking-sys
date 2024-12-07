package v1

import (
	"log"
	"net/http"
	"server/v1/middlewares"
	"server/v1/utils"
)

func Server() *http.Server {
	config := utils.GetConfig()

	//Establish database connection
	utils.NewDB(config.DbUrl)

	mux := http.NewServeMux()

	//Register routes
	mux.Handle("/", NewRoute().Register())

	//Apply middlewares
	routeWithMiddlewares := middlewares.Apply(mux, []middlewares.Middleware{
		middlewares.Logger,
		middlewares.Header,
	})

	//Server configuration
	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: routeWithMiddlewares,
	}

	log.Println("Listening on port", config.Port)

	return server
}
