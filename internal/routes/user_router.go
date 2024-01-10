package routes

import (
	"fmt"
	controller "ganasa18/attandance-be-mobile/internal/controller/user"
	"ganasa18/attandance-be-mobile/internal/middleware"
	"net/http"

	_ "ganasa18/attandance-be-mobile/docs"

	"github.com/julienschmidt/httprouter"

	httpSwagger "github.com/swaggo/http-swagger"
)

func swaggerHandler(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3002/swagger/doc.json"),
	).ServeHTTP(writer, r)
}

func ApiUserRoute(router *httprouter.Router, userController controller.UserController) {
	router.GET("/api/v1/user", middleware.CustomAuthMiddleware(userController.FindAll))
	router.GET("/api/v1/user/:userId", middleware.CustomAuthMiddleware(userController.FindOne))
}

func ApiRouteGlobal(router *httprouter.Router) {
	router.GET("/api/test", func(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "HELLO TEST")
	})

	router.GET("/swagger/", swaggerHandler)
	router.GET("/swagger/:name", swaggerHandler)

}
