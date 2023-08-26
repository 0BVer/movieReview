package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"movieReview/config/database"
	_ "movieReview/docs"
	movieDelivery "movieReview/movie/delivery"
	movieRepository "movieReview/movie/repository"
	movieUsecase "movieReview/movie/usecase"
	reviewDel "movieReview/review/delivery"
	reviewRepository "movieReview/review/repository"
	reviewUsecase "movieReview/review/usecase"
)

// @title Movie Review API
// @description This is a Movie Review API server.
// @version 1
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	db := database.GetMysql()

	movieRepo := movieRepository.NewRepository(db)
	movieUse := movieUsecase.NewUseCase(&movieRepo)
	movieDelivery.NewDelivery(v1, &movieUse)

	reviewRepo := reviewRepository.NewRepository(db)
	reviewUse := reviewUsecase.NewUseCase(&reviewRepo)
	reviewDel.NewDelivery(v1, &reviewUse)

	_ = r.Run(":8080")
}
