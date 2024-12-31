package auth_contracts

type PasswordHashContract interface {
	HashText(text string) (string, error)
	CompareHashedText(text, hashedText string) bool
}
