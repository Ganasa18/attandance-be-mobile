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

type RoleAccessControllerImpl struct {
	RoleAccessService service.RoleAccessService
}

func NewRoleAccessController(roleAccessService service.RoleAccessService) RoleAccessController {
	return &RoleAccessControllerImpl{
		RoleAccessService: roleAccessService,
	}
}

func (controller *RoleAccessControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Get query parameters from the request
	queryParams := request.URL.Query()

	ctxWithParams := context.WithValue(request.Context(), "queryParams", queryParams)
	roleAccessResponse, rowCount := controller.RoleAccessService.FindAll(ctxWithParams)

	// Retrieve total count and calculate total pages
	countTotal := rowCount

	perPage, _ := strconv.Atoi(queryParams.Get("limit"))
	currentPage, _ := strconv.Atoi(queryParams.Get("page"))
	totalPage := 0
	if countTotal > 0 {
		totalPage = int(math.Ceil(float64(countTotal) / float64(perPage)))
	}

	if len(roleAccessResponse) == 0 {
		countTotal = len(roleAccessResponse)
	}

	totalCount := countTotal

	var data []web.RoleAccessMasterResponseRequest
	if roleAccessResponse != nil {
		data = roleAccessResponse
	} else {
		data = []web.RoleAccessMasterResponseRequest{} // Return an empty array if subCategoryResponse is nil
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

func (controller *RoleAccessControllerImpl) CreateRoleAccess(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleAccessRequest := web.RoleAccessMasterCreateResponseRequest{}
	helper.ReadFromRequestBody(request, &roleAccessRequest)

	roleAccessResponse := controller.RoleAccessService.CreateRoleAccess(request.Context(), roleAccessRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleAccessResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
