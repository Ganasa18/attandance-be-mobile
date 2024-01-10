package exception

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

// type CustomFoundError struct {
// 	Code   int    `json:"code"`
// 	Status string `json:"status"`
// 	Error  string `json:"error"`
// }

// func CustomHandleError(code int, error string) CustomFoundError {
// 	return CustomFoundError{
// 		Code:  code,
// 		Error: error,
// 	}
// }
