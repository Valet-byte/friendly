package user_description_service

import (
	"context"
	"fmt"
	"friendly/internal/cache/user_description_cache"
	"friendly/internal/repository/user_description_repository"
	"friendly/internal/utils"
)

type DefaultUserDescriptionService struct {
	cache user_description_cache.UserDescriptionCache
	repo  user_description_repository.DescriptionRepository
}

func NewDefaultUserDescriptionService(repo user_description_repository.DescriptionRepository,
	cache user_description_cache.UserDescriptionCache) *DefaultUserDescriptionService {
	return &DefaultUserDescriptionService{
		cache: cache,
		repo:  repo,
	}
}

func (d *DefaultUserDescriptionService) SaveDescription(ctx context.Context, uid, description string) error {

	err := d.repo.UpdateUserDescription(ctx, uid, description)

	if err != nil {
		utils.Log(fmt.Sprintf("Failed save user description. description { %s }, uid { %s }", description, uid), err)
		return err
	}

	err = d.cache.PutUserDescription(uid, description)

	if err != nil {
		utils.Log(fmt.Sprintf("Failed save user description. description { %s }, uid { %s }", description, uid), err)
		return err
	}

	return nil

}

func (d *DefaultUserDescriptionService) GetUserDescription(ctx context.Context, uid string) (string, error) {
	desc, err := d.cache.GetUserDescription(uid)

	if err != nil {
		desc, err = d.repo.GetUserDescription(ctx, uid)
		if err != nil {
			utils.Log(fmt.Sprintf("Not found user description. uid { %s }", uid), err)
			return "", err
		}
	}

	return desc, nil
}
