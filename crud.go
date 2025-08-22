package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bookapi/config"
	"bookapi/models"
)

// ----------------- HELPERS -----------------

// getUserID extracts the user_id from the Gin context, which should be set by JWT middleware.
// Returns the user ID and a boolean indicating success.
func getUserID(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		// If user_id is not found in context, return 401 Unauthorized
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return 0, false
	}
	return userID.(uint), true
}

// getNextUserLocalID calculates the next sequential ID for a user's Books or Movies.
// This is used so each user has their own independent numbering for books/movies.
func getNextUserLocalID(table string, userID uint) uint {
	var lastID uint
	switch table {
	case "books":
		var lastBook models.Books
		// Fetch the last book for the user, ordered by user_local_id descending
		config.DB.Where("user_id = ?", userID).Order("user_local_id desc").First(&lastBook)
		lastID = lastBook.UserLocalID
	case "movies":
		var lastMovie models.Movies
		// Fetch the last movie for the user, ordered by user_local_id descending
		config.DB.Where("user_id = ?", userID).Order("user_local_id desc").First(&lastMovie)
		lastID = lastMovie.UserLocalID
	}
	return lastID + 1
}

// ----------------- BOOKS -----------------

// GetBooks returns all books belonging to the authenticated user
func GetBooks(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	var books []models.Books
	config.DB.Where("user_id = ?", userID).Find(&books)
	c.JSON(http.StatusOK, books)
}

// GetBookByID returns a specific book for the authenticated user based on its local ID
func GetBookByID(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	id := c.Param("id")
	var book models.Books
	if err := config.DB.Where("user_id = ? AND user_local_id = ?", userID, id).First(&book).Error; err != nil {
		// Return 404 if the book is not found
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// PostBook allows the authenticated user to add one or more books
func PostBook(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	var newBooks []models.Books
	if err := c.ShouldBindJSON(&newBooks); err != nil {
		// Return 400 if the JSON body is invalid
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign the user ID and next sequential user_local_id for each book
	for i := range newBooks {
		newBooks[i].UserID = userID
		newBooks[i].UserLocalID = getNextUserLocalID("books", userID)
	}

	config.DB.Create(&newBooks) // Insert the books into the database
	c.JSON(http.StatusCreated, newBooks)
}

// UpdateBook updates a specific book of the authenticated user
func UpdateBook(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	id := c.Param("id")
	var book models.Books
	if err := config.DB.Where("user_id = ? AND user_local_id = ?", userID, id).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var updatedData models.Books
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update book fields
	book.Title = updatedData.Title
	book.Author = updatedData.Author
	book.Publisher = updatedData.Publisher
	book.ReleaseYear = updatedData.ReleaseYear

	config.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

// DeleteBook deletes a specific book of the authenticated user
func DeleteBook(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	id := c.Param("id")
	var book models.Books
	if err := config.DB.Where("user_id = ? AND user_local_id = ?", userID, id).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	config.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// ----------------- MOVIES -----------------

// GetMovies returns all movies belonging to the authenticated user
func GetMovies(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	var movies []models.Movies
	config.DB.Where("user_id = ?", userID).Find(&movies)
	c.JSON(http.StatusOK, movies)
}

// GetMovieByID returns a specific movie for the authenticated user
func GetMovieByID(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	id := c.Param("id")
	var movie models.Movies
	if err := config.DB.Where("user_id = ? AND user_local_id = ?", userID, id).First(&movie).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// PostMovie allows the authenticated user to add one or more movies
func PostMovie(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	var newMovies []models.Movies
	if err := c.ShouldBindJSON(&newMovies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign the user ID and next sequential user_local_id for each movie
	for i := range newMovies {
		newMovies[i].UserID = userID
		newMovies[i].UserLocalID = getNextUserLocalID("movies", userID)
	}

	config.DB.Create(&newMovies) // Insert the movies into the database
	c.JSON(http.StatusCreated, newMovies)
}

// UpdateMovie updates a specific movie of the authenticated user
func UpdateMovie(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	id := c.Param("id")
	var movie models.Movies
	if err := config.DB.Where("user_id = ? AND user_local_id = ?", userID, id).First(&movie).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	var updatedData models.Movies
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update movie fields
	movie.Title = updatedData.Title
	movie.ReleaseYear = updatedData.ReleaseYear
	movie.Rating = updatedData.Rating

	config.DB.Save(&movie)
	c.JSON(http.StatusOK, movie)
}

// DeleteMovie deletes a specific movie of the authenticated user
func DeleteMovie(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	id := c.Param("id")
	var movie models.Movies
	if err := config.DB.Where("user_id = ? AND user_local_id = ?", userID, id).First(&movie).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	config.DB.Delete(&movie)
	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}



