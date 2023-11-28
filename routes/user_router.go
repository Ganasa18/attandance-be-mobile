package routes

import (
	"fmt"
	"ganasa18/attandance-be-mobile/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ApiUserRoute(router *httprouter.Router, userController controller.UserController) {
	router.GET("/api/v1/user", userController.FindAll)
}

func ApiRouteTest(router *httprouter.Router) {
	router.GET("/api/test", func(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "HELLO TEST")
	})
}
