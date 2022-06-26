package login

import (
	"context"
	login "github.com/starbuy-commerce/auth-server/protobuf/protobuf_login"
)

type Server struct {
	login.UnimplementedLoginServiceServer
}

func (s *Server) Login(ctx context.Context, request *login.LoginRequest) (*login.LoginResponse, error) {
	panic("Implement me")
}
