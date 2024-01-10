package helper

import (
	"encoding/json"
	"ganasa18/attandance-be-mobile/internal/model/web"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)

}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}

func CustomResponseError(writer http.ResponseWriter, request *http.Request, statusCode int, err interface{}) {
	// Set the provided custom status code
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	// Create a generic error response
	webResponse := web.WebResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   err,
	}

	encoder := json.NewEncoder(writer)
	if encodeErr := encoder.Encode(webResponse); encodeErr != nil {
		PanicIfError(encodeErr)
	}
}
