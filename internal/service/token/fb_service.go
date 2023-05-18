package service

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"friendly/internal/model"
	"github.com/sirupsen/logrus"
)

const (
	userCollection        = "users"
	descriptionCollection = "descriptions"
	descriptionParam      = "desc"
)

type FirebaseAppService struct {
	App *firebase.App
}

func NewFirebaseApp(app *firebase.App) *FirebaseAppService {
	return &FirebaseAppService{App: app}
}

func (f *FirebaseAppService) CheckToken(ctx context.Context, token string) (bool, error) {
	client, err := f.App.Auth(ctx)

	if err != nil {
		logrus.Error("FirebaseAppService CheckToken error: ", err)
		return false, err
	}

	_, err = client.VerifyIDToken(ctx, token)

	if err != nil {
		logrus.Error("FirebaseAppService CheckToken error: ", err)
		return false, err
	}

	return true, nil
}

func (f *FirebaseAppService) GetTokenWithClaims(ctx context.Context, uid string, claims map[string]interface{}) (string, error) {
	client, err := f.App.Auth(ctx)

	if err != nil {
		logrus.Error("FirebaseAppService GetToken error: ", err)
		return "", err
	}

	token, err := client.CustomTokenWithClaims(ctx, uid, claims)
	if err != nil {
		logrus.Error("FirebaseAppService GetToken error: ", err)
		return "", err
	}

	return token, nil
}

func (f *FirebaseAppService) GetToken(ctx context.Context, uid string) (string, error) {
	client, err := f.App.Auth(ctx)

	if err != nil {
		logrus.Error("FirebaseAppService GetToken error: ", err)
		return "", err
	}

	token, err := client.CustomToken(ctx, uid)
	if err != nil {
		logrus.Error("FirebaseAppService GetToken error: ", err)
		return "", err
	}

	return token, nil
}

func (f *FirebaseAppService) GetUserUid(ctx context.Context, token string) (string, error) {
	client, err := f.App.Auth(ctx)

	if err != nil {
		logrus.Error("FirebaseAppService GetKUserUid error: ", err)
		return "", err
	}

	_, err = client.VerifyIDToken(ctx, token)

	if err != nil {
		logrus.Error("FirebaseAppService GetKUserUid error: ", err)
		return "", err
	}

	fbToken, err := client.VerifyIDToken(ctx, token)

	if err != nil {
		logrus.Error("FirebaseAppService GetKUserUid error: ", err)
		return "", err
	}

	return fbToken.UID, nil
}

func (f *FirebaseAppService) AddUserData(ctx context.Context, uid string, user model.User) (bool, error) {
	db, err := f.App.Firestore(ctx)

	if err != nil {
		logrus.Error("FirebaseAppService AddUserData error: ", err)
		return false, err
	}

	_, err = db.Collection(userCollection).Doc(uid).Set(ctx, user)

	if err != nil {
		logrus.Error("FirebaseAppService AddUserData error: ", err)
		return false, err
	}

	return true, nil
}

func (f *FirebaseAppService) RegistrationUser(ctx context.Context, data model.RegistrationUserData) (string, error) {
	params := (&auth.UserToCreate{}).
		Email(data.Email).
		Password(data.Password).
		EmailVerified(false).
		DisplayName(data.UserName).
		Disabled(false)

	client, err := f.App.Auth(ctx)

	if err != nil {
		logrus.Error("FirebaseAppService registrationUser error (create client): ", err)
		return "", err
	}

	record, err := client.CreateUser(ctx, params)

	if err != nil {
		logrus.Error("FirebaseAppService registrationUser error (user registration): ", err)
		return "", err
	}

	return record.UID, nil
}

func (f *FirebaseAppService) GetUserData(ctx context.Context, uid string) (*model.User, error) {

	db, err := f.App.Firestore(ctx)
	var retrievedUser = new(model.User)

	if err != nil {
		logrus.Error("FirebaseAppService GetUserData error: ", err)
		return retrievedUser, err
	}

	data, err := db.Collection(userCollection).Doc(uid).Get(ctx)

	if err != nil {
		logrus.Error("FirebaseAppService GetUserData error: ", err)
		return retrievedUser, err
	}

	err = data.DataTo(&retrievedUser)

	if err != nil {
		logrus.Error("FirebaseAppService GetUserData error: ", err)
		return retrievedUser, err
	}

	retrievedUser.Uid = uid

	return retrievedUser, nil
}
func (f *FirebaseAppService) DeleteUser(ctx context.Context, uid string) error {

	client, err := f.App.Auth(ctx)
	if err != nil {
		return err
	}

	err = client.DeleteUser(ctx, uid)
	if err != nil {
		logrus.Error("failed delete user! error: ", err)
		return err
	}

	db, err := f.App.Firestore(ctx)

	if err != nil {
		logrus.Error("FirebaseAppService deleteUser error: ", err)
		return err
	}

	_, err = db.Collection(userCollection).Doc(uid).Delete(ctx)
	if err != nil {
		logrus.Error("FirebaseAppService deleteUser. Failed delete user data! error: ", err)
		return err
	}
	_, err = db.Collection(descriptionCollection).Doc(uid).Delete(ctx)

	if err != nil {
		logrus.Error("FirebaseAppService deleteUser. Failed delete user description! error: ", err)
		return err
	}

	return nil
}
func (f *FirebaseAppService) AddDescription(ctx context.Context, uid, description string) (bool, error) {
	db, err := f.App.Firestore(ctx)

	if err != nil {
		logrus.Error("FirebaseAppService AddDescription error: ", err)
		return false, err
	}

	_, err = db.Collection(descriptionCollection).Doc(uid).Set(ctx, map[string]string{descriptionParam: description})

	if err != nil {
		logrus.Error("FirebaseAppService AddDescription error: ", err)
		return false, err
	}

	return true, nil
}
func (f *FirebaseAppService) GetUserDescription(ctx context.Context, uid string) (string, error) {
	var res = new(string)
	db, err := f.App.Firestore(ctx)

	if err != nil {
		logrus.Error("FirebaseAppService GetUserDescription error: ", err)
		return "", err
	}

	data, err := db.Collection(descriptionCollection).Doc(uid).Get(ctx)

	*res = data.Data()[descriptionParam].(string)

	return *res, nil
}
