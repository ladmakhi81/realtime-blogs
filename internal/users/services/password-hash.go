package users_services

import "golang.org/x/crypto/bcrypt"

type PasswordHashService struct{}

func NewPasswordHashService() PasswordHashService {
	return PasswordHashService{}
}

func (PasswordHashService) HashText(text string) (string, error) {
	hashedText, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return string(hashedText), err
}

func (PasswordHashService) CompareHashedText(text, hashedText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(text))
	return err == nil
}
