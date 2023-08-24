package review

type createRequest struct {
	MovieID uint   `json:"movieID" binding:"required"`
	Score   int    `json:"score" binding:"required,min=0,max=5"`
	Comment string `json:"comment" binding:"required"`
}

func (r createRequest) toEntity() Review {
	return Review{
		MovieID: r.MovieID,
		Score:   r.Score,
		Comment: r.Comment,
	}
}

type findAllParams struct {
	MovieID  uint `form:"movieId" binding:"omitempty"`
	ScoreCap int  `form:"scoreCap" binding:"omitempty,min=0,max=5"`
}
