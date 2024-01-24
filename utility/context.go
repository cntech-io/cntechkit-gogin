package utility

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func getHeader(c *gin.Context, key string) ([]string, error) {
    value := c.Request.Header[key]
    if len(value) == 0 {
        return nil, fmt.Errorf("%s header not found", key)
    }
    return value, nil
}

func GetClientIPFromContext(c *gin.Context) (string, error) {
	xForwardedFor, err := getHeader(c, "X-Forwarded-For")
	if err != nil {
		return "", err
	}

	IPs := strings.Split(xForwardedFor[0], ",")
	if len(IPs) == 0 {
		return "", errors.New("X-Forwarded-For values not found")
	}

	clientIP := strings.TrimSpace(IPs[0])
	return clientIP, nil
}

func GetBearerTokenFromContext(c *gin.Context) (string, error) {
	authHeader, err := getHeader(c, "Authorization")
	if err != nil {
		return "", err
	}

	bearerToken := strings.Fields(authHeader[0])
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		return "", fmt.Errorf("invalid bearer token")
	}

	return bearerToken[1], nil
}

func GetCustomHeaderValueFromContext(c *gin.Context, key string) (string, error) {
	valueArr, err := getHeader(c, key)
	if err != nil {
		return "", err
	}
	return valueArr[0], nil
}

func ValidateContextBody(c *gin.Context, obj interface{}) (bool, error) {
	if err := c.BindJSON(obj); err != nil {
		return false, err
	}

	if err := validator.New().Struct(obj); err != nil {
		return false, err
	}

	return true, nil
}

func ValidateContextForm(c *gin.Context, obj interface{}) (bool, error) {
	if err := c.ShouldBindQuery(obj); err != nil {
		return false, err
	}

	if err := validator.New().Struct(obj); err != nil {
		return false, err
	}

	return true, nil
}

func ValidateContextQuery(c *gin.Context, obj interface{}) (bool, error) {
	if err := c.ShouldBindUri(obj); err != nil {
		return false, err
	}

	if err := validator.New().Struct(obj); err != nil {
		return false, err
	}

	return true, nil
}
