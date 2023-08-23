package movie

import (
	"gorm.io/gorm"
	"time"
)

type movie struct {
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
