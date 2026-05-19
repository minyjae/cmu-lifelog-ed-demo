
# CMU LIFELONG EDUCATION FOR STAFF MANAGEMENT API #Backend

Backend Hexagonal or Clean  Architecture with go language and fiber, gorm

## Architecture
```
fiber-lifelong-ed-api/
├── cmd/
│   ├── api/                   # Application entry point
│   │   └── main.go
│   └── migrate/               # Database migration CLI
│       └── main.go
├── internal/
│   ├── adapters/              # External adapters
│   │   ├── http/              # HTTP layer (handlers, middleware, routes)
│   │   │   ├── handlers/      # HTTP request handlers
│   │   │   ├── middleware/    # HTTP middleware
│   │   │   └── routes/        # Route definitions
│   │   └── persistence/       # Database layer
│   │       ├── models/        # Database models (GORM)
│   │       └── repositories/  # Data access layer
│   ├── config/                # Configuration management
│   │   ├── config.go          # App configuration
│   │   ├── database.go        # Database setup & migration
│   │   └── seeder.go          # Database seeding (Admin user)
│   └── core/                  # Business logic core
│       ├── domain/            # Domain entities and interfaces
│       │   ├── entities/      # Business entities
│       │   └── ports/         # Interfaces (ports)
│       │       ├── repositories/  # Repository interfaces
│       │       └── services/      # Service interfaces
│       └── services/          # Business logic services
├── pkg/utils/                 # Shared utilities
│   ├── jwt.go                 # JWT utilities
│   └── validator.go           # Validation utilities
├── docs/                      # API documentation (Swagger)
├── scripts/                   # Utility scripts
├── tmp/                       # Temporary build files
├── .env                       # Environment variables (local)
├── .env.example               # Environment variables template
├── .gitignore                 # Git ignore rules
├── .air.toml                  # Hot reload configuration
├── docker-compose.yml         # Docker services
├── Makefile                   # Build commands
├── go.mod                     # Go modules
└── go.sum                     # Go modules checksum
```

## Tech Stack

- **Go** with [Fiber](https://gofiber.io/)
- **JWT** for authentication
- **Swagger** for API documentation
- **Air** for live-reloading
- **PostgreSQL** (assumed)


## SET UP STEP

 - go mod tidy
 - docker compose up -d
 - go run .\cmd\api\main.go
 - or can hot reload following this step
    - swag init -g cmd/api/main.go -o docs
    - air


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file and you can see example in .env.example


## API Reference

#### SignIn

```http
  POST /api/auth
```
## User and Staff can access
#### Get ListQueue By Faculty

```http
  GET /api/listqueue/faculty
```

#### Get ListQueue By UserStatusID

```http
  GET /api/listqueue/user/status/:user_status_id

  ex. /api/listqueue/user/status/1
```
