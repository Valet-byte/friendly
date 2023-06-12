package repository

import (
	"context"
	"friendly/internal/model"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByUid(ctx context.Context, uid string) (model.User, error)
	DeleteUser(ctx context.Context, uid string) error
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}
