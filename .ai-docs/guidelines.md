## Project Overview

FortuneSpinner is a gacha (lottery) web application with a point system. Users can spin the gacha to earn items of different rarities and accumulate points. The project implements Clean Architecture principles in both backend and frontend.

## Architecture

### Backend (Go)
Located in `/backend`, implementing Clean Architecture with the following layers:

- **Domain Layer** (`/backend/domain/`)
  - `model/`: Entity definitions (User, Gacha, GachaResult, Point)
  - `repository/`: Repository interfaces

- **Use Case Layer** (`/backend/usecase/`)
  - `gacha/`: Gacha execution and history logic
  - `point/`: Point balance and transaction logic

- **Interface Layer** (`/backend/interface/`)
  - `handler/`: HTTP handlers for REST API
  - `repository/`: Repository implementations

- **Infrastructure Layer** (`/backend/infrastructure/`)
  - `mysql/`: Database connection and configuration

**Key Features:**
- Pure Go implementation using only standard library (no web frameworks)
- Module name: `github.com/Amane-Fujiwara11/FortuneSpinner/backend`
- Go version: 1.20

### Frontend (React + TypeScript)
Located in `/frontend`, also implementing Clean Architecture:

- **Domain Layer** (`/frontend/src/domain/`)
  - Type definitions for User, Gacha, Point entities

- **Use Cases Layer** (`/frontend/src/usecases/`)
  - Business logic for gacha, points, and user operations

- **Interface Layer** (`/frontend/src/interface/`)
  - `components/`: Reusable UI components (PointDisplay, GachaSpinner, GachaHistory)
  - `pages/`: Page components (GachaPage)

- **Infrastructure Layer** (`/frontend/src/infrastructure/`)
  - `api/`: API client and endpoint definitions

### Database (MySQL)
- Running in Docker container
- Port: 3306
- Database name: fortunespinner
- Tables: users, gacha_results, user_points, point_transactions

## Development Commands

### Quick Start
```bash
# 1. Start MySQL
docker-compose up -d mysql

# 2. Run database migrations (first time only)
cd migrations && ./migrate.sh

# 3. Start all services
docker-compose up -d

# Access the application
# Frontend: http://localhost:3000
# Backend API: http://localhost:8080
```

### Docker Services
```bash
# Start all services
docker-compose up -d

# Stop services (containers remain)
docker-compose stop

# Start stopped services
docker-compose start

# Stop and remove containers
docker-compose down

# Stop and remove everything including data
docker-compose down -v

# View logs
docker-compose logs -f [service_name]

# Rebuild and restart specific service
docker-compose up -d --build [service_name]
```

### Database Operations
```bash
# Connect to MySQL
docker exec -it fortunespinner-mysql mysql -uroot -prootpassword fortunespinner

# Run migrations manually
docker exec -i fortunespinner-mysql mysql -uroot -prootpassword fortunespinner < migrations/001_initial_schema.sql

# View table data
docker exec fortunespinner-mysql mysql -uroot -prootpassword fortunespinner -e "SELECT * FROM users;"
```

### Backend Development
```bash
cd backend

# Download dependencies
go mod download
go mod tidy

# Run locally (requires MySQL to be running)
go run main.go

# Build
go build -o main .
```

### Frontend Development
```bash
cd frontend

# Install dependencies
npm install

# Run development server
npm start

# Build for production
npm run build

# Run tests
npm test
```

## Project Structure

```
.
├── backend/                    # Go backend application
│   ├── domain/                # Domain layer (entities, interfaces)
│   │   ├── model/            # Entity definitions
│   │   └── repository/       # Repository interfaces
│   ├── usecase/              # Use case layer (business logic)
│   │   ├── gacha/           # Gacha-related use cases
│   │   └── point/           # Point-related use cases
│   ├── interface/            # Interface layer
│   │   ├── handler/         # HTTP handlers
│   │   └── repository/      # Repository implementations
│   ├── infrastructure/       # Infrastructure layer
│   │   └── mysql/          # Database configuration
│   ├── go.mod               # Go module definition
│   ├── go.sum               # Go dependencies lock file
│   ├── main.go              # Application entry point
│   └── Dockerfile           # Docker configuration
│
├── frontend/                  # React frontend application
│   ├── src/
│   │   ├── domain/          # Domain models (TypeScript types)
│   │   ├── usecases/        # Business logic
│   │   ├── interface/       # UI layer
│   │   │   ├── components/  # Reusable components
│   │   │   └── pages/      # Page components
│   │   ├── infrastructure/  # External services
│   │   │   └── api/        # API client
│   │   ├── App.tsx         # Root component
│   │   ├── App.css         # Global styles
│   │   └── index.tsx       # Entry point
│   ├── package.json         # Node dependencies
│   ├── tsconfig.json        # TypeScript configuration
│   ├── Dockerfile           # Docker configuration
│   └── nginx.conf          # Nginx configuration for production
│
├── migrations/               # Database migrations
│   ├── 001_initial_schema.sql  # Initial database schema
│   └── migrate.sh              # Migration runner script
│
├── docker-compose.yml       # Docker services configuration
├── README.md               # Project readme
└── .ai-docs/              # AI assistant documentation
    └── guidelines.md      # This file
```

