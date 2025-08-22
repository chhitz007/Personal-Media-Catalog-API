package models

// User defines the structure for the users table in the database.
type User struct {
	ID       uint    `gorm:"primaryKey" json:"id"`  
	// ID is the primary key for the user table. GORM will auto-increment it.

	Username string  `gorm:"unique" json:"username" binding:"required,min=3"`  
	// Username must be unique in the database.
	// JSON tag allows this field to be serialized/deserialized as "username".
	// Binding ensures that when receiving input (e.g., via API), 
	// the username is required and must be at least 3 characters long.

	Password string  `json:"password" binding:"required,min=6"`  
	// Password is required for input and must be at least 6 characters long.
	// Stored as plain text here (consider hashing before saving for security!).

	Books    []Books  `gorm:"foreignKey:UserID"`  
	// One-to-many relationship: a user can have multiple books.
	// 'UserID' in the Books table will act as the foreign key linking back to this user.

	Movies   []Movies `gorm:"foreignKey:UserID"`  
	// One-to-many relationship: a user can have multiple movies.
	// 'UserID' in the Movies table links movies to their respective user.
}

