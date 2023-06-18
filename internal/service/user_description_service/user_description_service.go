package user_description_service

import "context"

type UserDescriptionService interface {
	SaveDescription(ctx context.Context, uid, description string) error
	GetUserDescription(ctx context.Context, uid string) (string, error)
}
