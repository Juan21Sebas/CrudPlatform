# CrudPlatform

CrudPlatform is a Go-based backend API that manages users, challenges, and videos. It provides CRUD operations for each entity and includes pagination.

## Project Structure

```
CrudPlatform/
├── .vscode/
├── cmd/
│   └── config/
│       └── db/
├── internal/
│   ├── adapters/
│   │   ├── handlers/
│   │   │   └── http/
│   │   │       ├── middleware/
│   │   │       ├── challengeHandlers.go
│   │   │       ├── routes.go
│   │   │       ├── server.go
│   │   │       ├── userHandlers.go
│   │   │       └── videoHandlers.go
│   │   └── repository/
│   │       ├── api.go
│   │       ├── challengeTransactions.go
│   │       ├── usersTransactions.go
│   │       └── videoTransaction.go
│   └── core/
│       ├── domain/
│       │   └── repository/
│       │       ├── model/
│       │       │   ├── challenges/
│       │       │   ├── users/
│       │       │   └── videos/
│       │       └── schema/
│       │           ├── challenges/
│       │           ├── users/
│       │           └── videos/
│       ├── ports/
│       │   └── mocks/
│       └── services/
│           ├── challengeServices.go
│           ├── usersServices.go
│           └── videoServices.go
├── k8s/
├── postmanCollection/
├── swagger/
├── vendor/
├── scripts/
├── .gitignore
├── .gitlab-ci.yml
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md
```

## Features

- CRUD operations for users, challenges, and videos
- Pagination with a maximum of 10 results per page
- Authentication middleware
- Hexagonal architecture (ports and adapters)
- Domain-driven design
- Swagger documentation
- Docker support
- Kubernetes configuration
- Unit tests for core services and repositories

## Prerequisites

- Go 1.x
- Docker (optional)
- Kubernetes (optional)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/Juan21Sebas/CrudPlatform.git
   ```

2. Navigate to the project directory:
   ```
   cd CrudPlatform
   ```

3. Install dependencies:
   ```
   go mod download
   ```

## Running the Application

### Locally

1. Set up your environment variables (if any).

2. Run the application:
   ```
   go run main.go
   ```

### Using Docker

1. Build the Docker image:
   ```
   docker build -t crudplatform .
   ```

2. Run the container:
   ```
   docker run -p 8080:8080 crudplatform
   ```

### Using Docker Compose

Run the application and its dependencies:

```
docker-compose up
```

## API Documentation

API documentation is available via Swagger. After running the application, visit `/swagger/index.html` in your browser.

A Postman collection is also available in the `postmanCollection/` directory for testing the API endpoints.

## Testing

Run the tests using:

```
go test ./...
```

The project includes unit tests for core services and repository functions.

## CI/CD

This project includes a `.gitlab-ci.yml` file for GitLab CI/CD pipelines. Adjust as needed for your CI/CD platform.

## Deployment

Kubernetes configurations are available in the `k8s/` directory. These include:

- `deployment.yaml`: Defines the deployment for the application
- `services.yaml`: Defines the Kubernetes services
- `ingress.yaml`: Configures the ingress for external access

Modify these as needed for your specific deployment environment.

## Project Structure Details

- `cmd/`: Contains the main application setup, including database configuration.
- `internal/`: Houses the core application code.
  - `adapters/`: Implements the interface adapters (handlers and repositories).
  - `core/`: Contains the business logic and domain models.
    - `domain/`: Defines the domain models and schemas.
    - `ports/`: Defines the interfaces for the application.
    - `services/`: Implements the core business logic.
- `k8s/`: Kubernetes configuration files.
- `postmanCollection/`: Contains a Postman collection for API testing.
- `swagger/`: Swagger documentation for the API.

### Autor
#### Juan Sebastian Sanchez Arteta

