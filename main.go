package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/starbuy-commerce/auth-server/database"
	"github.com/starbuy-commerce/auth-server/login"
	token "github.com/starbuy-commerce/auth-server/token"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if err := database.Connect(); err != nil {
		log.Fatalf("Failed while connecting to database: %v", err.Error())
		return
	}

	r := gin.Default()

	r.POST("/login", login.Login)
	r.POST("/token", token.ValidateToken)

	r.Run(fmt.Sprintf(":%v", port))

}
