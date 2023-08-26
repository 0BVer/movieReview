package delivery

import (
	"github.com/gin-gonic/gin"
	"movieReview/review/domain"
	"net/http"
)

type Delivery struct {
	useCase domain.ReviewUseCase
}

func NewDelivery(api *gin.RouterGroup, useCase domain.ReviewUseCase) {
	handler := Delivery{
		useCase: useCase,
	}

	api.POST("/reviews", handler.Create)
	api.GET("/reviews", handler.FindAllByMovieId)
}

func (d *Delivery) Create(c *gin.Context) {
	var req domain.CreateRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := d.useCase.Create(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (d *Delivery) FindAllByMovieId(c *gin.Context) {
	var params domain.FindAllParams

	err := c.BindQuery(&params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := d.useCase.FindAllByMovieId(params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
