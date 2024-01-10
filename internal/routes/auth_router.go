package routes

import (
	controller "ganasa18/attandance-be-mobile/internal/controller/auth"

	"github.com/julienschmidt/httprouter"
)

func ApiAuthRoute(router *httprouter.Router, authController controller.AuthController) {
	router.POST("/api/v1/userRegister", authController.Register)
	router.POST("/api/v1/userLogin", authController.Login)
	router.POST("/api/v1/userLoginWeb", authController.LoginWithOutOtp)
}
