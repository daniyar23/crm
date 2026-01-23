package grpc

import (
	"github.com/daniyar23/crm/internal/feature/core"
	userpb "github.com/daniyar23/crm/internal/feature/user/delivery/grpc"
)

func toProtoUser(u core.User) *userpb.UserResponse {
	return &userpb.UserResponse{
		Id:    uint64(u.ID),
		Name:  u.Name,
		Email: u.Email,
	}
}
