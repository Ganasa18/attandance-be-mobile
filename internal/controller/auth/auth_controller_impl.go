package controller

import (
	"ganasa18/attandance-be-mobile/internal/helper"
	"ganasa18/attandance-be-mobile/internal/model/web"
	service "ganasa18/attandance-be-mobile/internal/service/auth"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	registerRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &registerRequest)

	registerResponse := controller.AuthService.Register(request.Context(), registerRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   registerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.AuthLoginRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	loginResponse, err := controller.AuthService.Login(request.Context(), loginRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   loginResponse,
	}

	if err != nil {
		helper.CustomResponseError(writer, request, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *AuthControllerImpl) LoginWithOutOtp(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.AuthLoginRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	loginResponse, err := controller.AuthService.Login(request.Context(), loginRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   loginResponse,
	}

	if err != nil {
		helper.CustomResponseError(writer, request, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteToResponseBody(writer, webResponse)

}
