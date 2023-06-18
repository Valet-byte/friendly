package friends_service

import (
	"context"
	"friendly/internal/model"
)

type FriendsService interface {
	GetFriends(ctx context.Context, uid string, limit, page int) ([]model.User, error)
	AddFriends(ctx context.Context, userUid, friendUid string) error
	DeleteFriends(ctx context.Context, userUid, friendUid string) error
	GetFriendsToken(uid string) (string, error)
	AddFriendsByFriendsToken(ctx context.Context, uid, token string) error
}
