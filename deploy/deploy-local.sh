#!/bin/bash
# Local deployment script using .env file
# Usage: ./deploy-local.sh

set -e

echo "🚀 Deploying FortuneSpinner locally"

# Navigate to project root
cd "$(dirname "$0")/.."

# Check if .env file exists
if [ ! -f ".env" ]; then
    echo "❌ .env file not found!"
    echo "   Please create it from .env.example:"
    echo "   cp .env.example .env"
    echo "   Then edit .env with your configuration values"
    exit 1
fi

# Load environment variables from .env file
export $(cat .env | grep -v '^#' | xargs)

# Validate required environment variables
if [ "$DB_PASSWORD" = "CHANGE_ME" ] || [ -z "$DB_PASSWORD" ]; then
    echo "❌ DB_PASSWORD is not set or still has default value"
    echo "   Please edit .env file and set a strong password"
    exit 1
fi

echo "✅ Environment variables loaded from .env"

# Start Docker containers
echo "🐳 Starting Docker containers..."
docker-compose up -d --build

# Wait for MySQL to be ready
echo "⏳ Waiting for MySQL to be ready..."
until docker exec fortunespinner-mysql mysqladmin ping -h localhost --silent; do
    echo "   Waiting for MySQL..."
    sleep 2
done

# Run migrations
echo "🗄️  Running database migrations..."
docker exec -i fortunespinner-mysql mysql -uroot -p${DB_PASSWORD} ${DB_NAME} < migrations/001_initial_schema.sql || {
    echo "   Migrations might have already been applied, continuing..."
}

# Health check
echo "🏥 Performing health check..."
sleep 5
if curl -f http://localhost:8080/health > /dev/null 2>&1; then
    echo "✅ Deployment successful!"
    echo "🌐 Frontend: http://localhost:3000"
    echo "🔧 Backend API: http://localhost:8080"
else
    echo "⚠️  Health check failed, but services might still be starting..."
    echo "   Check status with: docker-compose ps"
fi

echo ""
echo "📊 Useful commands:"
echo "   View logs: docker-compose logs -f"
echo "   Stop services: docker-compose stop"
echo "   Remove everything: docker-compose down -v"