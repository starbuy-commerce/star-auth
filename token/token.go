package token

import (
	"context"
	"github.com/starbuy-commerce/auth-server/authorization"
	token "github.com/starbuy-commerce/auth-server/protobuf/protobuf_token"
)

type Server struct {
	token.UnimplementedTokenValidationServiceServer
}

func (t *Server) ValidateToken(ctx context.Context, request *token.TokenValidationRequest) (*token.TokenValidationResponse, error) {
	username, ok := authorization.ValidateToken(request.GetToken())
	return &token.TokenValidationResponse{Username: username, Value: ok}, nil
}
