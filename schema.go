package models // This code belongs to the models package

// Defines what a Book looks like in the database
type Books struct {
	ID          uint   `json:"id" gorm:"primaryKey"`         // The unique number for this book in the database
	UserLocalID uint   `json:"user_local_id"`                // The number for this book that only matters to its owner (e.g., user's 1st book, 2nd book)
	Title       string `json:"title" binding:"required,min=2"`       // Book title, must be provided and at least 2 letters
	Author      string `json:"author" binding:"required,min=3"`      // Author name, must be provided and at least 3 letters
	Publisher   string `json:"publisher" binding:"required,min=3"`   // Publisher name, must be provided and at least 3 letters
	ReleaseYear int    `json:"release_year" binding:"required,gt=0"` // The year the book came out, must be a positive number
	UserID      uint   `json:"user_id"`                      // The ID of the user who owns this book
}

// Defines what a Movie looks like in the database
type Movies struct {
	ID          uint   `json:"id" gorm:"primaryKey"`                 // The unique number for this movie in the database
	UserLocalID uint   `json:"user_local_id"`                        // The number for this movie that only matters to its owner (e.g., user's 1st movie)
	Title       string `json:"title" binding:"required,min=2"`               // Movie title, must be provided and at least 2 letters
	ReleaseYear int    `json:"release_year" binding:"required,gt=1800"`      // The year the movie came out, must be after 1800
	Rating      int    `json:"rating" binding:"required,gte=1,lte=10"`       // User's rating (1-10 stars), must be between 1 and 10
	UserID      uint   `json:"user_id"`                              // The ID of the user who owns this movie
}