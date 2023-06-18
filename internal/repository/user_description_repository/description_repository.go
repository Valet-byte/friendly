package user_description_repository

import "context"

type DescriptionRepository interface {
	UpdateUserDescription(ctx context.Context, uid, description string) error
	GetUserDescription(ctx context.Context, uid string) (string, error)
}
