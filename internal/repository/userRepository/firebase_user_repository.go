package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"friendly/internal/model"
	"friendly/internal/utils"
)

const userCollection = "users"

type FirebaseUserRepository struct {
	fb *firebase.App
}

func NewFirebaseUserRepository(fb *firebase.App) *FirebaseUserRepository {
	return &FirebaseUserRepository{fb: fb}
}

func (f *FirebaseUserRepository) SaveUser(ctx context.Context, user model.User) (model.User, error) {
	db, err := f.fb.Firestore(ctx)

	if err != nil {
		utils.Log("Failed to connect to Firestore", err)
		return user, err
	}

	if user.Uid == "" {
		_, err = db.Collection(userCollection).Doc(user.Uid).Set(ctx, user)

		if err != nil {
			utils.Log(fmt.Sprintf("Failed to update user. user: { %v }", user), err)
			return user, err
		}

		return user, nil
	} else {
		doc, _, err := db.Collection(userCollection).Add(ctx, user)

		if err != nil {
			utils.Log(fmt.Sprintf("Failed to save userdata to Firestore. user: { %v }", user), err)
			return user, err
		}
		user.Uid = doc.ID
		return user, nil
	}
}

func (f *FirebaseUserRepository) GetUserByUid(ctx context.Context, uid string) (model.User, error) {
	db, err := f.fb.Firestore(ctx)

	if err != nil {
		utils.Log("Failed to connect to Firestore", err)
		return model.User{}, err
	}

	var retrievedUser = new(model.User)
	data, err := db.Collection(userCollection).Doc(uid).Get(ctx)

	err = data.DataTo(&retrievedUser)

	if err != nil {
		utils.Log("Failed to parse user data", err)
		return *retrievedUser, err
	}

	return *retrievedUser, nil
}

func (f *FirebaseUserRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	db, err := f.fb.Firestore(ctx)

	if err != nil {
		utils.Log("Failed to connect to Firestore", err)
		return model.User{}, err
	}

	var retrievedUser = new(model.User)
	data, err := db.Collection(userCollection).Where("emil", "==", email).Documents(ctx).Next()

	if err != nil {
		utils.Log(fmt.Sprintf("Failed to select user details by email: { %s }", email), err)
		return model.User{}, err
	}

	err = data.DataTo(&retrievedUser)

	if err != nil {
		utils.Log("Failed parse user data", err)
		return model.User{}, err
	}

	return *retrievedUser, nil

}

func (f *FirebaseUserRepository) DeleteUser(ctx context.Context, uid string) error {
	db, err := f.fb.Firestore(ctx)

	if err != nil {
		utils.Log("Failed connection to firestore", err)
		return err
	}

	_, err = db.Collection(userCollection).Doc(uid).Update(ctx, []firestore.Update{
		{Path: "deleted", Value: true},
	})
	if err != nil {
		utils.Log("Failed update user data", err)
		return err
	}

	return nil
}
