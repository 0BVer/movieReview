package review

import (
	"github.com/gin-gonic/gin"
	"log"
	"movieReview/config/database"
	"net/http"
)

var db = database.GetDB()

func Config(api *gin.RouterGroup) {
	_ = db.DB.AutoMigrate(&Review{})
	api.POST("/reviews", create)
	api.GET("/reviews", findAllByMovieId)
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

func findAllByMovieId(c *gin.Context) {
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
			"FROM reviews "+
			"WHERE (@MovieID = '' OR movie_id = @MovieID) "+
			"AND (@ScoreCap = '' OR score >= @ScoreCap) "+
			"ORDER BY created_at DESC",
		params,
	).Scan(&res).Error
	if dbFindAllErr != nil {
		log.Println(dbFindAllErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": dbFindAllErrMsg})
		return
	}

	c.JSON(http.StatusOK, res)
}
