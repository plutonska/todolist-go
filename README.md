﻿
# Todo List API (todolist-go)

The Todo List API is a Go-based application designed to manage a to-do list. It is built using clean architecture with key components such as controllers, services, repositories, and databases. This project also supports Docker for easy deployment.

## Features
- CRUD (Create, Read, Update, Delete) operations for the to-do list.
- Mark tasks as completed.
- Store and retrieve data from Redis (via repository).
- Use GORM for database management.

## Requirements
Before running this project, make sure you have:
- Go (minimum version 1.18)
- Docker & Docker Compose
- Redis (if running locally without Docker)

## Installation and Running the Application

### 1. Clone the Repository
Clone this repository to your local directory:

```bash
git clone github.com/plutonska/todolist-go
cd todolist-go
```

### 2. Using Docker
This project comes with a `docker-compose.yml` file to help you run the application and its dependencies like Redis.

1. **Build and run the services:**
   ```bash
   docker-compose up --build
   ```

2. The application will be available at `http://localhost:8080`.

### 3. Running Locally
If you don't want to use Docker, you can also run the application locally.

1. **Install dependencies:**
   ```bash
   go mod download
   ```

2. **Run the application:**
   ```bash
   go run cmd/todolist/main.go
   ```

3. The application will run at `http://localhost:8080`.

## Configuration
You can configure some settings by editing the `.env` or `.env.example` file.

### Example `.env` configuration:
```env
REDIS_URL=redis://localhost:6379
DATABASE_URL=postgres://user:password@localhost:5432/todolist
```

## Project Structure
Here is the main structure of the project:

```
.
├── cmd
│   └── todolist
│       └── main.go              # Application entry point
├── internal
│   ├── app
│   │   ├── controller
│   │   │   └── todo_controller.go  # Todo Controller
│   │   ├── models
│   │   │   └── domain
│   │   │       └── todo.go        # Todo Model
│   │   ├── repository
│   │   │   └── todo_repository.go # Repository for managing Todo
│   │   └── service
│   │       └── todo_service.go    # Business logic for Todo
│   └── pkg
│       └── databases
│           ├── init.go            # Database initialization
│           └── migration.go       # Database migration
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

## API Endpoints
Here are the main endpoints of the Todo List API:

- **GET /todos**: Retrieve all todos.
- **GET /todos/{id}**: Retrieve a todo by ID.
- **POST /todos**: Create a new todo.
- **PUT /todos/{id}**: Update a todo by ID.
- **DELETE /todos/{id}**: Delete a todo by ID.

## License
This project is licensed under the [MIT License](LICENSE.md).
