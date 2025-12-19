[README.md](https://github.com/user-attachments/files/24260534/README.md)
# AgeCal -- RESTful User API (Go + Fiber) {#agecal-restful-user-api-go-fiber}

AgeCal is a clean, production-style RESTful API built using **Go** and
**Fiber** that manages users with their **name** and **date of birth
(DOB)** and **calculates age dynamically** when fetching user details.

This project was developed as part of an **internship backend
assignment**, with a strong focus on: - Clean architecture - Separation
of concerns - Input validation - Proper HTTP status codes - Structured
logging

## 🚀 Features {#features}

- Create users with name and DOB
- Fetch user details with dynamically calculated age
- Input validation using `go-playground/validator`
- PostgreSQL database integration
- Structured logging using Uber `zap`
- Middleware-based request logging
- Clean folder structure (handler, service, repository)
- Proper HTTP status codes and error handling

## 🧱 Tech Stack {#tech-stack}

- **Language:** Go
- **Framework:** Fiber
- **Database:** PostgreSQL And SQLC(for type-safe code)
- **Validation:** go-playground/validator
- **Logging:** Uber zap
- **API Testing:** Postman

## 📁 Project Structure {#project-structure}

    AgeCal/
    │
    ├── cmd/
    │   └── server/
    │       └── main.go          # Application entry point
    │
    ├── internal/
    │   ├── handler/             # HTTP handlers (controllers)
    │   ├── service/             # Business logic
    │   ├── repository/          # Database access layer
    │   ├── middleware/          # Fiber middlewares (logger, request ID)
    │   ├── logger/              # Zap logger initialization
    │   ├── validator/           # Input validation logic
    │   └── dto/                 # Request/response DTOs
    │
    ├── migrations/               # SQL migration files
    ├── go.mod
    └── go.sum

## ⚙️ Setup Instructions {#setup-instructions}

### 1️⃣ Prerequisites {#prerequisites}

- Go 1.20+
- PostgreSQL
- Git

### 2️⃣ Clone the Repository {#clone-the-repository}

    git clone https://github.com/geekynoob097/AgeCal.git
   cd AgeCal

### 3️⃣ Database Setup {#database-setup}

Create a PostgreSQL database:

    CREATE DATABASE agecal;

Create `users` table:

    CREATE TABLE users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        dob DATE NOT NULL
    );

### 4️⃣ Configure Database Connection {#configure-database-connection}

Update database credentials inside your DB config file (or environment
variables):

    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=your_password
    DB_NAME=agecal

### 5️⃣ Run the Application {#run-the-application}

    go run cmd/server/main.go

Server will start at:

    http://localhost:8080

## 📌 API Endpoints {#api-endpoints}

### ➕ Create User {#create-user}

**POST** `/users`

Request Body:

    {
      "name": "Alice",
      "dob": "1990-05-10"
    }

Response (201 Created):

    {
      "id": 1,
      "name": "Alice",
      "dob": "1990-05-10",
      "age": 35
    }

### 📥 Get All Users {#get-all-users}

**GET** `/users`

Response (200 OK):

    [
      {
        "id": 1,
        "name": "Alice",
        "dob": "1990-05-10",
        "age": 35
      }
    ]

## 🛑 Validation & Error Handling {#validation-error-handling}

- Invalid JSON → `400 Bad Request`
- Validation failure → `400 Bad Request`
- Resource not found → `404 Not Found`
- Duplicate resource → `409 Conflict`
- Server error → `500 Internal Server Error`

Example validation error response:

    {
      "message": "Validation failed",
      "errors": {
        "Name": "is required",
        "DOB": "must be YYYY-MM-DD"
      }
    }

## 🧾 Logging {#logging}

- Structured logs using **zap**
- Request logging via middleware
- Request ID injected per request

Example log:

    {
      "level": "info",
      "msg": "request",
      "method": "GET",
      "path": "/users",
      "request_id": "abc-123"
    }

## 🧪 Testing {#testing}

- API tested using Postman
- Manual testing for edge cases and validation

## 📚 Learning Outcomes {#learning-outcomes}

- REST API design in Go
- Clean architecture & separation of concerns
- Input validation patterns
- Middleware usage in Fiber
- Structured logging
- Error handling best practices

## 👨‍💻 Author {#author}

**U Venkata Tharun**  
Intern Backend Developer

## 📜 License {#license}

This project is for **educational and internship purposes**.
