package entity

type AuthManager interface {
	MakeToken(userID uint) (string, error)
	FetchToken(token string) (*map[string]string, error)
}
