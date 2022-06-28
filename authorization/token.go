package authorization

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"strings"
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

func extractToken(c *gin.Context) string {
	raw := c.GetHeader("Authorization")

	if len(strings.Split(raw, " ")) != 2 {
		return ""
	}
	return strings.Split(raw, " ")[1]
}

func checkSecurityKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %s", token.Header["alg"])
	}

	return []byte(os.Getenv("JWT_SIGN")), nil
}

func ExtractUser(c *gin.Context) (string, error) {
	raw := extractToken(c)
	token, err := jwt.Parse(raw, checkSecurityKey)
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := fmt.Sprintf("%v", claims["username"])
		return username, nil
	}

	return "", errors.New("Invalid token")
}
