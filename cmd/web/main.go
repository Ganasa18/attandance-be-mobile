package main

import (
	"fmt"
	"ganasa18/attandance-be-mobile/internal/app"
	"ganasa18/attandance-be-mobile/internal/exception"
	"ganasa18/attandance-be-mobile/internal/helper"
	"ganasa18/attandance-be-mobile/internal/middleware"
	"ganasa18/attandance-be-mobile/internal/routes"
	"net/http"
	"os"

	routeshandler "ganasa18/attandance-be-mobile/internal/routes/handler"

	_ "ganasa18/attandance-be-mobile/docs"

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
	authRouteHandler := routeshandler.SetupAuthHandlerRoute(db, validate)
	roleRouteHanlder := routeshandler.SetupRoleHandlerRoute(db, validate)
	menuRouteHanlder := routeshandler.SetupMenuHandlerRoute(db, validate)
	userMenuAccessRouteHandler := routeshandler.SetupUserMenuAccessHandlerRoute(db, validate)
	roleAccessRouteHandler := routeshandler.SetupRoleAccessHandlerRoute(db, validate)

	// ROUTES PANIC
	router := httprouter.New()

	routes.ApiUserRoute(router, userRouteHandler)
	routes.ApiAuthRoute(router, authRouteHandler)
	routes.ApiRoleRoute(router, roleRouteHanlder, roleAccessRouteHandler)
	routes.ApiMenuRoute(router, menuRouteHanlder)
	routes.ApiUserMenuAccessRoute(router, userMenuAccessRouteHandler)
	routes.ApiRouteGlobal(router)

	router.PanicHandler = exception.ErrorHandler

	//SERVER
	server := http.Server{
		Addr:    host + ":" + port,
		Handler: middleware.NewBaseMiddleware(router),
	}

	// LISTEN SERVER
	fmt.Printf("Server listening on port %s\n", server.Addr)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
