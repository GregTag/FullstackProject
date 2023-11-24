package helpers

import "golang.org/x/crypto/bcrypt"

func ComparePasswordWithHash(passwordRaw, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordRaw))
	return err
}

func GeneratePasswordHash(passwordRaw string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(passwordRaw), bcrypt.DefaultCost)
	return string(passwordHash), err
}
