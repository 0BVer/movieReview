package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movieReview/movie/domain"
	"time"
)

type Review struct {
	gorm.Model
	MovieID   uint `gorm:"foreignKey:MovieID"`
	Movie     domain.Movie
	Score     int
	Comment   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

//go:generate mockery --name ReviewRepository
type ReviewRepository interface {
	Create(m Review) (Review, error)
	FindAllByMovieId(params FindAllParams) ([]Response, error)
}

//go:generate mockery --name ReviewUseCase
type ReviewUseCase interface {
	Create(req CreateRequest) (Response, error)
	FindAllByMovieId(params FindAllParams) ([]Response, error)
}

//go:generate mockery --name ReviewDelivery
type ReviewDelivery interface {
	Create(c *gin.Context)
	FindAllByMovieId(c *gin.Context)
}
