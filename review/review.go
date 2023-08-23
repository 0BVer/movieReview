package review

import (
	"gorm.io/gorm"
	"time"
)

type review struct {
	gorm.Model
	MovieID   uint `gorm:"foreignKey:MovieID"` // foreign key
	Score     int
	Comment   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
