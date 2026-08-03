package utils

import (
    "errors"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func InitJWTSecret(secret string) {
    jwtSecret = []byte(secret)
}

type Claims struct {
    UserID int `json:"user_id"`
    jwt.RegisteredClaims
}

// GenerateJWT ek naya token banata hai jisme userID aur expiry time hota hai
func GenerateJWT(userID int) (string, error) {
    claims := Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 ghante valid
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// ValidateJWT token ko verify karta hai aur userID nikalta hai
func ValidateJWT(tokenStr string) (int, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil || !token.Valid {
        return 0, errors.New("invalid or expired token")
    }

    return claims.UserID, nil
}