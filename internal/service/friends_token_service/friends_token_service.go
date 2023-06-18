package friends_token_service

type FriendsTokenService interface {
	GetFriendsToken(uid string) (string, error)
	ParseFriendsToken(token string) (string, error)
}
