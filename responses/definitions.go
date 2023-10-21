package responses

import (
	"net/http"

	"github.com/cntech-io/cntechkit-gogin/errorcodes"
)

type response struct {
	Status         bool                 `json:"status"`
	EndUserMessage string               `json:"end_user_message"`
	Err            errorcodes.ErrorCode `json:"error,omitempty"`
	Data           any                  `json:"data,omitempty"`
}

func New() *response {
	return &response{}
}

func (r *response) BadRequest(errcode errorcodes.ErrorCode) (int, response) {
	r.Status = false
	r.Err = errcode
	return http.StatusBadRequest, *r
}

func (r *response) Unauthorized(errcode errorcodes.ErrorCode) (int, response) {
	r.Status = false
	r.Err = errcode
	return http.StatusUnauthorized, *r
}

func (r *response) NotFound(errcode errorcodes.ErrorCode) (int, response) {
	r.Status = false
	r.Err = errcode
	return http.StatusNotFound, *r
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
