package auth

type AuthUsecase struct {
	AuthRepository AuthRepository
}

func NewAuthUsecase(ar AuthRepository) *AuthUsecase {
	return &AuthUsecase{
		AuthRepository: ar,
	}
}
