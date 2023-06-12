package service

import (
	"context"
	"fmt"
	cache "friendly/internal/cache/user_cache"
	"friendly/internal/model"
	repository "friendly/internal/repository/userRepository"
	"friendly/internal/utils"
)

type DefaultUserService struct {
	repo      repository.UserRepository
	userCache cache.UserCache
}

func NewDefaultUserService(repo repository.UserRepository, userCache cache.UserCache) *DefaultUserService {
	return &DefaultUserService{repo: repo, userCache: userCache}
}

func (d *DefaultUserService) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	res, err := d.repo.SaveUser(ctx, user)

	if err != nil {
		utils.Log("Failed save user / update user", err)
		return model.User{}, err
	}

	if err = d.userCache.PutUser(res); err != nil {
		utils.Log(fmt.Sprintf("Failed add user to cache. uset: { %v }", user), err)
	}

}

func (d *DefaultUserService) DeleteUser(ctx context.Context, uid string) error {
	//TODO implement me
	panic("implement me")
}

func (d *DefaultUserService) GetUser(ctx context.Context, uid string) (model.User, error) {
	//TODO implement me
	panic("implement me")
}
