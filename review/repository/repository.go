package repository

import (
	"gorm.io/gorm"
	"movieReview/review/domain"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	err := db.AutoMigrate(&domain.Review{})
	if err != nil {
		panic(err)
	}
	return Repository{db: db}
}

func (r *Repository) Create(m domain.Review) (domain.Review, error) {
	err := r.db.Create(&m).Error
	if err != nil {
		return domain.Review{}, err
	}

	return m, nil
}

func (r *Repository) FindAllByMovieId(params domain.FindAllParams) ([]domain.Response, error) {
	var res []domain.Response

	err := r.db.Raw(
		"SELECT * "+
			"FROM reviews "+
			"WHERE (@MovieID = '' OR movie_id = @MovieID) "+
			"AND (@ScoreCap = '' OR score >= @ScoreCap) "+
			"ORDER BY created_at DESC",
		params,
	).Scan(&res).Error
	if err != nil {
		return []domain.Response{}, err
	}

	return res, nil
}
