package utils

import (
	"errors"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/system/request"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

var (
	TokenExpired      = errors.New("token is expired")
	TokenNotValidYet  = errors.New("token not active yet")
	TokenMalformed    = errors.New("that's not even a token")
	TokenInvalid      = errors.New("couldn't handle this token")
	SignatureNotFound = errors.New("signature for the token was not found")
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}

// CreateClaims .
func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	bf, _ := ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second),
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"ABC"},
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // Time when the token becomes valid
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // Expiration time (7 days) as configured
			Issuer:    global.GVA_CONFIG.JWT.Issuer,              // Issuer of the signature
		},
	}
	return claims
}

// CreateToken Create a token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	if j.SigningKey == nil {
		return "", SignatureNotFound
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken old token replaces new token using merge back to source to avoid concurrency issues
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.GVA_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken .
func (j *JWT) ParseToken(tk string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tk, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}

	return nil, TokenInvalid
}

// GetToken extracts id of the c.Header
func GetToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	s := strings.Split(authHeader, "Bearer")
	if len(s) != 2 {
		return ""
	}
	token := strings.TrimSpace(s[1])
	return token
}