## API Endpoints

All API endpoints return JSON with the following structure:
```json
{
  "success": true|false,
  "data": {...} | null,
  "error": "error message" | null
}
```

### User Management
- `POST /api/users` - Create a new user
  - Body: `{"name": "username"}`
  - Returns: User object with ID

### Gacha Operations
- `POST /api/gacha/execute` - Execute a gacha spin
  - Body: `{"user_id": 1}`
  - Returns: GachaResult with item details and points earned

- `GET /api/gacha/history?user_id={id}&limit={limit}` - Get gacha history
  - Returns: Array of GachaResult objects

### Point Management
- `GET /api/points/balance?user_id={id}` - Get user's point balance
  - Returns: UserPoint object with current balance

- `GET /api/points/transactions?user_id={id}&limit={limit}` - Get point transaction history
  - Returns: Array of PointTransaction objects

### Health Check
- `GET /health` - Check if backend is running
  - Returns: `{"status": "ok"}`

## Database Schema

### users
```sql
id INT PRIMARY KEY AUTO_INCREMENT
name VARCHAR(255) NOT NULL
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
```

### gacha_results
```sql
id INT PRIMARY KEY AUTO_INCREMENT
user_id INT NOT NULL (FK -> users.id)
item_id INT NOT NULL
item_name VARCHAR(255) NOT NULL
rarity INT NOT NULL (1=Common, 2=Rare, 3=Epic, 4=Legendary)
points_earned INT NOT NULL
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
```

### user_points
```sql
id INT PRIMARY KEY AUTO_INCREMENT
user_id INT NOT NULL UNIQUE (FK -> users.id)
balance INT NOT NULL DEFAULT 0
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
```

### point_transactions
```sql
id INT PRIMARY KEY AUTO_INCREMENT
user_id INT NOT NULL (FK -> users.id)
amount INT NOT NULL
type VARCHAR(50) NOT NULL ('gacha' | 'spend')
description TEXT
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
```

## Gacha System

### Items and Probabilities
- Bronze Coin (Common): 60% chance, 10 points
- Silver Coin (Rare): 30% chance, 50 points
- Gold Coin (Epic): 8% chance, 200 points
- Diamond (Legendary): 2% chance, 1000 points

## Current Implementation Status

### Completed Features
- ✅ User creation and management
- ✅ Gacha spinning with probability-based item selection
- ✅ Point accumulation system
- ✅ Gacha history tracking
- ✅ Point transaction history
- ✅ Clean Architecture implementation (backend & frontend)
- ✅ Docker containerization
- ✅ Database migrations
- ✅ Responsive UI with animations

### Pending Features
- ⏳ Google OAuth login
- ⏳ Fortune telling feature (using points)
- ⏳ Ad network integration
- ⏳ User rankings
- ⏳ More gacha items and animations

## Important Notes

### Backend
- Uses Go standard library exclusively (no Gin, Echo, etc.)
- All handlers follow the `http.HandlerFunc` interface
- CORS is handled manually in middleware
- Database operations use `database/sql` directly

### Frontend
- API calls go through nginx proxy in production
- Development uses proxy configuration in package.json
- All API responses must handle potential null values
- Components follow Clean Architecture separation

### Database
- MySQL runs in Docker with persistent volume
- Migrations should be versioned sequentially (001_, 002_, etc.)
- All tables use InnoDB engine with utf8mb4 charset

### Docker
- Frontend nginx serves static files and proxies API calls
- Backend connects to MySQL using container name
- Health checks ensure proper startup order