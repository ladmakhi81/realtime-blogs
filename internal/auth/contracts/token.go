package auth_contracts

type TokenRepositoryContract interface {
	CreateToken()
	DeleteTokensByUserId()
	GetTokenByUserId()
}
