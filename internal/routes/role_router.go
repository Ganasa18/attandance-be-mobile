package routes

import (
	controller "ganasa18/attandance-be-mobile/internal/controller/crud"
	"ganasa18/attandance-be-mobile/internal/middleware"

	"github.com/julienschmidt/httprouter"
)

func ApiRoleRoute(router *httprouter.Router, roleController controller.RoleController, roleAccessController controller.RoleAccessController) {
	router.POST("/api/v1/role", middleware.CustomAuthMiddleware(roleController.CreateRole))
	router.GET("/api/v1/role", middleware.CustomAuthMiddleware(roleController.FindAll))
	router.POST("/api/v1/role-access", middleware.CustomAuthMiddleware(roleAccessController.CreateRoleAccess))
	router.GET("/api/v1/role-access", middleware.CustomAuthMiddleware(roleAccessController.FindAll))
}
