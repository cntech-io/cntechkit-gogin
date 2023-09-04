package cntechkitgogin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	Error          string `json:"error,omitempty"`
	Description    string `json:"description,omitempty"`
}

func NewResponse(endUserMessage string) *response {
	return &response{
		EndUserMessage: endUserMessage,
	}
}

func (r *response) Err(err string) *response {
	r.Error = err
	return r
}

func (r *response) success() *response {
	r.Status = true
	return r
}

func (r *response) fail(errType errorType) *response {
	r.Status = false
	r.ErrorType = string(errType)
	return r
}
func (r *response) successWithData(data any) *response {
	r.Data = data
	return r
}

func (r *response) OK(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, r.success())
}

func (r *response) Created(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, r.success())
}

func (r *response) BadRequest(ctx *gin.Context, errType errorType) {
	ctx.JSON(http.StatusBadRequest, r.fail(errType))
}

func (r *response) Unauthorized(ctx *gin.Context, errType errorType) {
	ctx.JSON(http.StatusUnauthorized, r.fail(errType))
}

func (r *response) Forbidden(ctx *gin.Context, errType errorType) {
	ctx.JSON(http.StatusForbidden, r.fail(errType))
}

func (r *response) NotFound(ctx *gin.Context, errType errorType) {
	ctx.JSON(http.StatusNotFound, r.fail(errType))
}

func (r *response) WithData(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, r.successWithData(data))
}
