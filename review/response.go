package review

import (
	"time"
)

type response struct {
	ID        uint      `json:"id"`
	MovieID   uint      `json:"movieId"`
	Score     int       `json:"score"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m Review) fromEntity() response {
	return response{
		ID:        m.ID,
		MovieID:   m.MovieID,
		Score:     m.Score,
		Comment:   m.Comment,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
