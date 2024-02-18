package jwt

import (
	"encoding/json"
	"errors"

	errormessage "github.com/cntech-io/cntechkit-gogin/v2/error-message"
	"github.com/cntech-io/cntechkit-gogin/v2/response"
	"github.com/cntech-io/cntechkit-gogin/v2/utility"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type jwtConfig struct {
	SecretKey string
	claims    *jwt.Claims
}

func NewJwt(secret string) *jwtConfig {
	return &jwtConfig{
		SecretKey: secret,
	}
}

func (cfg *jwtConfig) AddClaims(claims jwt.Claims) *jwtConfig {
	cfg.claims = &claims
	return cfg
}

func (cfg *jwtConfig) sign() string {
	var token *jwt.Token
	if cfg.claims == nil {
		token = jwt.New(jwt.SigningMethodHS256)
	} else {
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, *cfg.claims)
	}

	tokenString, _ := token.SignedString([]byte(cfg.SecretKey))
	return tokenString
}

func (cfg *jwtConfig) authenticate(token string) (jwt.Claims, error) {
	var claims jwt.Claims
	if cfg.claims == nil {
		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.SecretKey), nil
		})
		if err != nil {
			return nil, err
		}
		claims = parsedToken.Claims
		if !parsedToken.Valid {
			return nil, errors.New("invalid token")
		}
		return claims, nil
	}

	parsedToken, err := jwt.ParseWithClaims(token, *cfg.claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}
	claims = parsedToken.Claims
	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func (cfg *jwtConfig) AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utility.GetBearerTokenFromContext(c)
		if err != nil {
			c.AbortWithStatusJSON(response.New().BadRequest(errormessage.ERR_INVALID_BEARER_AUTH_HEADER))
			return
		}

		claims, err := cfg.authenticate(token)
		if err != nil {
			c.AbortWithStatusJSON(response.New().BadRequest(errormessage.ERR_USER_NOT_AUTHENTICATED))
			return
		}

		claimsBytes, _ := json.Marshal(claims)
		var claimsData map[string]interface{}
		_ = json.Unmarshal(claimsBytes, &claimsData)
		c.Set("claims", claimsData)

		c.Next()
	}
}

func (cfg *jwtConfig) GenerateToken() string {
	return cfg.sign()
}
