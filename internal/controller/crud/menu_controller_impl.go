package controller

import (
	"context"
	"ganasa18/attandance-be-mobile/internal/helper"
	"ganasa18/attandance-be-mobile/internal/model/web"
	service "ganasa18/attandance-be-mobile/internal/service/crud"
	"math"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type MenuControllerImpl struct {
	MenuService service.MenuService
}

func NewMenuController(MenuService service.MenuService) MenuController {
	return &MenuControllerImpl{
		MenuService: MenuService,
	}
}

func (controller *MenuControllerImpl) CreateMenu(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	menuRequest := web.MenuMasterCreateResponseRequest{}
	helper.ReadFromRequestBody(request, &menuRequest)

	menuResponse := controller.MenuService.CreateMenu(request.Context(), menuRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MenuControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Get query parameters from the request
	queryParams := request.URL.Query()

	ctxWithParams := context.WithValue(request.Context(), "queryParams", queryParams)
	menuResponse, rowCount := controller.MenuService.FindAll(ctxWithParams)

	// Retrieve total count and calculate total pages
	countTotal := rowCount

	perPage, _ := strconv.Atoi(queryParams.Get("limit"))
	currentPage, _ := strconv.Atoi(queryParams.Get("page"))
	totalPage := 0
	if countTotal > 0 {
		totalPage = int(math.Ceil(float64(countTotal) / float64(perPage)))
	}

	if len(menuResponse) == 0 {
		countTotal = len(menuResponse)
	}

	totalCount := countTotal

	var data []web.MenuMasterResponseRequest
	if menuResponse != nil {
		data = menuResponse
	} else {
		data = []web.MenuMasterResponseRequest{} // Return an empty array if subCategoryResponse is nil
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
