package response

import (
	"net/http"

	errormessage "github.com/cntech-io/cntechkit-gogin/v2/error-message"
)

type response struct {
	Status         bool                       `json:"status"`
	EndUserMessage string                     `json:"endUserMessage"`
	Err            *errormessage.ErrorMessage `json:"error,omitempty"`
	Data           any                        `json:"data,omitempty"`
}

func New() *response {
	return &response{}
}

func (r *response) BadRequest(err errormessage.ErrorMessage) (int, response) {
	r.Status = false
	r.Err = &err
	return http.StatusBadRequest, *r
}

func (r *response) Forbitten(err errormessage.ErrorMessage) (int, response) {
	r.Status = false
	r.Err = &err
	return http.StatusForbidden, *r
}

func (r *response) Unauthorized(err errormessage.ErrorMessage) (int, response) {
	r.Status = false
	r.Err = &err
	return http.StatusUnauthorized, *r
}

func (r *response) NotFound(err errormessage.ErrorMessage) (int, response) {
	r.Status = false
	r.Err = &err
	return http.StatusNotFound, *r
}

func (r *response) TooManyRequests(err errormessage.ErrorMessage) (int, response) {
	r.Status = false
	r.Err = &err
	return http.StatusTooManyRequests, *r
}

func (r *response) NoContent(message string) (int, response) {
	r.Status = true
	r.EndUserMessage = message
	return http.StatusNoContent, *r
}

func (r *response) Created(message string) (int, response) {
	r.Status = true
	r.EndUserMessage = message
	return http.StatusCreated, *r
}

func (r *response) OK(message string) (int, response) {
	r.Status = true
	r.EndUserMessage = message
	return http.StatusOK, *r
}

func (r *response) WithData(message string, data any) (int, response) {
	r.Status = true
	r.EndUserMessage = message
	r.Data = data
	return http.StatusOK, *r
}
