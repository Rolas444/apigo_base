package repository

type ValidTokenRepository interface {
	FindByToken(token string) (bool, error)
	Create(token string) error
	Delete(token string) error
	IsValidToken(token string) (bool, error)
}
