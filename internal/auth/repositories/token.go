package auth_repositories

import (
	auth_entities "github.com/ladmakhi81/realtime-blogs/internal/auth/entities"
	pkg_storage "github.com/ladmakhi81/realtime-blogs/pkg/storage"
)

type TokenRepository struct {
	DBStorage pkg_storage.Storage
}

func NewTokenRepository(dbStorage pkg_storage.Storage) TokenRepository {
	return TokenRepository{DBStorage: dbStorage}
}

func (tokenRepo TokenRepository) CreateToken(token *auth_entities.Token) error {
	command := `
		INSERT INTO "_tokens"
		("access_token", "refresh_token", "user_id")
		VALUES
		($1, $2, $3)
		RETURNING "id", "created_at", "updated_at", "access_token", "refresh_token";
	`
	row := tokenRepo.DBStorage.DB.QueryRow(command, token.AccessToken, token.RefreshToken, token.User.ID)
	scanErr := row.Scan(
		&token.ID,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.AccessToken,
		&token.RefreshToken,
	)
	if scanErr != nil {
		return scanErr
	}
	return nil
}

func (tokenRepo TokenRepository) DeleteTokensByUserId() {}

func (tokenRepo TokenRepository) GetTokenByUserId() {}
