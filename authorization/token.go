package authorization

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"time"
)

func GenerateToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   username,
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
		"nbf":        time.Date(2022, 6, 26, 20, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SIGN")))

	if err != nil {
		log.Fatalf("Failed while signing token: %v", err.Error())
	}

	return tokenString
}

func ValidateToken(tokenString string) (string, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("PORT")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return fmt.Sprintf("%v", claims["username"]), true
	}

	if err != nil {
		log.Fatalf("Error while validating token: %v", err.Error())
	}

	return "", false
}
