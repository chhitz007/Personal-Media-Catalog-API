package routes

import (
	"github.com/gin-gonic/gin"
	"bookapi/controllers"
	"bookapi/middleware"
)

func RegisterRoutes(router *gin.Engine) {
	// ===== BOOK ROUTES =====
	// Protected routes (JWT required)
	bookRoutes := router.Group("/books")
	bookRoutes.Use(middleware.JWTMiddleware())
	{
		bookRoutes.GET("", controllers.GetBooks)          // Get all books for logged-in user
		bookRoutes.GET("/:id", controllers.GetBookByID)  // Get specific book by ID
		bookRoutes.POST("", controllers.PostBook)        // Create book
		bookRoutes.PUT("/:id", controllers.UpdateBook)   // Update book
		bookRoutes.DELETE("/:id", controllers.DeleteBook) // Delete book
	}

	// ===== MOVIE ROUTES =====
	// Protected routes (JWT required)
	movieRoutes := router.Group("/movies")
	movieRoutes.Use(middleware.JWTMiddleware())
	{
		movieRoutes.GET("", controllers.GetMovies)          // Get all movies for logged-in user
		movieRoutes.GET("/:id", controllers.GetMovieByID)   // Get specific movie by ID
		movieRoutes.POST("", controllers.PostMovie)         // Create movie
		movieRoutes.PUT("/:id", controllers.UpdateMovie)    // Update movie
		movieRoutes.DELETE("/:id", controllers.DeleteMovie) // Delete movie
	}
}



