package controller

import (
	"context"
	"ganasa18/attandance-be-mobile/helper"
	"ganasa18/attandance-be-mobile/model/web"
	"ganasa18/attandance-be-mobile/service"
	"math"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Get query parameters from the request
	queryParams := request.URL.Query()

	ctxWithParams := context.WithValue(request.Context(), "queryParams", queryParams)
	userResponse, rowCount := controller.UserService.FindAll(ctxWithParams)

	// Retrieve total count and calculate total pages
	countTotal := rowCount

	perPage, _ := strconv.Atoi(queryParams.Get("limit"))
	currentPage, _ := strconv.Atoi(queryParams.Get("page"))
	totalPage := 0
	if countTotal > 0 {
		totalPage = int(math.Ceil(float64(countTotal) / float64(perPage)))
	}

	if len(userResponse) == 0 {
		countTotal = len(userResponse)
	}

	totalCount := countTotal

	var data []web.UserResponseRequest
	if userResponse != nil {
		data = userResponse
	} else {
		data = []web.UserResponseRequest{} // Return an empty array if subCategoryResponse is nil
	}

	pageInfo := web.PageInfoResponse{
		Total:       totalCount,
		PerPage:     perPage,
		CurrentPage: currentPage,
		TotalPage:   totalPage,
	}

	webResponse := web.WebResponsePaginate{
		Code:     200,
		Status:   "OK",
		Data:     data,
		PageInfo: pageInfo,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
