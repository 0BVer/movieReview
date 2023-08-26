package usecase

import (
	"movieReview/movie/domain"
)

type UseCase struct {
	repo domain.MovieRepository
}

func NewUseCase(repo domain.MovieRepository) UseCase {
	return UseCase{repo: repo}
}

func (u *UseCase) Create(req domain.CreateRequest) (domain.Response, error) {
	movie, err := u.repo.Create(req.ToEntity())
	if err != nil {
		return domain.Response{}, err
	}

	return movie.FromEntity(), nil
}

func (u *UseCase) FindAll(params domain.FindAllParams) ([]domain.Response, error) {
	res, err := u.repo.FindAll(params)
	if err != nil {
		return []domain.Response{}, err
	}

	return res, nil
}

func (u *UseCase) FindById(id string) (domain.Response, error) {
	movie, err := u.repo.FindById(id)
	if err != nil {
		return domain.Response{}, err
	}

	return movie.FromEntity(), nil
}

func (u *UseCase) Update(id string, req domain.CreateRequest) (domain.Response, error) {
	res, err := u.repo.Update(id, req)
	if err != nil {
		return domain.Response{}, err
	}

	return res.FromEntity(), nil
}

func (u *UseCase) Delete(id string) error {
	err := u.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UseCase) FindAllByScore(params domain.PaginationParams) ([]domain.ScoreRankResponse, error) {
	res, err := u.repo.FindAllByScore(params)
	if err != nil {
		return []domain.ScoreRankResponse{}, err
	}

	return res, nil
}
