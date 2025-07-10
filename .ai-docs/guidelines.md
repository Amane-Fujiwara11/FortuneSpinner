## Project Overview

FortuneSpinner is a web application project in early development stages with a Go backend and planned frontend. The project uses Docker for containerization with MySQL as the database.

## Architecture

- **Backend**: Go (Golang) application located in `/backend`
  - Entry point: `backend/main.go`
  - Module name: `backend`
  - Go version: 1.24.1
- **Frontend**: Located in `/frontend` (currently empty, awaiting implementation)

- **Database**: MySQL (latest) running in Docker container
  - Port: 3306
  - Root password: root
  - Data persisted in Docker volume

## Development Commands

### Docker Services

```bash
# Start the MySQL database
docker-compose up -d

# Stop services
docker-compose down

# View logs
docker-compose logs -f
```

### Backend Development

```bash
# Navigate to backend directory
cd backend

# Run the Go application
go run main.go

# Build the application
go build

# Download dependencies (if any are added)
go mod download
```

## Project Structure

```
.
├── backend/           # Go backend application
│   ├── go.mod        # Go module definition
│   ├── main.go       # Application entry point
│   └── .gitignore    # Go-specific gitignore
├── frontend/         # Frontend (to be implemented)
├── docker-compose.yml # Docker services configuration
└── README.md         # Project readme (minimal)
```

## Current Implementation Status

- Basic Go backend structure is set up
- Docker Compose configured with MySQL database
- Frontend directory created but not yet implemented
- No API routes or database connections established yet

## Important Notes

- The backend is not yet connected to the MySQL database
- No frontend framework has been chosen or implemented
- The project is in initial setup phase and ready for feature development
