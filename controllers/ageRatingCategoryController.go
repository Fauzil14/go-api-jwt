package controllers

import (
	"go-api-jwt/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// struct for input
type AgeRatingCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Get All Rating godoc
// @Summary Get All Age Rating Category
// @Description Get List of Age Rating Category
// @Tags AgeRatingCategory
// @Produce json
// @Success 200 {object} []models.AgeRatingCategory
// @Router /age-rating-categories [get]
func GetAllRating(c *gin.Context) {
	// gunakan *gin.Context ketika ada request
	// get db from gin context, because db connection is in main function
	db := c.MustGet("db").(*gorm.DB)

	var ratings []models.AgeRatingCategory

	db.Find(&ratings)

	// return json
	c.JSON(http.StatusOK, gin.H{"data": ratings})
}

// Create Rating godoc
// @Summary Create Age Rating Category
// @Description Get List of Age Rating Category
// @Tags AgeRatingCategory
// @Param Body body AgeRatingCategoryInput true "the body to create new age rating category"
// @Param Authorizattion header string true "Authorization : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories [post]
func CreateRating(c *gin.Context) {
	var input AgeRatingCategoryInput

	// Validasi inputan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rating := models.AgeRatingCategory{Name: input.Name, Description: input.Description}

	db := c.MustGet("db").(*gorm.DB)

	db.Create(&rating)
	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// Get a Rating godoc
// @Summary Get an Age Rating Category by id
// @Description Get one Age Rating Category by id
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "Age Rating Category id"
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories/{id} [get]
func GetRatingById(c *gin.Context) {
	var rating models.AgeRatingCategory

	db := c.MustGet("db").(*gorm.DB)

	// tabel yang di cari berdasarkan struct yang di masukkan melalui parameter .First(&struct)
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// return json
	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// Get movies from one Rating godoc
// @Summary Get movies by Age Rating Category by id
// @Description Get all movies of Age Rating Category by id
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "Age Rating Category id"
// @Success 200 {object} []models.Movie
// @Router /age-rating-categories/{id}/movies [get]
func GetMoviesByAgeRatingCategoryId(c *gin.Context) {
	var movies []models.Movie

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("age_rating_category_id = ?", c.Param("id")).Find(&movies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// return json
	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// Update Rating godoc
// @Summary Update an Age Rating Category by id
// @Description Update one Age Rating Category by id
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "Age Rating Category id"
// @Param Body body AgeRatingCategoryInput true "the body to update Age Rating Category"
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories/{id} [patch]
func UpdateRating(c *gin.Context) {
	// validasi inputan
	var input AgeRatingCategoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find data in database (exist or not)
	var rating models.AgeRatingCategory

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var updatedInputRating models.AgeRatingCategory

	updatedInputRating.Name = input.Name
	updatedInputRating.Description = input.Description
	updatedInputRating.UpdatedAt = time.Now()

	db.Model(&rating).Updates(updatedInputRating)
	// return json
	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// Delete a Rating godoc
// @Summary Delete an Age Rating Category by id
// @Description Delete one Age Rating Category by id
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "Age Rating Category id"
// @Success 200 {object} map[string]boolean
// @Router /age-rating-categories/{id} [delete]
func DeleteRating(c *gin.Context) {
	var rating models.AgeRatingCategory

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&rating)

	// return json
	c.JSON(http.StatusOK, gin.H{"data": true})
}
