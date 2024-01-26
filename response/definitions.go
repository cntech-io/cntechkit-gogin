package response

import (
	"net/http"

	errormessage "github.com/cntech-io/cntechkit-gogin/v2/error-message"
)

type response struct {
	status         bool
	endUserMessage string
	err            errormessage.ErrorMessage
	data           any
}

func New() *response {
	return &response{}
}

func (r *response) BadRequest(err errormessage.ErrorMessage) (int, response) {
	r.status = false
	r.err = err
	return http.StatusBadRequest, *r
}

func (r *response) Forbitten(err errormessage.ErrorMessage) (int, response) {
	r.status = false
	r.err = err
	return http.StatusForbidden, *r
}

func (r *response) Unauthorized(err errormessage.ErrorMessage) (int, response) {
	r.status = false
	r.err = err
	return http.StatusUnauthorized, *r
}

func (r *response) NotFound(err errormessage.ErrorMessage) (int, response) {
	r.status = false
	r.err = err
	return http.StatusNotFound, *r
}

func (r *response) NoContent(message string) (int, response) {
	r.status = true
	r.endUserMessage = message
	return http.StatusNoContent, *r
}

func (r *response) Created(message string) (int, response) {
	r.status = true
	r.endUserMessage = message
	return http.StatusCreated, *r
}

func (r *response) OK(message string) (int, response) {
	r.status = true
	r.endUserMessage = message
	return http.StatusOK, *r
}

func (r *response) Data(message string, data any) (int, response) {
	r.status = true
	r.data = data
	return http.StatusOK, *r
}
