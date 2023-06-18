package user_friends_repository

import (
	"context"
	"friendly/internal/model"
)

type FriendsRepository interface {
	GetFriendsByUserUid(ctx context.Context, userUid string, limit, page int) ([]model.User, error)
	AddFriend(ctx context.Context, userUid, friendUid string) error
	DeleteFriend(ctx context.Context, userUid, friendUid string) error
}
