package service

import (
	"context"
	"friendly/internal/model"
)

type UserService interface {
	UpdateUser(ctx context.Context, user model.User) (model.User, error)
	DeleteUser(ctx context.Context, uid string) error
	GetUserByUid(ctx context.Context, uid string) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}
