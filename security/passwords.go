package security

import (
	"golang.org/x/crypto/bcrypt"
)

func GetPasswordHash(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func CheckPasswordHash(hash []byte, pwd string) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(pwd))
	return err == nil
}
