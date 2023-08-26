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

// Create @Summary Create a movie
// @Description Create a movie
// @Tags movies
// @Accept json
// @Produce json
// @Param movie body domain.CreateRequest true "Movie"
// @Success 201 {object} domain.Response
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /movies [post]
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

// FindAll @Summary Find all movies
// @Description Find all movies
// @Tags movies
// @Produce json
// @Param genre query string false "Genre"
// @Param isShowing query string false "IsShowing"
// @Success 200 {array} domain.Response
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /movies [get]
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

// FindById @Summary Find a movie by id
// @Description Find a movie by id
// @Tags movies
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} domain.Response
// @Failure 404 {object} map[string]any
// @Router /movies/{id} [get]
func (d *Delivery) FindById(c *gin.Context) {
	id := c.Param("id")

	res, err := d.useCase.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Update @Summary Update a movie
// @Description Update a movie
// @Tags movies
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param movie body domain.CreateRequest true "Movie"
// @Success 200 {object} domain.Response
// @Failure 400 {object} map[string]any
// @Failure 404 {object} map[string]any
// @Router /movies/{id} [put]
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

// Delete @Summary Delete a movie
// @Description Delete a movie
// @Tags movies
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 204
// @Failure 404 {object} map[string]any
// @Router /movies/{id} [delete]
func (d *Delivery) Delete(c *gin.Context) {
	id := c.Param("id")

	err := d.useCase.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// FindAllByScore @Summary Find all movies by score
// @Description Find all movies by score
// @Tags movies
// @Produce json
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Success 200 {array} domain.ScoreRankResponse
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /movies/score [get]
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
