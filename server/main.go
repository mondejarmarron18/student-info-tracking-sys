package main

import v1 "server/v1"

func main() {
	server := v1.Server()

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
