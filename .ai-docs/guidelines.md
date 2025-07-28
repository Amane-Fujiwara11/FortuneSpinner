## Project Overview

FortuneSpinner is a gacha (lottery) web application with a point system. Users can spin the gacha to earn items of different rarities and accumulate points. The project implements Clean Architecture principles in both backend and frontend.

### Key Features
- **Gacha System**: Probability-based item lottery with 4 rarity tiers
- **Point System**: Earn points from gacha spins, track balances and transaction history
- **Clean Architecture**: Separation of concerns across domain, use case, interface, and infrastructure layers
- **Containerized**: Fully dockerized development and production environments
- **Type-Safe**: Full TypeScript support in frontend, strongly typed Go backend

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

- **Infrastructure Layer** (`/backend/infrastructure/`)
  - `mysql/`: Database connection and configuration
  - `repository/`: Repository implementations (moved from interface layer)
  - `container.go`: Dependency injection container

**Key Features:**
- Pure Go implementation using only standard library (no web frameworks)
- Clean Architecture with proper dependency direction (Domain ← UseCase ← Interface ← Infrastructure)
- Rich domain models with business rules and validation
- Dependency injection container for clean component initialization
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

## Code Style Guidelines

### General Principles
- **Follow Clean Architecture**: Keep dependencies pointing inward (Domain ← UseCase ← Interface ← Infrastructure)
- **No Framework Lock-in**: Backend uses Go standard library, frontend uses minimal dependencies
- **Type Safety First**: Use proper types in TypeScript, avoid `any`
- **Error Handling**: Always handle errors explicitly, return meaningful error messages
- **Consistent Naming**: Use descriptive names, follow language conventions (camelCase for JS/TS, PascalCase for Go types)

### Backend (Go)
- Use standard Go error handling patterns
- Keep handlers thin - business logic belongs in use cases
- Repository interfaces in domain layer, implementations in infrastructure
- No global variables except for configuration
- Always validate inputs at the handler level

### Frontend (React/TypeScript)
- Functional components with hooks
- Keep components focused and reusable
- Business logic in use cases, not components
- Use proper TypeScript types for all API responses
- Handle loading and error states in all API calls
- LocalStorage for session persistence (user data saved on creation)

## Development Commands

### Quick Start (Fast Development Mode)
```bash
# Use the new Makefile for faster development
make dev      # Start development environment with hot reload
make dev-logs # View logs
make stop     # Stop services
make down     # Remove containers

# Or manually:
docker-compose -f docker-compose.dev.yml up -d
```

### Makefile Commands
```bash
make help           # Show all available commands
make dev           # Start development environment
make prod          # Start production environment
make build         # Build Docker images
make rebuild       # Force rebuild without cache
make mysql         # Connect to MySQL CLI
make backend-logs  # View backend logs
make frontend-logs # View frontend logs
make ps            # Show running containers
```

