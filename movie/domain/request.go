package domain

import (
	"time"
)

type CreateRequest struct {
	Title      string    `form:"title" binding:"required"`
	Genre      string    `form:"genre" binding:"required"`
	ReleasedAt time.Time `form:"releasedAt" binding:"required"`
	EndAt      time.Time `form:"endAt" binding:"required"`
}

func (r CreateRequest) ToEntity() Movie {
	return Movie{
		Title:      r.Title,
		Genre:      r.Genre,
		IsShowing:  r.ReleasedAt.Before(time.Now()) && r.EndAt.After(time.Now()),
		ReleasedAt: r.ReleasedAt,
		EndAt:      r.EndAt,
	}
}

type FindAllParams struct {
	Genre     string `form:"genre" binding:"omitempty"`               // 원하는 like 조건에 따라 %를 붙여서 사용
	IsShowing string `form:"isShowing" binding:"omitempty,oneof=0 1"` // 0: false, 1: true
}

type PaginationParams struct {
	Page int `form:"page" binding:"required,min=0"`
	Size int `form:"size" binding:"required,min=1,max=100"`
}
