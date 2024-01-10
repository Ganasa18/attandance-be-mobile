package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MenuController interface {
	CreateMenu(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
