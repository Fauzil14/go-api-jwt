package routes

import (
	"go-api-jwt/controllers"
	"go-api-jwt/middlewares"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"
)

func SetUpRouter(db *gorm.DB) *gin.Engine {
	// menandakan router
	r := gin.Default()

	// set db to gin context -> that in can acess globally i
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	// middleware for rating
	// group with prefix
	ratingMiddlewareRuote := r.Group("/age-rating-categories")
	ratingMiddlewareRuote.Use(middlewares.JwtAuthMiddleware())
	ratingMiddlewareRuote.POST("", controllers.CreateRating)
	ratingMiddlewareRuote.PATCH("/:id", controllers.UpdateRating)
	ratingMiddlewareRuote.DELETE("/:id", controllers.DeleteRating)

	r.GET("/age-rating-categories", controllers.GetAllRating)
	r.GET("/age-rating-categories/:id", controllers.GetRatingById)
	r.GET("/age-rating-categories/:id/movies", controllers.GetMoviesByAgeRatingCategoryId)

	moviesMiddlewareRoute := r.Group("/movies")
	moviesMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	moviesMiddlewareRoute.POST("", controllers.CreateMovie)
	moviesMiddlewareRoute.PATCH("/:id", controllers.UpdateMovie)
	moviesMiddlewareRoute.DELETE("/:id", controllers.DeleteMovie)

	r.GET("/movies", controllers.GetAllMovie)
	r.GET("/movies/:id", controllers.GetMovieById)

	// for swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
