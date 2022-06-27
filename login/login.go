package login

import (
	"context"
	"database/sql"
	"errors"
	"github.com/starbuy-commerce/auth-server/authorization"
	"github.com/starbuy-commerce/auth-server/database"
	login "github.com/starbuy-commerce/auth-server/protobuf/protobuf_login"
	"github.com/starbuy-commerce/auth-server/security"
)

type Server struct {
	login.UnimplementedLoginServiceServer
}

type loginData struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

func (s *Server) Login(ctx context.Context, request *login.LoginRequest) (*login.LoginResponse, error) {
	db := database.GrabDB()
	var userLogin loginData
	if err := db.Get(&userLogin, "SELECT * FROM login WHERE username=$1", request.Username); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Usuário não encontrado")
		}
	}

	if err := security.ComparePassword(userLogin.Password, request.Password); err != nil {
		return nil, errors.New("Senha incorreta")
	}

	token := authorization.GenerateToken(userLogin.Username)
	return &login.LoginResponse{
		Message: "Sessão iniciada com sucesso",
		Jwt:     token,
		Status:  true,
	}, nil
}
