package movie

import (
	"github.com/gin-gonic/gin"
	"log"
	"movieReview/config/database"
	"net/http"
	"time"
)

var db = database.GetDB()

func Config(api *gin.RouterGroup) {
	_ = db.DB.AutoMigrate(&Movie{})
	api.GET("/movies", findAll)
	api.POST("/movies", create)
	api.GET("/movies/:id", findById)
	api.PUT("/movies/:id", update)
	api.DELETE("/movies/:id", remove)
	api.GET("/movies/score", findAllByScore)
}

func create(c *gin.Context) {
	var req createRequest

	shouldBindErr := c.ShouldBind(&req)
	if shouldBindErr != nil {
		log.Println(shouldBindErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": shouldBindErrMsg + " " + shouldBindErr.Error()})
		return
	}

	m := req.toEntity()

	dbCreateErr := db.DB.Create(&m).Error
	if dbCreateErr != nil {
		log.Println(dbCreateErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": dbCreateErrMsg + " " + dbCreateErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, m.fromEntity())
}

func findAll(c *gin.Context) {
	var res []response
	var params findAllParams

	bindQueryErr := c.BindQuery(&params)
	if bindQueryErr != nil {
		log.Println(bindQueryErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": bindQueryErrMsg + " " + bindQueryErr.Error()})
		return
	}

	res = []response{}

	dbFindAllErr := db.DB.Raw(
		"SELECT * "+
			"FROM movies "+
			"WHERE (@Genre = '' OR genre LIKE @Genre)"+
			"AND (@IsShowing = '' OR is_showing = @IsShowing)"+
			"ORDER BY released_at",
		params,
	).Scan(&res).Error
	if dbFindAllErr != nil {
		log.Println(dbFindAllErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": dbFindAllErrMsg + " " + dbFindAllErr.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func findById(c *gin.Context) {
	id := c.Param("id")
	var m Movie

	dbWhereErr := db.DB.Where("id = ?", id).First(&m).Error
	if dbWhereErr != nil {
		log.Println(dbWhereErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": dbWhereErrMsg + " " + dbWhereErr.Error()})
		return
	}
	c.JSON(http.StatusOK, m.fromEntity())
}

func update(c *gin.Context) {
	id := c.Param("id")
	var req createRequest

	shouldBindErr := c.ShouldBind(&req)
	if shouldBindErr != nil {
		log.Println(shouldBindErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": shouldBindErrMsg + " " + shouldBindErr.Error()})
		return
	}

	var m Movie

	dbWhereErr := db.DB.Where("id = ?", id).First(&m).Error
	if dbWhereErr != nil {
		log.Println(dbWhereErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": dbWhereErrMsg + " " + dbWhereErr.Error()})
		return
	}

	m.Title = req.Title
	m.Genre = req.Genre
	m.IsShowing = req.ReleasedAt.Before(time.Now()) && req.EndAt.After(time.Now())
	m.ReleasedAt = req.ReleasedAt
	m.EndAt = req.EndAt

	dbSaveErr := db.DB.Save(&m).Error
	if dbSaveErr != nil {
		log.Println(dbSaveErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": dbSaveErrMsg + " " + dbSaveErr.Error()})
		return
	}

	c.JSON(http.StatusOK, m.fromEntity())
}

func remove(c *gin.Context) {
	id := c.Param("id")
	var m Movie

	dbWhereErr := db.DB.Where("id = ?", id).First(&m).Error
	if dbWhereErr != nil {
		log.Println(dbWhereErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": dbWhereErrMsg + " " + dbWhereErr.Error()})
		return
	}

	dbDeleteErr := db.DB.Delete(&m).Error
	if dbDeleteErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": dbDeleteErrMsg + " " + dbDeleteErr.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// TODO: Implement this function
func findAllByScore(c *gin.Context) {
	var res []scoreRankResponse

	var params paginationParams

	bindQueryErr := c.BindQuery(&params)
	if bindQueryErr != nil {
		log.Println(bindQueryErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": bindQueryErrMsg + " " + bindQueryErr.Error()})
		return
	}

	dbFindAllErr := db.DB.Raw(
		"SELECT movies.*, AVG(reviews.score) as ScoreAvg "+
			"FROM movies "+
			"LEFT JOIN reviews ON movies.id = reviews.movie_id "+
			"GROUP BY movies.id "+
			"ORDER BY ScoreAvg DESC "+
			"LIMIT ? OFFSET ?",
		params.Size, params.Page*params.Size,
	).Scan(&res).Error
	if dbFindAllErr != nil {
		log.Println(dbFindAllErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": dbFindAllErrMsg + " " + dbFindAllErr.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
