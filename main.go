package main

import (
	"github.com/gin-gonic/gin"

	"movieReview/config/database"
	movieDelivery "movieReview/movie/delivery"
	movieRepository "movieReview/movie/repository"
	movieUsecase "movieReview/movie/usecase"
	reviewDel "movieReview/review/delivery"
	reviewRepository "movieReview/review/repository"
	reviewUsecase "movieReview/review/usecase"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	db := database.GetMysql()

	movieRepo := movieRepository.NewRepository(db)
	movieUse := movieUsecase.NewUseCase(&movieRepo)
	movieDelivery.NewDelivery(v1, &movieUse)

	reviewRepo := reviewRepository.NewRepository(db)
	reviewUse := reviewUsecase.NewUseCase(&reviewRepo)
	reviewDel.NewDelivery(v1, &reviewUse)

	_ = r.Run(":8080")
}
