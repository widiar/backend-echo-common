package response

import "net/http"

type Common struct {
	Status      int    `json:"status"`
	ErrorSchema string `json:"error_schema"`
	Payload     any    `json:"payload"`
	Validation  any    `json:"validation"`
	Total       int    `json:"total"`
}

type RequestValidationError struct {
	Message string
	Details []map[string]string
}

func (e *RequestValidationError) Error() string {
	return e.Message
}

func SuccessResponse(data any) Common {
	return Common{
		Status:      http.StatusOK,
		ErrorSchema: "Success",
		Payload:     data,
		Validation:  nil,
	}
}

func ErrorResponse() Common {
	return Common{
		Status:      http.StatusInternalServerError,
		ErrorSchema: "Internal Server Error",
		Payload:     nil,
		Validation:  nil,
	}
}

func FailedResponse(msg string, status int, validation *[]map[string]string) Common {
	return Common{
		Status:      status,
		ErrorSchema: msg,
		Payload:     nil,
		Validation:  validation,
	}
}
