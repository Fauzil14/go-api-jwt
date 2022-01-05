package controllers

import (
	"go-api-jwt/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// struct for input
type MovieInput struct {
	Title               string `json:"title"`
	Year                int    `json:"year"`
	AgeRatingCategoryID int    `json:"age_rating_category_id"`
}

// Get All Movie godoc
// @Summary Get All Movie
// @Description Get List of Movie
// @Tags Movie
// @Produce json
// @Success 200 {object} []models.Movie
// @Router /movies [get]
func GetAllMovie(c *gin.Context) {
	// gunakan *gin.Context ketika ada request
	// get db from gin gi context, because db connection is in main function
	db := c.MustGet("db").(*gorm.DB)

	var movies []models.Movie

	db.Find(&movies)

	// return json
	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// Create Movie godoc
// @Summary Create Movie
// @Description Get List of Movie
// @Tags Movie
// @Param Body body MovieInput true "the body to create new Movie"
// @Produce json
// @Success 200 {object} models.Movie
// @Router /movies [post]
func CreateMovie(c *gin.Context) {
	var input MovieInput

	// Validasi inputan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie := models.Movie{Title: input.Title, Year: input.Year, AgeRatingCategoryID: input.AgeRatingCategoryID}

	db := c.MustGet("db").(*gorm.DB)

	db.Create(&movie)
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// Get a Movie godoc
// @Summary Get a Movie by id
// @Description Get one Movie by id
// @Tags Movie
// @Produce json
// @Param id path string true "Movie id"
// @Success 200 {object} models.Movie
// @Router /movies/{id} [get]
func GetMovieById(c *gin.Context) {
	var movie models.Movie

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// return json
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// Update Movie godoc
// @Summary Update a Movie by id
// @Description Update one Movie by id
// @Tags Movie
// @Produce json
// @Param id path string true "Movie id"
// @Param Body body MovieInput true "the body to update Movie"
// @Success 200 {object} models.Movie
// @Router /movies/{id} [patch]
func UpdateMovie(c *gin.Context) {
	// validasi inputan
	var input MovieInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find data in database (exist or not)
	var movie models.Movie

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var updatedInputMovie models.Movie

	updatedInputMovie.Title = input.Title
	updatedInputMovie.Year = input.Year
	updatedInputMovie.AgeRatingCategoryID = input.AgeRatingCategoryID
	updatedInputMovie.UpdatedAt = time.Now()

	db.Model(&movie).Updates(updatedInputMovie)
	// return json
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// Delete a Movie godoc
// @Summary Delete a Movie by id
// @Description Delete one Movie by id
// @Tags Movie
// @Produce json
// @Param id path string true "Movie id"
// @Success 200 {object} map[string]boolean
// @Router /movies/{id} [delete]
func DeleteMovie(c *gin.Context) {
	var movie models.Movie

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&movie)

	// return json
	c.JSON(http.StatusOK, gin.H{"data": true})
}
