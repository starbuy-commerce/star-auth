package token

import (
	"context"
	token "github.com/starbuy-commerce/auth-server/protobuf/protobuf_token"
)

type Server struct {
	token.UnimplementedTokenValidationServiceServer
}

func (t *Server) ValidateToken(ctx context.Context, request *token.TokenValidationRequest) (*token.TokenValidationResponse, error) {
	panic("implement me")
}
