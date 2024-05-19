package grpcapi

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	authpb "user/gen/go/auth"
	"user/pkg/utils"
)

type ServerAPI struct {
	authpb.UnimplementedAuthServiceServer
	port int
}

func NewServerAPI() *ServerAPI {
	return &ServerAPI{
		port: 0,
	}
}

func (s *ServerAPI) VerifyToken(ctx context.Context, req *authpb.VerifyTokenRequest) (*authpb.VerifyTokenResponse, error) {
	if req.Token == "" {
		return &authpb.VerifyTokenResponse{
			Valid:   false,
			Message: "missing Authorization header",
		}, fmt.Errorf("missing Authorization header")
	}

	parts := strings.Split(req.Token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return &authpb.VerifyTokenResponse{
			Valid:   false,
			Message: "Invalid Authorization header format",
		}, fmt.Errorf("Invalid Authorization header format")
	}

	token := parts[1]

	userIDStr, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err)
		return &authpb.VerifyTokenResponse{
			Valid:   false,
			Message: err.Error(),
		}, err
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		log.Println(err)
		return &authpb.VerifyTokenResponse{
			Valid:   false,
			Message: err.Error(),
		}, err
	}

	return &authpb.VerifyTokenResponse{
		Valid:   true,
		UserID:  userID,
		Message: "Success",
	}, nil
}
