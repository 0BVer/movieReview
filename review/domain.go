package review

import (
	"gorm.io/gorm"
	"movieReview/movie"
	"time"
)

type Review struct {
	gorm.Model
	MovieID   uint `gorm:"foreignKey:MovieID"`
	Movie     movie.Movie
	Score     int
	Comment   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
