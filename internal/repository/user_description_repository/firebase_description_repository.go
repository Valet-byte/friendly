package user_description_repository

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/sirupsen/logrus"
)

const (
	descriptionParam      = "desc"
	descriptionCollection = "descriptions"
)

type FirebaseDescriptionRepository struct {
	App *firebase.App
}

func NewFirebaseDescriptionRepository(app *firebase.App) *FirebaseDescriptionRepository {
	return &FirebaseDescriptionRepository{App: app}
}

func (f *FirebaseDescriptionRepository) UpdateUserDescription(ctx context.Context, uid, description string) error {
	db, err := f.App.Firestore(ctx)
	defer db.Close()

	if err != nil {
		logrus.Error("FirebaseAppService AddDescription error: ", err)
		return err
	}

	_, err = db.Collection(descriptionCollection).Doc(uid).Set(ctx, map[string]string{descriptionParam: description})

	if err != nil {
		logrus.Error("FirebaseAppService AddDescription error: ", err)
		return err
	}

	return nil
}

func (f *FirebaseDescriptionRepository) GetUserDescription(ctx context.Context, uid string) (string, error) {
	var res = new(string)
	db, err := f.App.Firestore(ctx)
	defer db.Close()

	if err != nil {
		logrus.Error("FirebaseAppService GetUserDescription error: ", err)
		return "", err
	}

	data, err := db.Collection(descriptionCollection).Doc(uid).Get(ctx)

	*res = data.Data()[descriptionParam].(string)

	return *res, nil
}
