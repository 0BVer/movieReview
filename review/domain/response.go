package domain

import (
	"time"
)

type Response struct {
	ID        uint      `json:"id"`
	MovieID   uint      `json:"movieId"`
	Score     int       `json:"score"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m Review) FromEntity() Response {
	return Response{
		ID:        m.ID,
		MovieID:   m.MovieID,
		Score:     m.Score,
		Comment:   m.Comment,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
