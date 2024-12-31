package auth_utils

import "golang.org/x/crypto/bcrypt"

type PasswordHashUtil struct{}

func NewPasswordHashUtil() PasswordHashUtil {
	return PasswordHashUtil{}
}

func (PasswordHashUtil) HashText(text string) (string, error) {
	hashedText, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return string(hashedText), err
}

func (PasswordHashUtil) CompareHashedText(text, hashedText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(text))
	return err == nil
}
