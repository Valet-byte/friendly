package friends_service

import (
	"context"
	"friendly/internal/model"
	"friendly/internal/repository/user_friends_repository"
	"friendly/internal/service/friends_token_service"
	"friendly/internal/utils"
)

type DefaultUserFriendsService struct {
	repo         user_friends_repository.FriendsRepository
	tokenService friends_token_service.FriendsTokenService
}

func (d *DefaultUserFriendsService) GetFriends(ctx context.Context, uid string, limit, page int) ([]model.User, error) {
	return d.repo.GetFriendsByUserUid(ctx, uid, limit, page)
}

func (d *DefaultUserFriendsService) AddFriends(ctx context.Context, userUid, friendUid string) error {
	return d.repo.AddFriend(ctx, userUid, friendUid)
}

func (d *DefaultUserFriendsService) DeleteFriends(ctx context.Context, userUid, friendUid string) error {
	return d.repo.DeleteFriend(ctx, userUid, friendUid)
}

func (d *DefaultUserFriendsService) GetFriendsToken(uid string) (string, error) {
	return d.tokenService.GetFriendsToken(uid)
}

func (d *DefaultUserFriendsService) AddFriendsByFriendsToken(ctx context.Context, uid, token string) error {
	friendId, err := d.tokenService.ParseFriendsToken(token)
	if err != nil {
		utils.Log("Failed add user by friend token", err)
		return err
	}

	return d.AddFriends(ctx, uid, friendId)
}
