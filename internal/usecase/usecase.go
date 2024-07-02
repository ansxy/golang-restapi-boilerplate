package usecase

import "github.com/ansxy/niaga-catering-be/internal/repository"

type Usecase struct {
	Repo repository.IFaceRepository
}

func NewUsecase(u *Usecase) *Usecase {
	return &Usecase{
		Repo: u.Repo,
	}
}
