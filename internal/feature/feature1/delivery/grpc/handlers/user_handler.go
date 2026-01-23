package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/daniyar23/crm/internal/feature/feature1/services"
	userpb "github.com/daniyar23/crm/proto"
)

type UserGRPCHandler struct {
	userpb.UnimplementedUserServiceServer
	service services.UserService
}

func NewUserGRPCHandler(s services.UserService) *UserGRPCHandler {
	return &UserGRPCHandler{service: s}
}

func (h *UserGRPCHandler) GetUserByID(
	ctx context.Context,
	req *userpb.GetUserByIDRequest,
) (*userpb.UserResponse, error) {

	user, err := h.service.GetUserByID(ctx, uint(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return toProtoUser(user), nil
}

func (h *UserGRPCHandler) ListUsers(
	ctx context.Context,
	_ *userpb.Empty,
) (*userpb.UserListResponse, error) {

	users, err := h.service.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list users")
	}

	resp := &userpb.UserListResponse{}
	for _, u := range users {
		resp.Users = append(resp.Users, toProtoUser(u))
	}

	return resp, nil
}
