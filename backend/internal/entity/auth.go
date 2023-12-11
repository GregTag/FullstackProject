package entity

type AuthManager interface {
	MakeToken(user_id uint) (string, error)
	FetchToken(token string) (*map[string]string, error)
}
