package user_service

import (
	"context"
	"fmt"
	cache "friendly/internal/cache/user_cache"
	"friendly/internal/model"
	repository "friendly/internal/repository/user_repository"
	"friendly/internal/utils"
)

type DefaultUserService struct {
	userRepo  repository.UserRepository
	userCache cache.UserCache
}

func NewDefaultUserService(repo repository.UserRepository, userCache cache.UserCache) *DefaultUserService {
	return &DefaultUserService{userRepo: repo, userCache: userCache}
}

func (d *DefaultUserService) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	res, err := d.userRepo.SaveUser(ctx, user)

	if err != nil {
		utils.Log("Failed save user / update user", err)
		return model.User{}, err
	}

	if err = d.userCache.PutUser(res); err != nil {
		utils.Log(fmt.Sprintf("Failed add user to cache. user: { %v }", user), err)
	}
	return res, nil

}

func (d *DefaultUserService) DeleteUser(ctx context.Context, uid string) error {
	err := d.userRepo.DeleteUser(ctx, uid)

	if err != nil {
		utils.Log("Failed delete user from cache", err)
		return err
	}

	err = d.userCache.DeleteUser(uid)
	if err != nil {
		utils.Log("Failed delete user from cache", err)
		return err
	}

	return nil
}

func (d *DefaultUserService) GetUserByUid(ctx context.Context, uid string) (model.User, error) {
	res, err := d.userCache.GetUser(uid)

	if err != nil {
		utils.Log("Failed get user from cache", err)
	} else {
		return res, nil
	}

	res, err = d.userRepo.GetUserByUid(ctx, uid)

	if err != nil {
		utils.Log("Failed get user from db", err)
		return res, err
	}

	if err = d.userCache.PutUser(res); err != nil {
		utils.Log(fmt.Sprintf("Failed add user to cache. user: { %v }", res), err)
	}

	return res, nil
}

func (d *DefaultUserService) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	res, err := d.userRepo.GetUserByEmail(ctx, email)

	if err != nil {
		utils.Log("Failed get user from db", err)
		return res, err
	}

	if err = d.userCache.PutUser(res); err != nil {
		utils.Log(fmt.Sprintf("Failed add user to cache. user: { %v }", res), err)
	}

	return res, nil
}
