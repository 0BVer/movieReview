package main

import (
	"github.com/gin-gonic/gin"
	"movieReview/movie"
	"movieReview/review"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	movie.Config(v1)
	review.Config(v1)
	_ = r.Run(":8080")
}
