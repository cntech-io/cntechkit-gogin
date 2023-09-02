package cntechkitgogin

import "net/http"

type errorType string

const (
	InvalidParams  errorType = "invalid_params"
	QueryError     errorType = "query_error"
	AlgorthymError errorType = "algorthym_error"
)

type response struct {
	Status         bool   `json:"status"`
	EndUserMessage string `json:"end_user_message"`
	ErrorType      string `json:"error_type,omitempty"`
	Data           any    `json:"data,omitempty"`
}

func Response(endUserMessage string) *response {
	return &response{
		EndUserMessage: endUserMessage,
	}
}

func (r *response) Success() *response {
	r.Status = true
	return r
}

func (r *response) Fail(respType errorType) *response {
	r.Status = false
	r.ErrorType = string(respType)
	return r
}
func (r *response) WithData(data any) *response {
	r.Data = data
	return r
}

func OKResponse(endUserMessage string) (int, response) {
	return http.StatusOK, *Response(endUserMessage).Success()
}

func DataResponse(endUserMessage string, data any) (int, response) {
	return http.StatusOK, *Response(endUserMessage).Success().WithData(data)
}

func CreatedResponse(endUserMessage string) (int, response) {
	return http.StatusCreated, *Response(endUserMessage).Success()
}

func ForbittenResponse(endUserMessage string, respType errorType) (int, response) {
	return http.StatusForbidden, *Response(endUserMessage).Fail(respType)
}

func BadRequestResponse(endUserMessage string, respType errorType) (int, response) {
	return http.StatusBadRequest, *Response(endUserMessage).Fail(respType)
}

func UnauthorizedResponse(endUserMessage string, respType errorType) (int, response) {
	return http.StatusUnauthorized, *Response(endUserMessage).Fail(respType)
}
