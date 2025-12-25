# Go User API (Fiber + SQLC + Zap)

A RESTful API built with Go to manage users with their `name` and `dob` (date of birth). The API calculates and returns a user's age dynamically when fetching user details.

## 🚀 Features

- RESTful API with GoFiber
- PostgreSQL database with SQLC for type-safe queries
- Dynamic age calculation from date of birth
- Input validation with go-playground/validator
- Structured logging with Uber Zap
- Request ID middleware for tracing
- Request duration logging
- Pagination support for listing users
- Docker support for easy deployment
- Unit tests for age calculation

## 📋 Requirements

- Go 1.21+
- PostgreSQL 12+
- `sqlc` (for code generation)
- (Optional) Docker & docker-compose

## 🛠️ Setup

### Option 1: Using Docker (Recommended)

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd go-backend
   ```

2. Start the services:
   ```bash
   docker-compose up -d
   ```

   This will:
   - Start PostgreSQL database
   - Run migrations automatically
   - Build and start the Go application

3. The API will be available at `http://localhost:8080`

4. To stop the services:
   ```bash
   docker-compose down
   ```

### Option 2: Manual Setup

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Install SQLC (if not already installed):
   ```bash
   go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
   ```

3. Set up PostgreSQL database:
   ```bash
   # Create database
   createdb users_db
   
   # Run migrations
   psql users_db < db/migrations/0001_create_users.sql
   ```

   Or using psql:
   ```bash
   psql -U postgres -d users_db -f db/migrations/0001_create_users.sql
   ```

4. Generate SQLC code:
   ```bash
   sqlc generate
   ```

5. Set environment variables:
   ```bash
   # Windows PowerShell
   $env:APP_PORT="8080"
   $env:DB_HOST="localhost"
   $env:DB_PORT="5432"
   $env:DB_USER="postgres"
   $env:DB_PASSWORD="postgres"
   $env:DB_NAME="users_db"
   $env:DB_SSLMODE="disable"
   ```

   ```bash
   # Linux/Mac
   export APP_PORT=8080
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASSWORD=postgres
   export DB_NAME=users_db
   export DB_SSLMODE=disable
   ```

6. Run the server:
   ```bash
   go run ./cmd/server
   ```

## 📡 API Endpoints

### Create User

**POST** `/users`

**Request:**
```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

**Response:**
```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}
```

### Get User by ID

**GET** `/users/:id`

**Response:**
```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}
```

### Update User

**PUT** `/users/:id`

**Request:**
```json
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

**Response:**
```json
{
  "id": 1,
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

### Delete User

**DELETE** `/users/:id`

**Response:**
- HTTP `204 No Content`

### List All Users

**GET** `/users?page=1&page_size=20`

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `page_size` (optional): Number of items per page (default: 20, max: 100)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 34
  },
  {
    "id": 2,
    "name": "Bob",
    "dob": "1985-12-25",
    "age": 39
  }
]
```

### Health Check

**GET** `/health`

**Response:**
```json
{
  "status": "ok"
}
```

## 🏗️ Project Structure

```
go-backend/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── config/
│   └── config.go            # Configuration management
├── db/
│   ├── migrations/          # Database migrations
│   │   └── 0001_create_users.sql
│   ├── queries/             # SQL queries for SQLC
│   │   └── users.sql
│   └── sqlc/                # Generated SQLC code
├── internal/
│   ├── handler/             # HTTP handlers
│   │   └── user_handler.go
│   ├── middleware/          # HTTP middleware
│   │   ├── request_id.go
│   │   └── logger.go
│   ├── models/              # Data models
│   │   └── user.go
│   ├── repository/          # Data access layer
│   │   └── user_repository.go
│   ├── routes/              # Route definitions
│   │   └── routes.go
│   ├── service/             # Business logic
│   │   ├── user_service.go
│   │   └── user_service_test.go
│   └── logger/              # Logger setup
│       └── logger.go
├── docker-compose.yml       # Docker Compose configuration
├── Dockerfile              # Docker image definition
├── go.mod                  # Go dependencies
├── go.sum                  # Go dependencies checksum
├── sqlc.yaml              # SQLC configuration
└── README.md              # This file
```

## 🧪 Testing

Run unit tests for age calculation:

```bash
go test ./internal/service
```

Run all tests with coverage:

```bash
go test ./... -cover
```

## 🔧 Middleware

The application includes the following middleware:

1. **Request ID Middleware**: 
   - Adds `X-Request-ID` header to all responses
   - Generates a unique UUID for each request if not provided

2. **Request Logger Middleware**:
   - Logs request method, path, status code, duration, and request ID
   - Uses structured logging with Zap

## 📝 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_PORT` | Server port | `8080` |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `5432` |
| `DB_USER` | Database user | `postgres` |
| `DB_PASSWORD` | Database password | `postgres` |
| `DB_NAME` | Database name | `users_db` |
| `DB_SSLMODE` | SSL mode | `disable` |

## 🐳 Docker

### Build Docker Image

```bash
docker build -t go-backend .
```

### Run with Docker Compose

```bash
docker-compose up -d
```

### View Logs

```bash
docker-compose logs -f app
```

### Stop Services

```bash
docker-compose down
```

### Remove Volumes (Clean Database)

```bash
docker-compose down -v
```

## 🔍 Validation

The API validates:

- **Name**: Required, 1-100 characters
- **DOB**: Required, must be in `YYYY-MM-DD` format

## 📊 Database Schema

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🚦 Error Handling

The API returns appropriate HTTP status codes:

- `200 OK`: Successful GET/PUT request
- `201 Created`: Successful POST request
- `204 No Content`: Successful DELETE request
- `400 Bad Request`: Invalid input or validation error
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

Error responses follow this format:

```json
{
  "error": "Error message description"
}
```

## 📦 Tech Stack

- **Framework**: [GoFiber](https://gofiber.io/) v2
- **Database**: PostgreSQL with [SQLC](https://sqlc.dev/) for type-safe queries
- **Logging**: [Uber Zap](https://github.com/uber-go/zap)
- **Validation**: [go-playground/validator](https://github.com/go-playground/validator)

## 📄 License

This project is open source and available under the MIT License.

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!

## 📧 Contact

For questions or support, please open an issue in the repository.
