package user_description_cache

type UserDescriptionCache interface {
	PutUserDescription(uid, description string) error
	GetUserDescription(uid string) (string, error)
}
