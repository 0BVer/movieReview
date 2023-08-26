package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	gorm.Model
	Title      string `gorm:"uniqueKey"`
	Genre      string
	IsShowing  bool
	ReleasedAt time.Time
	EndAt      time.Time
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt
}

//go:generate mockery --name MovieRepository
type MovieRepository interface {
	Create(movie Movie) (Movie, error)
	FindAll(params FindAllParams) ([]Response, error)
	FindById(id string) (Movie, error)
	Update(id string, movie CreateRequest) (Movie, error)
	Delete(id string) error
	FindAllByScore(params PaginationParams) ([]ScoreRankResponse, error)
}

//go:generate mockery --name MovieUseCase
type MovieUseCase interface {
	Create(req CreateRequest) (Response, error)
	FindAll(params FindAllParams) ([]Response, error)
	FindById(id string) (Response, error)
	Update(id string, req CreateRequest) (Response, error)
	Delete(id string) error
	FindAllByScore(params PaginationParams) ([]ScoreRankResponse, error)
}

//go:generate mockery --name MovieDelivery
type MovieDelivery interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindAllByScore(c *gin.Context)
}

func (m Movie) Update(req CreateRequest) Movie {
	m.Title = req.Title
	m.Genre = req.Genre
	m.IsShowing = req.ReleasedAt.Before(time.Now()) && req.EndAt.After(time.Now())
	m.ReleasedAt = req.ReleasedAt
	m.EndAt = req.EndAt
	return m
}
