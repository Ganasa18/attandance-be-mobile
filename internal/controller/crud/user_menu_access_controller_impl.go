package controller

import (
	"ganasa18/attandance-be-mobile/internal/helper"
	"ganasa18/attandance-be-mobile/internal/model/web"
	service "ganasa18/attandance-be-mobile/internal/service/crud"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserMenuAccessControllerImpl struct {
	UserMenuAccessService service.UserMenuAccessService
}

func NewUserMenuAccessController(userMenuAccessService service.UserMenuAccessService) UserMenuAccessController {
	return &UserMenuAccessControllerImpl{
		UserMenuAccessService: userMenuAccessService,
	}
}

func (controller *UserMenuAccessControllerImpl) CreateUserMenuAccess(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userMenuAccessRequest := web.UserMenuCreateAccessResponseRequest{}
	helper.ReadFromRequestBody(request, &userMenuAccessRequest)

	roleResponse := controller.UserMenuAccessService.CreateUserMenuAccess(request.Context(), userMenuAccessRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserMenuAccessControllerImpl) GetUserMenuAccess(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userMenuAccessResponse := controller.UserMenuAccessService.GetUserMenuAccess(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userMenuAccessResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
