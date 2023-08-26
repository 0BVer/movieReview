package usecase

import "movieReview/review/domain"

type UseCase struct {
	repo domain.ReviewRepository
}

func NewUseCase(repo domain.ReviewRepository) UseCase {
	return UseCase{repo: repo}
}

func (u *UseCase) Create(req domain.CreateRequest) (domain.Response, error) {
	movie, err := u.repo.Create(req.ToEntity())
	if err != nil {
		return domain.Response{}, err
	}

	return movie.FromEntity(), nil
}

func (u *UseCase) FindAllByMovieId(params domain.FindAllParams) ([]domain.Response, error) {
	res, err := u.repo.FindAllByMovieId(params)
	if err != nil {
		return []domain.Response{}, err
	}

	return res, nil
}
