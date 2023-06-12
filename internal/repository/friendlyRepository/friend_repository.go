package repository

import (
	"context"
	"friendly/internal/model"
)

type FriendsRepository interface {
	GetFriendsByUserUid(ctx context.Context, userUid string, limit, page int) ([]string, error)
	AddFriendRequest(ctx context.Context, userUid, friendUid string) (model.Friends, error)
	AddFriend(ctx context.Context, userUid, friendUid string) (model.Friends, error)
	DeleteFriend(ctx context.Context, userUid, friendUid string) error
}
