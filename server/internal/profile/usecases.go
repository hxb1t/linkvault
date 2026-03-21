package profile

type ProfileUsecase struct {
	ProfileRepository ProfileRepository
}

func NewProfileUsecase(pr ProfileRepository) *ProfileUsecase {
	return &ProfileUsecase{
		ProfileRepository: pr,
	}
}