### Production Mode
```bash
# Start production environment
make prod

# Or manually:
docker-compose up -d --build

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
│   │   ├── model/            # Entity definitions with business rules
│   │   └── repository/       # Repository interfaces
│   ├── usecase/              # Use case layer (business logic)
│   │   ├── gacha/           # Gacha-related use cases
│   │   └── point/           # Point-related use cases
│   ├── interface/            # Interface layer
│   │   └── handler/         # HTTP handlers
│   ├── infrastructure/       # Infrastructure layer
│   │   ├── mysql/          # Database configuration
│   │   ├── repository/     # Repository implementations
│   │   └── container.go    # Dependency injection container
│   ├── go.mod               # Go module definition
│   ├── go.sum               # Go dependencies lock file
│   ├── main.go              # Application entry point
│   ├── Dockerfile           # Docker configuration
│   └── .dockerignore        # Docker build exclusions
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
│   ├── package.json         # Node dependencies (includes proxy config)
│   ├── tsconfig.json        # TypeScript configuration
│   ├── Dockerfile           # Docker configuration
│   ├── nginx.conf          # Nginx configuration for production
│   └── .dockerignore        # Docker build exclusions
│
├── migrations/               # Database migrations
│   ├── 001_initial_schema.sql  # Initial database schema
│   └── migrate.sh              # Migration runner script
│
├── deploy/                    # Deployment scripts
│   ├── aws/                  # AWS-specific deployment
│   ├── local/                # Local development scripts
│   └── production/           # Production deployment scripts
├── docker-compose.yml       # Docker services configuration (production)
├── docker-compose.dev.yml   # Docker services configuration (development with hot reload)
├── Makefile                # Development command shortcuts
├── README.md               # Project readme
├── DEPLOYMENT.md           # Deployment documentation
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
- ✅ LocalStorage for user session persistence (auto-restore on page reload)
- ✅ Date display formatting in history
- ✅ AWS deployment infrastructure
- ✅ Fast development mode with hot reload (docker-compose.dev.yml)
- ✅ Makefile for convenient development commands
- ✅ Optimized Docker builds with .dockerignore files

### Pending Features
- ⏳ Google OAuth login
- ⏳ Fortune telling feature (using points)
- ⏳ Ad network integration
- ⏳ User rankings
- ⏳ More gacha items and animations

## Testing Guidelines

### Backend Testing
```bash
cd backend
go test ./...
go test -v ./usecase/...  # Test specific package
go test -cover ./...      # With coverage
```

### Frontend Testing
```bash
cd frontend
npm test                  # Run all tests
npm test -- --coverage    # With coverage
npm test -- --watch       # Watch mode
```

### Integration Testing
- Test API endpoints with actual database
- Use test containers for isolated database instances
- Mock external services when necessary

## Troubleshooting

### Common Issues

1. **Port Already in Use**
   ```bash
   # Find process using port
   lsof -i :3000  # Frontend
   lsof -i :8080  # Backend
   lsof -i :3306  # MySQL
   
   # Kill process
   kill -9 <PID>
   ```

2. **Database Connection Errors**
   - Ensure MySQL container is running: `docker ps`
   - Check logs: `docker-compose logs mysql`
   - Verify connection string in backend configuration

3. **CORS Issues**
   - Backend CORS middleware should allow frontend origin
   - In development, proxy is configured in frontend package.json
   - In production, nginx handles proxying

4. **Build Failures**
   - Clear Docker cache: `docker-compose build --no-cache`
   - Remove node_modules and reinstall: `rm -rf node_modules && npm install`
   - Update Go dependencies: `go mod tidy`

5. **JSON Parse Errors (Frontend)**
   - Ensure proxy is set in package.json: `"proxy": "http://backend:8080"`
   - Restart frontend container after adding proxy
   - Clear browser cache and reload

## Important Notes

### Backend
- Uses Go standard library exclusively (no Gin, Echo, etc.)
- All handlers follow the `http.HandlerFunc` interface
- CORS is handled manually in middleware
- Database operations use `database/sql` directly

### Frontend
- API calls go through nginx proxy in production
- Development uses proxy configuration in package.json (`"proxy": "http://backend:8080"`)
- All API responses must handle potential null values
- Components follow Clean Architecture separation
- **Important**: Proxy setting in package.json is required for development API calls

### Database
- MySQL runs in Docker with persistent volume
- Migrations should be versioned sequentially (001_, 002_, etc.)
- All tables use InnoDB engine with utf8mb4 charset

### Docker
- Frontend nginx serves static files and proxies API calls
- Backend connects to MySQL using container name
- Health checks ensure proper startup order

### Security
- Sensitive information stored in AWS Parameter Store / GCP Secret Manager / Azure Key Vault
- .env files excluded from version control
- Production deployment uses environment variables from secure storage
- Deploy scripts in `/deploy` directory for different environments
- Input validation at all entry points
- SQL injection prevention through parameterized queries
- XSS prevention through proper output escaping

## Contributing Guidelines

### Git Workflow
1. Create feature branch from main: `git checkout -b feature/your-feature`
2. Make atomic commits with clear messages
3. Run tests before committing
4. Push branch and create pull request
5. Ensure CI/CD passes before merging

### Commit Message Format
```
type(scope): subject

body (optional)

footer (optional)
```

Types: feat, fix, docs, style, refactor, test, chore

Example:
```
feat(gacha): add animation for legendary items

Implemented particle effects and sound for legendary drops
to enhance user experience.

Closes #123
```

### Code Review Checklist
- [ ] Follows Clean Architecture principles
- [ ] Has appropriate tests
- [ ] Handles errors properly
- [ ] Updates documentation if needed
- [ ] No hardcoded values or secrets
- [ ] Performance considerations addressed
- [ ] Accessibility requirements met (frontend)

## Performance Optimization

### Backend
- Use connection pooling for database
- Implement caching for frequently accessed data
- Batch database operations where possible
- Profile CPU and memory usage regularly

### Frontend
- Lazy load components and routes
- Optimize bundle size with code splitting
- Use React.memo for expensive components
- Implement virtual scrolling for long lists
- Optimize images and assets

## Deployment

### Environment Variables
Required environment variables for each service:

**Backend:**
- `DB_HOST`: MySQL host
- `DB_PORT`: MySQL port (default: 3306)
- `DB_USER`: Database user
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `PORT`: API server port (default: 8080)

**Frontend:**
- `REACT_APP_API_URL`: Backend API URL
- `NODE_ENV`: Environment (development/production)

### Production Checklist
- [ ] Environment variables configured
- [ ] Database migrations applied
- [ ] SSL certificates configured
- [ ] Monitoring and logging set up
- [ ] Backup strategy implemented
- [ ] Rate limiting configured
- [ ] Health checks passing