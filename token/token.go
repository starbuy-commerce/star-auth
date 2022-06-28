package token

import (
	"github.com/gin-gonic/gin"
	"github.com/starbuy-commerce/auth-server/authorization"
	"net/http"
)

type tokenDTO struct {
	Token string `json:"token"`
}

func ValidateToken(c *gin.Context) {

	token := tokenDTO{}
	if err := c.BindJSON(&token); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request", "user": "", "jwt": ""})
		return
	}

	username, err := authorization.ExtractUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "invalid token"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Token válido", "user": username, "jwt": token})
}
