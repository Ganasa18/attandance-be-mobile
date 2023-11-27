package main

import (
	"fmt"
	"ganasa18/attandance-be-mobile/app"
	"ganasa18/attandance-be-mobile/helper"
	"net/http"
	"os"
)

func main() {
	// App environtment ...
	helper.Env()
	host := os.Getenv("APP_URL")
	// CHECK HOST NOT EMPTY
	if host == "" {
		host = "localhost"
	}
	// CHECK PORT NOT EMPTY
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	// DB CONNECTION INITIAL
	// db := app.NewDB()
	app.NewDB()

	//SERVER
	server := http.Server{
		Addr: host + ":" + port,
		// Handler: middleware.NewAuthMiddleware(router),
	}

	// LISTEN SERVER
	fmt.Printf("Server listening on port %s\n", server.Addr)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
