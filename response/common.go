package response

import "net/http"

type Http struct {
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

func SuccessResponse(data any) Http {
	return Http{
		Status:      http.StatusOK,
		ErrorSchema: "Success",
		Payload:     data,
		Validation:  nil,
	}
}

func ErrorResponse() Http {
	return Http{
		Status:      http.StatusInternalServerError,
		ErrorSchema: "Internal Server Error",
		Payload:     nil,
		Validation:  nil,
	}
}

func FailedResponse(msg string, status int, validation *[]map[string]string) Http {
	return Http{
		Status:      status,
		ErrorSchema: msg,
		Payload:     nil,
		Validation:  validation,
	}
}
