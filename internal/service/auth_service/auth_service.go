package auth_service

import "context"

type AuthService interface {
	GetUserUid(ctx context.Context, token string) (string, error)
}
