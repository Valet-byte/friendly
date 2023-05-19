package service

import (
	"encoding/json"
	"errors"
	"friendly/internal/model"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	userPrefix        = "user_"
	descriptionPrefix = "description_"
	ipPrefix          = "ip_"
)

type RedisService struct {
	client    *redis.Client
	validTime time.Duration
}

func NewRedisService(client *redis.Client, validTime time.Duration) *RedisService {
	return &RedisService{client: client, validTime: validTime}
}

func (r *RedisService) PutUser(uid string, user *model.User) (string, error) {
	data, err := json.Marshal(user)
	if err != nil {
		logrus.Error("Failed convert value :", err)
		return "", err
	}

	_, err = r.client.Set(userPrefix+uid, data, r.validTime).Result()

	if err != nil {
		logrus.Error("RedisService set error :", err)
		return "", err
	}

	return uid, nil
}
func (r *RedisService) GetUser(uid string) (*model.User, error) {
	userData, err := r.client.Get(userPrefix + uid).Bytes()

	var retrievedUser model.User

	if err == redis.Nil {
		return nil, errors.New("not exist user")
	}
	if err != nil {
		logrus.Error("RedisService set error :", err)
		return &retrievedUser, err
	}
	logrus.Info(string(userData[:]))
	err = json.Unmarshal(userData, &retrievedUser)

	if err != nil {
		logrus.Error("RedisService set error :", err)
		return &retrievedUser, err
	}

	return &retrievedUser, nil
}
func (r *RedisService) IncrementRequestCount(ip string) error {
	err := r.client.Incr(ipPrefix + ip).Err()
	if err != nil {
		return err
	}

	err = r.client.Expire(ipPrefix+ip, 1*time.Minute).Err()
	return err
}
func (r *RedisService) GetRequestCount(ip string) (int, error) {
	count, err := r.client.Get(ipPrefix + ip).Int()
	if err != nil {
		return 0, err
	}

	return count, nil
}
func (r *RedisService) DeleteUser(uid string) error {
	return r.client.Del(userPrefix + uid).Err()
}
func (r *RedisService) PutUserDescription(uid, description string) (string, error) {
	_, err := r.client.Set(descriptionPrefix+uid, description, r.validTime*5).Result()
	if err != nil {
		logrus.Error("RedisService set error :", err)
		return "", err
	}

	return uid, nil
}
func (r *RedisService) GetUserDescription(uid string) (string, error) {
	desc, err := r.client.Get(descriptionPrefix + uid).Bytes()
	if err == redis.Nil {
		return "", errors.New("not exist user")
	}
	if err != nil {
		logrus.Error("RedisService set error :", err)
		return "", err
	}
	return string(desc[:]), nil
}
func (r *RedisService) DeleteUserDescription(uid string) error {
	return r.client.Del(descriptionPrefix + uid).Err()
}
