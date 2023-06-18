package user_friends_repository

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"friendly/internal/model"
	"friendly/internal/repository/user_repository"
	"friendly/internal/utils"
)

const friendsCollection = "friends"

type FirebaseFriendsRepository struct {
	app *firebase.App
}

func NewFirebaseFriendsRepository(app *firebase.App) *FirebaseFriendsRepository {
	return &FirebaseFriendsRepository{app: app}
}

func (f *FirebaseFriendsRepository) GetFriendsByUserUid(ctx context.Context, userUid string, limit, page int) ([]model.User, error) {
	db, err := f.app.Firestore(ctx)
	defer db.Close()

	if err != nil {
		utils.Log("Failed to connect to Firestore", err)
		return []model.User{}, err
	}

	data, err := db.Collection(friendsCollection).Select("friend_id").
		Where("user_id", "==", userUid).
		Limit(limit).Offset(page * limit).Documents(ctx).Next()
	if err != nil {
		utils.Log("Failed to select friends_uid from Firestore", err)
		return []model.User{}, err
	}

	ids := new([]string)

	err = data.DataTo(&ids)
	if err != nil {
		utils.Log(fmt.Sprintf("Failde maped data from firestore. user_id: { %s }", userUid), err)
		return []model.User{}, err
	}

	users := new([]model.User)

	data, err = db.Collection(user_repository.UserCollection).Where("uid", "in", ids).Documents(ctx).Next()
	if err != nil {
		utils.Log("Failed to select users from Firestore", err)
		return []model.User{}, err
	}

	err = data.DataTo(&users)
	if err != nil {
		utils.Log(fmt.Sprintf("Failde maped data from firestore. user_ids: { %v }", ids), err)
		return []model.User{}, err
	}

	return *users, nil

}
func (f *FirebaseFriendsRepository) AddFriend(ctx context.Context, userUid, friendUid string) error {
	db, err := f.app.Firestore(ctx)
	defer db.Close()

	if err != nil {
		utils.Log("Failed to connect to Firestore", err)
		return err
	}

	data1 := model.Friends{
		FriendId: friendUid,
		UserId:   userUid,
	}
	data2 := model.Friends{
		FriendId: userUid,
		UserId:   friendUid,
	}

	err = db.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
		ref := db.Collection(friendsCollection).NewDoc()
		err := transaction.Set(ref, data1)

		if err != nil {
			return err
		}

		ref = db.Collection(friendsCollection).NewDoc()
		err = transaction.Set(ref, data2)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		utils.Log("Failed add friends data", err)
		return err
	}

	return nil
}
func (f *FirebaseFriendsRepository) DeleteFriend(ctx context.Context, userUid, friendUid string) error {
	db, err := f.app.Firestore(ctx)
	defer db.Close()

	if err != nil {
		utils.Log("Failed to connect to Firestore", err)
		return err
	}

	doc, err := db.Collection(friendsCollection).Where("user_id", "==", userUid).
		Where("friend_id", "==", friendUid).Documents(ctx).Next()

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
