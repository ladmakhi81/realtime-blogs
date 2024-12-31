package auth_contracts

type PasswordHashServiceContract interface {
	HashText(text string) (string, error)
	CompareHashedText(text, hashedText string) bool
}
