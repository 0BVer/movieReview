package delivery

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movieReview/movie/domain"
	"movieReview/movie/domain/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func testSetupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func getResponse() domain.Response {
	tempTime := time.Now()
	return domain.Response{
		Title:      "Test Movie",
		Genre:      "Action",
		IsShowing:  true,
		ReleasedAt: tempTime,
		EndAt:      tempTime.Add(time.Hour * 24 * 7),
		CreatedAt:  tempTime,
		UpdatedAt:  tempTime,
	}
}

func TestDelivery_Create_Success(t *testing.T) {
	r := testSetupRouter()
	mockUseCase := &mocks.MovieUseCase{}
	mockUseCase.On(
		"Create",
		mock.AnythingOfType("domain.CreateRequest"),
	).Return(
		getResponse(),
		nil,
	)

	NewDelivery(r.Group("/api/v1"), mockUseCase)

	reader := strings.NewReader(
		`{
				"title": "Test Movie", 
				"genre": "Action", 
				"releasedAt": "2023-09-24T00:00:00Z",
	  			"endAt": "2023-09-15T00:00:00Z"
			}`,
	)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/movies", reader)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	mockUseCase.AssertExpectations(t)
}

func TestDelivery_Create_BadRequest(t *testing.T) {
	r := testSetupRouter()
	mockUseCase := &mocks.MovieUseCase{}
	mockUseCase.On(
		"Create",
		mock.AnythingOfType("domain.CreateRequest"),
	).Return(
		domain.Response{},
		errors.New("create error"),
	)

	NewDelivery(r.Group("/api/v1"), mockUseCase)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/movies", nil)
	//req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	mockUseCase.AssertExpectations(t)
}

func TestFindAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &mocks.MovieUseCase{}
	handler := Delivery{useCase: mockUseCase}

	router := gin.New()
	router.GET("/movies", handler.FindAll)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []domain.Response
	_ = json.Unmarshal(w.Body.Bytes(), &response)

}
