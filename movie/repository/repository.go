package repository

import (
	"gorm.io/gorm"
	"movieReview/movie/domain"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	err := db.AutoMigrate(&domain.Movie{})
	if err != nil {
		panic(err)
	}
	return Repository{db: db}
}

func (r *Repository) Create(m domain.Movie) (domain.Movie, error) {
	err := r.db.Create(&m).Error
	if err != nil {
		return domain.Movie{}, err
	}

	return m, nil
}

func (r *Repository) FindAll(params domain.FindAllParams) ([]domain.Response, error) {
	var res []domain.Response
	err := r.db.Raw(
		"SELECT * "+
			"FROM movies "+
			"WHERE (@Genre = '' OR genre LIKE @Genre) "+
			"AND (@IsShowing = '' OR is_showing = @IsShowing) "+
			"AND deleted_at IS NULL "+
			"ORDER BY released_at",
		params,
	).Scan(&res).Error
	if err != nil {
		return []domain.Response{}, err
	}

	return res, nil
}

func (r *Repository) FindById(id string) (domain.Movie, error) {
	var m domain.Movie
	err := r.db.Where("id = ?", id).First(&m).Error
	if err != nil {
		return domain.Movie{}, err
	}

	return m, nil
}

func (r *Repository) Update(id string, req domain.CreateRequest) (domain.Movie, error) {
	m, err := r.FindById(id)
	if err != nil {
		return domain.Movie{}, err
	}

	m = m.Update(req)

	err = r.db.Save(&m).Error
	if err != nil {
		return domain.Movie{}, err
	}

	return m, nil
}

func (r *Repository) Delete(id string) error {
	m, err := r.FindById(id)
	if err != nil {
		return err
	}

	err = r.db.Delete(&m).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindAllByScore(params domain.PaginationParams) ([]domain.ScoreRankResponse, error) {
	var res []domain.ScoreRankResponse
	err := r.db.Raw(
		"SELECT movies.*, AVG(reviews.score) as ScoreAvg "+
			"FROM movies "+
			"LEFT JOIN reviews ON movies.id = reviews.movie_id "+
			"WHERE movies.deleted_at IS NULL "+
			"GROUP BY movies.id "+
			"ORDER BY ScoreAvg DESC "+
			"LIMIT ? OFFSET ?",
		params.Size, params.Page*params.Size,
	).Scan(&res).Error
	if err != nil {
		return []domain.ScoreRankResponse{}, err
	}

	return res, nil
}
