package delivery

import (
	"github.com/gin-gonic/gin"
	"movieReview/movie/domain"
	"net/http"
)

type Delivery struct {
	useCase domain.MovieUseCase
}

func NewDelivery(api *gin.RouterGroup, useCase domain.MovieUseCase) {
	handler := Delivery{
		useCase: useCase,
	}

	api.GET("/movies", handler.FindAll)
	api.POST("/movies", handler.Create)
	api.GET("/movies/:id", handler.FindById)
	api.PUT("/movies/:id", handler.Update)
	api.DELETE("/movies/:id", handler.Delete)
	api.GET("/movies/score", handler.FindAllByScore)
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

func (d *Delivery) FindAll(c *gin.Context) {
	var params domain.FindAllParams

	err := c.BindQuery(&params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := d.useCase.FindAll(params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (d *Delivery) FindById(c *gin.Context) {
	id := c.Param("id")

	res, err := d.useCase.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (d *Delivery) Update(c *gin.Context) {
	id := c.Param("id")
	var req domain.CreateRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := d.useCase.Update(id, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (d *Delivery) Delete(c *gin.Context) {
	id := c.Param("id")

	err := d.useCase.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (d *Delivery) FindAllByScore(c *gin.Context) {
	var params domain.PaginationParams

	err := c.BindQuery(&params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := d.useCase.FindAllByScore(params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
