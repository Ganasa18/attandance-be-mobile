package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type RoleAccessController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CreateRoleAccess(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
