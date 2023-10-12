package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const AUTHORIZATION = "authorization"
const UserId = "user_id"

var (
	jwtSecret        = ""
	jwtSecretRefresh = ""
)

func MustInit(secret, secretRefresh string) {
	jwtSecret = secret
	jwtSecretRefresh = secretRefresh
}

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, username string) (string, string, error) {
	if jwtSecret == "" {
		return "", "", errors.New("no jwt secret")
	}
	if jwtSecretRefresh == "" {
		return "", "", errors.New("no jwt refresh secret")
	}
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)

	claims := Claims{
		userID,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expireTime},
			Issuer:    "bic-gin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err1 := tokenClaims.SignedString(jwtSecret)
	if err1 != nil {
		return "", "", err1
	}
	tokenRefresh, err2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	return token, tokenRefresh, err2
}

func ParseRefreshToken(token string) (*Claims, error) {
	if token == "" {
		return nil, errors.New("empty token")
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretRefresh, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func ParseToken(token string) (*Claims, error) {
	if token == "" {
		return nil, errors.New("empty token")
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
