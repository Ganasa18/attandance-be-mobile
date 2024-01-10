package routes

import (
	controller "ganasa18/attandance-be-mobile/internal/controller/crud"
	"ganasa18/attandance-be-mobile/internal/middleware"

	"github.com/julienschmidt/httprouter"
)

func ApiUserMenuAccessRoute(router *httprouter.Router, userMenuAccessController controller.UserMenuAccessController) {
	router.POST("/api/v1/user-menu-access", middleware.CustomAuthMiddleware(userMenuAccessController.CreateUserMenuAccess))
	router.GET("/api/v1/user-menu-access/user", userMenuAccessController.GetUserMenuAccess)
}
