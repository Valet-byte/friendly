package friendly_token_service

type FriendlyTokenService interface {
	ParseToken(tokenString string) (string, int64, error)
	CreateToken(userId string) (string, error)
}
