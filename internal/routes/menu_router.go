package routes

import (
	controller "ganasa18/attandance-be-mobile/internal/controller/crud"

	"github.com/julienschmidt/httprouter"
)



func ApiMenuRoute(router *httprouter.Router, menuController controller.MenuController) {
	router.POST("/api/v1/menu", menuController.CreateMenu)
	router.GET("/api/v1/menu", menuController.FindAll)
}
