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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": shouldBindErrMsg})
		return
	}

	m := req.toEntity()

	dbCreateErr := db.DB.Create(&m).Error
	if dbCreateErr != nil {
		log.Println(dbCreateErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": dbCreateErrMsg})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": bindQueryErrMsg})
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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": dbFindAllErrMsg})
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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": dbWhereErrMsg})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": shouldBindErrMsg})
		return
	}

	var m Movie

	dbWhereErr := db.DB.Where("id = ?", id).First(&m).Error
	if dbWhereErr != nil {
		log.Println(dbWhereErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": dbWhereErrMsg})
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": dbSaveErrMsg})
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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": dbWhereErrMsg})
		return
	}

	dbDeleteErr := db.DB.Delete(&m).Error
	if dbDeleteErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": dbDeleteErrMsg})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// TODO: Implement this function
func findAllByScore(c *gin.Context) {
	var res []response

	var params paginationParams

	bindQueryErr := c.BindQuery(&params)
	if bindQueryErr != nil {
		log.Println(bindQueryErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": bindQueryErrMsg})
		return
	}

	dbFindAllErr := db.DB.Raw(
		"SELECT * "+
			"FROM movies "+
			"ORDER BY ? ? "+
			"LIMIT ? OFFSET ?",
		params.SortBy, params.Sort,
		params.Size, params.Page*params.Size,
	).Scan(&res).Error
	if dbFindAllErr != nil {
		log.Println(dbFindAllErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": dbFindAllErrMsg})
		return
	}

	c.JSON(http.StatusOK, res)
}
