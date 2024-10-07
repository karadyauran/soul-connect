package server

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/emptypb"
	"soul-connect/sc-auth/internal/generated"
	"soul-connect/sc-auth/internal/models"
	"soul-connect/sc-auth/internal/services"
)

type AuthServer struct {
	generated.UnimplementedAuthServiceServer
	authService *services.AuthService
}

func NewAuthServer(service *services.Service) *AuthServer {
	return &AuthServer{
		authService: service.AuthService,
	}
}

func (s *AuthServer) Register(ctx context.Context, request *generated.RegisterUserRequest) (*generated.RegisterUserResponse, error) {
	createUserReq := models.CreateUserRequest{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	registeredUser, err := s.authService.Register(ctx, createUserReq)
	if err != nil {
		return nil, err
	}

	response := &generated.RegisterUserResponse{
		Id:       registeredUser.ID,
		Username: registeredUser.Username,
		Email:    registeredUser.Email,
	}

	return response, nil
}

func (s *AuthServer) Login(ctx context.Context, request *generated.LoginUserRequest) (*generated.LoginUserResponse, error) {
	loginReq := models.UserLoginRequest{
		Username: request.Username,
		Password: request.Password,
	}

	loginUser, err := s.authService.Login(ctx, loginReq)
	if err != nil {
		return nil, err
	}

	response := &generated.LoginUserResponse{
		Id:           loginUser.ID,
		Username:     loginUser.Username,
		Email:        loginUser.Email,
		AccessToken:  loginUser.AccessToken,
		RefreshToken: loginUser.RefreshToken,
	}

	return response, nil
}

func (s *AuthServer) Logout(ctx context.Context, request *generated.LogoutUserRequest) (*emptypb.Empty, error) {
	if err := s.authService.Logout(ctx, request.Token); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *AuthServer) LogoutFromAllDevices(ctx context.Context, request *generated.LogoutUserFromAllDevicesRequest) (*emptypb.Empty, error) {
	userUUID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}
	userID := pgtype.UUID{Bytes: [16]byte(userUUID[:]), Valid: true}
	if err := s.authService.LogoutFromAllDevices(ctx, userID); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *AuthServer) DeleteUser(ctx context.Context, request *generated.DeleteUserRequest) (*emptypb.Empty, error) {
	userUUID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}
	userID := pgtype.UUID{Bytes: [16]byte(userUUID[:]), Valid: true}
	if err := s.authService.DeleteUser(ctx, userID); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
