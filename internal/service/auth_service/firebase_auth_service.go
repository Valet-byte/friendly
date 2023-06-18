package auth_service

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/sirupsen/logrus"
)

type FirebaseAuthService struct {
	App *firebase.App
}

func NewFirebaseAuthService(app *firebase.App) *FirebaseAuthService {
	return &FirebaseAuthService{App: app}
}

func (f *FirebaseAuthService) GetUserUid(ctx context.Context, token string) (string, error) {
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
