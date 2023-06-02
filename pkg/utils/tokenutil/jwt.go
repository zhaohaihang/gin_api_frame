package tokenutil

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"time"
)

var jwtSecret = []byte("FanOne")

type Claims struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		UserID:    id,
		Username:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin_api_frame",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(tokenStr string) (*Claims, error) {
	token := strings.Fields(tokenStr)
	if len(token) != 2 || strings.ToLower(token[0]) != "bearer" || token[1] == "" {
		return nil, fmt.Errorf("authorization header invaild")
	}

	tokenClaims, err := jwt.ParseWithClaims(token[1], &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func SetTokenClaimsToContext(c *gin.Context, claims *Claims) {
	if c == nil || claims == nil {
		return
	}
	c.Set("claims", claims)
}

func GetTokenClaimsFromContext(c *gin.Context) *Claims {
	if c == nil {
		return nil
	}

	val, ok := c.Get("claims")
	if !ok {
		return nil
	}

	claims, ok := val.(*Claims)
	if !ok {
		return nil
	}

	return claims
}
