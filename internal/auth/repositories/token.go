package auth_repositories

type TokenRepository struct{}

func (tokenRepo TokenRepository) CreateToken() {}

func (tokenRepo TokenRepository) DeleteTokensByUserId() {}

func (tokenRepo TokenRepository) GetTokenByUserId() {}
