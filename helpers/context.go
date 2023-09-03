package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func getHeader(ctx *gin.Context, key string) ([]string, error) {
	value, exists := ctx.Request.Header[key]
	if !exists {
		return nil, fmt.Errorf("%s header not found", key)
	}
	return value, nil
}

func GetClientIPFromContext(ctx *gin.Context) (string, error) {
	xForwardedFor, err := getHeader(ctx, "X-Forwarded-For")
	if err != nil {
		return "", err
	}

	ips := strings.Split(xForwardedFor[0], ",")
	if len(ips) == 0 {
		return "", errors.New("X-Forwarded-For values not found")
	}

	clientIP := strings.TrimSpace(ips[0])
	return clientIP, nil
}

func GetBearerTokenFromContext(ctx *gin.Context) (string, error) {
	authorization, err := getHeader(ctx, "Authorization")
	if err != nil {
		return "", err
	}

	authValue := authorization[0]

	authValueArr := strings.Split(authValue, " ")
	if len(authValueArr) != 2 {
		return "", fmt.Errorf("invalid authorization value")
	}
	if authValueArr[0] != "Bearer" {
		return "", fmt.Errorf("invalid bearer authorization value")
	}

	token := authValueArr[1]
	return token, nil
}

func GetCustomHeaderValueFromContext(ctx *gin.Context, key string) (string, error) {
	valueArr, err := getHeader(ctx, key)
	if err != nil {
		return "", err
	}
	value := valueArr[0]

	return value, nil
}

func ValidateContextBody(ctx *gin.Context, obj interface{}) (bool, error) {
	if err := ctx.BindJSON(obj); err != nil {
		return false, err
	}

	validator := validator.New()
	err := validator.Struct(obj)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ValidateContextForm(c *gin.Context, obj interface{}) (bool, error) {
	if err := c.ShouldBindQuery(obj); err != nil {
		return false, err
	}

	validator := validator.New()
	if err := validator.Struct(obj); err != nil {
		return false, err
	}

	return true, nil
}

func ValidateContextQuery(c *gin.Context, obj interface{}) (bool, error) {
	if err := c.ShouldBindUri(obj); err != nil {
		return false, err
	}

	validator := validator.New()
	if err := validator.Struct(obj); err != nil {
		return false, err
	}

	return true, nil
}
