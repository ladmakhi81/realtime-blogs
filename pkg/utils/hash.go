package pkg_utils

import "golang.org/x/crypto/bcrypt"

func HashText(text string) (string, error) {
	hashedText, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return string(hashedText), err
}

func CompareHashedText(text, hashedText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(text))
	return err == nil
}
