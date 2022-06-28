package login

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/starbuy-commerce/auth-server/authorization"
	"github.com/starbuy-commerce/auth-server/database"
	"github.com/starbuy-commerce/auth-server/security"
	"net/http"
)

type loginData struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

func Login(c *gin.Context) {
	loginInfo := loginData{}
	if err := c.BindJSON(&loginInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request", "user": "", "jwt": ""})
		return
	}

	db := database.GrabDB()
	var userLogin loginData
	if err := db.Get(&userLogin, "SELECT * FROM login WHERE username=$1", loginInfo.Username); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Usuário não encontrado", "user": "", "jwt": ""})
			return
		}
	}

	if err := security.ComparePassword(userLogin.Password, loginInfo.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Senha incorreta", "user": "", "jwt": ""})
		return
	}

	token := authorization.GenerateToken(userLogin.Username)
	c.JSON(http.StatusOK, gin.H{"status": false, "message": "Senha incorreta", "user": userLogin.Username, "jwt": token})
}
