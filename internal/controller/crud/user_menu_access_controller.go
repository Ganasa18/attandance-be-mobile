package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserMenuAccessController interface {
	CreateUserMenuAccess(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetUserMenuAccess(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
