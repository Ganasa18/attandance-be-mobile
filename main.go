package main

import (
	"fmt"
	"ganasa18/attandance-be-mobile/app"
	"ganasa18/attandance-be-mobile/exception"
	"ganasa18/attandance-be-mobile/helper"
	"ganasa18/attandance-be-mobile/routes"
	"net/http"
	"os"

	routeshandler "ganasa18/attandance-be-mobile/routes/handler"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
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
	db := app.NewDB()
	validate := validator.New()

	//  SETUP ROUTES CONTROLLER
	userRouteHandler := routeshandler.SetupUserHandlerRoute(db, validate)

	// ROUTES PANIC
	router := httprouter.New()
	routes.ApiUserRoute(router, userRouteHandler)
	routes.ApiRouteTest(router)

	router.PanicHandler = exception.ErrorHandler

	//SERVER
	server := http.Server{
		Addr:    host + ":" + port,
		Handler: router,
		// Handler: middleware.NewAuthMiddleware(router),
	}

	// LISTEN SERVER
	fmt.Printf("Server listening on port %s\n", server.Addr)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
