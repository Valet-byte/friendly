package cache

import "friendly/internal/model"

type UserCache interface {
	PutUser(user model.User) error
	GetUser(uid string) (*model.User, error)
	DeleteUser(uid string) error
}
