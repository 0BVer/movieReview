package movie

import "time"

type createRequest struct {
	Title      string    `form:"title" binding:"required"`
	Genre      string    `form:"genre" binding:"required"`
	ReleasedAt time.Time `form:"releasedAt" binding:"required"`
	EndAt      time.Time `form:"endAt" binding:"required"`
}

func (r createRequest) toEntity() Movie {
	return Movie{
		Title:      r.Title,
		Genre:      r.Genre,
		IsShowing:  r.ReleasedAt.Before(time.Now()) && r.EndAt.After(time.Now()),
		ReleasedAt: r.ReleasedAt,
		EndAt:      r.EndAt,
	}
}

type findAllParams struct {
	Genre     string `form:"genre" binding:"omitempty"`               // 원하는 like 조건에 따라 %를 붙여서 사용
	IsShowing string `form:"isShowing" binding:"omitempty,oneof=0 1"` // 0: false, 1: true
}

type paginationParams struct {
	Page int `form:"page" binding:"omitempty,min=0"`
	Size int `form:"size" binding:"omitempty,min=1,max=100"`
}
