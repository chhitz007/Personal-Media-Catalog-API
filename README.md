# Personal Media Catalog API

A secure, production-ready RESTful API built in Go (Golang) that allows users to manage their personal collections of books and movies. The API features robust JWT authentication and enforces strict data isolation, ensuring users can only access their own data.

## 🛠️ Tech Stack

**Backend:** Go (Golang)  
**Framework:** Gin-Gonic  
**ORM:** GORM  
**Database:** PostgreSQL  
**Authentication:** JWT (JSON Web Tokens)  
**Password Hashing:** bcrypt  

## ✨ Features

- **Secure User Authentication:** JWT-based login and registration with hashed passwords.
- **Complete CRUD Operations:** Create, read, update, and delete books and movies.
- **Data Isolation:** All user queries are scoped by user ID, preventing access to other users' data.
- **RESTful Design:** Clean and predictable API endpoints.
- **User-Friendly Sequential IDs:** Items are numbered sequentially (1, 2, 3...) per user, instead of using global database IDs.
- **Well-Structured Code:** Separation of concerns into models, config, controllers, middleware, and routes.

## 📁 Project Structure

```
bookapi/
├── config/
│   └── database.go          # Database connection and auto-migration
├── controllers/
│   ├── auth.go              # Handlers for user registration and login
│   └── crud.go      # CRUD handlers for books and movies
├── middleware/
│   └── jwt.go               # JWT authentication middleware
├── models/
│   ├── user.go              # User model schema
│   └── schema.go            # Book and Movie model schemas
├── routes/
│   └── routes.go            # API route definitions
├── main.go                  # Application entry point
├── .env.example             # Example environment variables file
└── go.mod                   # Go module dependencies
```

## 🚀 Getting Started

### Prerequisites

- Go 1.18+
- PostgreSQL database
- A tool like [Thunder Client](https://www.thunderclient.io/) or [Postman](https://www.postman.com/) to test the API

### Installation

1.  **Clone the repository**
    ```bash
    git clone https://github.com/your-username/bookapi.git
    cd bookapi
    ```

2.  **Install dependencies**
    ```bash
    go mod download
    ```

3.  **Set up your database**
    - Create a new PostgreSQL database for the project.

4.  **Configure environment variables**
    - Copy the `.env.example` file to a new file called `.env`.
    - Fill in your specific database credentials and a strong JWT secret.

    ```env
    # Database Settings
    DB_HOST=localhost
    DB_USER=your_db_username
    DB_PASSWORD=your_db_password
    DB_NAME=your_database_name
    DB_PORT=5432

    # JWT Secret (Use a long, random string in production!)
    JWT_SECRET=your_super_secret_key_here

    # Server Port (optional)
    PORT=8080
    ```

5.  **Run the application**
    ```bash
    go run main.go
    ```
    The server will start on `http://localhost:8080` (or your specified PORT), and it will automatically migrate the database tables.

## 📚 API Documentation

### Authentication Endpoints

| Method | Endpoint    | Description | Body |
| :----- | :---------- | :---------- | :--- |
| `POST` | `/register` | Register a new user | `{"username": "string", "password": "string"}` |
| `POST` | `/login`    | Login and receive a JWT token | `{"username": "string", "password": "string"}` |

### Protected Endpoints (Require JWT in `Authorization: Bearer <token>` Header)

#### Books

| Method | Endpoint | Description | Body |
| :----- | :------- | :---------- | :--- |
| `GET` | `/books` | Get all books for the authenticated user | - |
| `GET` | `/books/{id}` | Get a specific book by its user-local ID | - |
| `POST` | `/books` | Create one or more new books | `[{"title": "string", "author": "string", "publisher": "string", "release_year": number}]` |
| `PUT` | `/books/{id}` | Update a specific book | `{"title": "string", "author": "string", "publisher": "string", "release_year": number}` |
| `DELETE` | `/books/{id}` | Delete a specific book | - |

#### Movies

| Method | Endpoint | Description | Body |
| :----- | :------- | :---------- | :--- |
| `GET` | `/movies` | Get all movies for the authenticated user | - |
| `GET` | `/movies/{id}` | Get a specific movie by its user-local ID | - |
| `POST` | `/movies` | Create one or more new movies | `[{"title": "string", "release_year": number, "rating": number}]` |
| `PUT` | `/movies/{id}` | Update a specific movie | `{"title": "string", "release_year": number, "rating": number}` |
| `DELETE` | `/movies/{id}` | Delete a specific movie | - |

## 🧪 Testing the API

1.  **Register a User:**
    ```http
    POST http://localhost:8080/register
    Content-Type: application/json

    {
        "username": "john_doe",
        "password": "securepassword123"
    }
    ```

2.  **Login to Get a Token:**
    ```http
    POST http://localhost:8080/login
    Content-Type: application/json

    {
        "username": "john_doe",
        "password": "securepassword123"
    }
    ```
    *Save the `token` from the response.*

3.  **Access a Protected Route:**
    ```http
    GET http://localhost:8080/books
    Authorization: Bearer <your_token_here>
    ```

## 🔒 Security Highlights

- Passwords are hashed using **bcrypt** before being stored in the database.
- All protected routes require a valid **JWT token**.
- **Data isolation** is enforced at the database query level using `WHERE user_id = ?`, making it impossible for users to access each other's data.
- The JWT secret is stored as an environment variable, not in the codebase.



---

**⭐ Star this repo if you found it helpful!**
