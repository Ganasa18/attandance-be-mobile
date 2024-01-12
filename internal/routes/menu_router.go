package routes

import (
	controller "ganasa18/attandance-be-mobile/internal/controller/crud"
	"ganasa18/attandance-be-mobile/internal/middleware"

	"github.com/julienschmidt/httprouter"
)

func ApiMenuRoute(router *httprouter.Router, menuController controller.MenuController) {
	router.POST("/api/v1/menu", middleware.CustomAuthMiddleware(menuController.CreateMenu))
	router.GET("/api/v1/menu", middleware.CustomAuthMiddleware(menuController.FindAll))
}
