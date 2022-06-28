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

	username, ok := authorization.ValidateToken(token.Token)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Token inválido.", "user": "", "jwt": ""})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Token válido", "user": username, "jwt": token})
}
