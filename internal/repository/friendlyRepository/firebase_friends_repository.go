package repository

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"friendly/internal/model"
	"friendly/internal/utils"
)

const friendsCollection = "friends"

type FirebaseFriendsRepository struct {
	fb *firebase.App
}

func NewFirebaseFriendsRepository(fb *firebase.App) *FirebaseFriendsRepository {
	return &FirebaseFriendsRepository{fb: fb}
}

func (f *FirebaseFriendsRepository) GetFriendsByUserUid(ctx context.Context, userUid string, limit, page int) (ids []string, err error) {
	db, err := f.fb.Firestore(ctx)

	if err != nil {
		utils.Log("Failed to connect to Firestore", err)
		return
	}

	data, err := db.Collection(friendsCollection).Select("friend_id").Where("user_id", "==", userUid).Limit(limit).Offset(page * limit).Documents(ctx).Next()
	if err != nil {
		utils.Log("Failed to select friends_uid from Firestore", err)
		return
	}

	err = data.DataTo(&ids)
	if err != nil {
		utils.Log(fmt.Sprintf("Failde maped data from firestore. user_id: { %s }", userUid), err)
		return
	}

	return
}

func (f *FirebaseFriendsRepository) AddFriendRequest(ctx context.Context, userUid, friendUid string) (model.Friends, error) {
	db, err := f.fb.Firestore(ctx)

	if err != nil {
		utils.Log("Failed to connect to Firestore", err)
		return model.Friends{}, err
	}

	data := model.Friends{
		FriendId:  friendUid,
		UserId:    userUid,
		Confirmed: false,
	}

	add, _, err := db.Collection(friendsCollection).Add(ctx, data)
	if err != nil {
		utils.Log(fmt.Sprintf("Failed to add friend data to Firestore. user_id: { %s } friend_id: { %s }", userUid, friendUid), err)
		return model.Friends{}, err
	}

	data.Uid = add.ID
	return data, nil
}

func (f *FirebaseFriendsRepository) AddFriend(ctx context.Context, userUid, friendUid string) (model.Friends, error) {
	db, err := f.fb.Firestore(ctx)

	if err != nil {
		utils.Log("Failed to connect to Firestore", err)
		return model.Friends{}, err
	}

	data := model.Friends{
		FriendId:  friendUid,
		UserId:    userUid,
		Confirmed: true,
	}

	add, _, err := db.Collection(friendsCollection).Add(ctx, data)
	if err != nil {
		utils.Log(fmt.Sprintf("Failed to add friend data to Firestore. user_id: { %s } friend_id: { %s }", userUid, friendUid), err)
		return model.Friends{}, err
	}

	data.Uid = add.ID
	return data, nil
}

func (f *FirebaseFriendsRepository) DeleteFriend(ctx context.Context, userUid, friendUid string) error {
	db, err := f.fb.Firestore(ctx)

	if err != nil {
		utils.Log("Failed to connect to Firestore", err)
		return err
	}

	doc, err := db.Collection(friendsCollection).Where("user_id", "==", userUid).Where("friend_id", "==", friendUid).Documents(ctx).Next()

	if err != nil {
		utils.Log("Failed to select friends data", err)
		return err
	}

	_, err = doc.Ref.Delete(ctx)
	if err != nil {
		utils.Log(fmt.Sprintf("Failed to delete friends. user_id: { %s } friend_id: { %s }", userUid, friendUid), err)
		return err
	}

	return nil
}
