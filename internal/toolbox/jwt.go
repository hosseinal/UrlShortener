package toolbox

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrorAccessExpired = errors.New("access token expired")
var ErrorInvalidToken = errors.New("invalid token")
var ErrorInvalidClaims = errors.New("invalid claims")

// JWTClaims represents the custom claims we want to include in our JWT
type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(Username string, accessSecret string, refreshSecret string) (string, string, error) {
	// 1. Generate short-lived Access Token (15 minutes)
	accessClaims := JWTClaims{
		Username: Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "UrlShortner",
		},
	}
	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := accessTokenObj.SignedString([]byte(accessSecret))
	if err != nil {
		return "", "", err
	}

	// 2. Generate long-lived Refresh Token (7 days)
	refreshClaims := JWTClaims{
		Username: Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "UrlShortner",
		},
	}
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refreshTokenObj.SignedString([]byte(refreshSecret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func ParseJWT(tokenString, secret string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrorAccessExpired
		}
		return nil, ErrorInvalidToken
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrorInvalidClaims
}

func ValidateJWT(tokenString, secret string) (bool, error) {
	_, err := ParseJWT(tokenString, secret)
	if err != nil {
		return false, ErrorInvalidToken
	}
	return true, ErrorInvalidToken
}
