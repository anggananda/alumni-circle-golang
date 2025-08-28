package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	IDAlumni int64  `json:"id_alumni"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	IDAlumni string `json:"id_alumni"`
	jwt.StandardClaims
}

func GenerateJWT(IDAlumni int64, username string) (string, int64, error) {
	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	JWTDUration := os.Getenv("JWT_DURATION")
	durationHours := 8
	if parsedDuration, err := strconv.Atoi(JWTDUration); err == nil && parsedDuration > 0 {
		durationHours = parsedDuration
	}

	expirationTime := time.Now().Add(time.Duration(durationHours) * time.Hour).Unix() // Access token berlaku selama 8 jam

	claims := &Claims{
		IDAlumni: IDAlumni,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expirationTime, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
