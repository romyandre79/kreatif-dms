package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenPayload struct {
	UserID       uuid.UUID `json:"user_id"`
	Role         string    `json:"role"`
	DepartmentID uuid.UUID `json:"department_id"`
}

type Claims struct {
	UserID       uuid.UUID `json:"user_id"`
	Role         string    `json:"role"`
	DepartmentID uuid.UUID `json:"department_id"`
	jwt.RegisteredClaims
}

func GenerateToken(payload TokenPayload, secret string, ttl time.Duration) (string, error) {
	claims := Claims{
		UserID:       payload.UserID,
		Role:         payload.Role,
		DepartmentID: payload.DepartmentID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenStr string, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
