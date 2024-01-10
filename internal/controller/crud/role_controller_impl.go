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

type RoleControllerImpl struct {
	RoleService service.RoleService
}

func NewRoleController(roleService service.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: roleService,
	}
}

func (controller *RoleControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Get query parameters from the request
	queryParams := request.URL.Query()

	ctxWithParams := context.WithValue(request.Context(), "queryParams", queryParams)
	roleResponse, rowCount := controller.RoleService.FindAll(ctxWithParams)

	// Retrieve total count and calculate total pages
	countTotal := rowCount

	perPage, _ := strconv.Atoi(queryParams.Get("limit"))
	currentPage, _ := strconv.Atoi(queryParams.Get("page"))
	totalPage := 0
	if countTotal > 0 {
		totalPage = int(math.Ceil(float64(countTotal) / float64(perPage)))
	}

	if len(roleResponse) == 0 {
		countTotal = len(roleResponse)
	}

	totalCount := countTotal

	var data []web.RoleMasterResponseRequest
	if roleResponse != nil {
		data = roleResponse
	} else {
		data = []web.RoleMasterResponseRequest{} // Return an empty array if subCategoryResponse is nil
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

func (controller *RoleControllerImpl) CreateRole(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleRequest := web.RoleMasterCreateResponseRequest{}
	helper.ReadFromRequestBody(request, &roleRequest)

	roleResponse := controller.RoleService.CreateRole(request.Context(), roleRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
