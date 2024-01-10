package middleware

import (
	"ganasa18/attandance-be-mobile/internal/helper"
	"ganasa18/attandance-be-mobile/internal/model/web"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CustomAuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

		providedToken := request.Header.Get("Authorization")

		// CHECK HANDLER IF AUTHORIZATION HEADER EMPTY
		if providedToken == "" {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "TOKEN NOT PROVIDED",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}

		tokenValid, _ := helper.ValidateToken(providedToken)

		// CHECK HANDLER TOKEN INVALID
		if tokenValid == nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}
			helper.WriteToResponseBody(writer, webResponse)
		} else {
			next(writer, request, params)
		}
	}
}
