package domain

import (
	"time"
)

type Response struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Genre      string    `json:"genre"`
	IsShowing  bool      `json:"isShowing"`
	ReleasedAt time.Time `json:"releasedAt"`
	EndAt      time.Time `json:"endAt"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (m Movie) FromEntity() Response {
	return Response{
		ID:         m.ID,
		Title:      m.Title,
		Genre:      m.Genre,
		IsShowing:  m.IsShowing,
		ReleasedAt: m.ReleasedAt,
		EndAt:      m.EndAt,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

type ScoreRankResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Genre      string    `json:"genre"`
	ScoreAvg   float64   `json:"scoreAvg"`
	IsShowing  bool      `json:"isShowing"`
	ReleasedAt time.Time `json:"releasedAt"`
	EndAt      time.Time `json:"endAt"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
