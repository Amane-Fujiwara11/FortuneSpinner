#!/bin/bash
# AWS deployment script using Parameter Store or Secrets Manager
# Usage: ./deploy-aws.sh [environment]

set -e

ENVIRONMENT=${1:-production}
PROJECT_NAME="fortunespinner"

echo "🚀 Deploying FortuneSpinner to AWS ($ENVIRONMENT)"

# Check AWS CLI is installed
if ! command -v aws &> /dev/null; then
    echo "❌ AWS CLI is not installed. Please install it first."
    exit 1
fi

# Check AWS credentials
if ! aws sts get-caller-identity &> /dev/null; then
    echo "❌ AWS credentials not configured. Run 'aws configure' first."
    exit 1
fi

echo "📦 Loading environment variables from AWS Parameter Store..."

# Load parameters from AWS Systems Manager Parameter Store
export DB_HOST=$(aws ssm get-parameter \
    --name "/${PROJECT_NAME}/${ENVIRONMENT}/db_host" \
    --query 'Parameter.Value' \
    --output text 2>/dev/null || echo "mysql")

export DB_PORT=$(aws ssm get-parameter \
    --name "/${PROJECT_NAME}/${ENVIRONMENT}/db_port" \
    --query 'Parameter.Value' \
    --output text 2>/dev/null || echo "3306")

export DB_USER=$(aws ssm get-parameter \
    --name "/${PROJECT_NAME}/${ENVIRONMENT}/db_user" \
    --query 'Parameter.Value' \
    --output text 2>/dev/null || echo "root")

export DB_PASSWORD=$(aws ssm get-parameter \
    --name "/${PROJECT_NAME}/${ENVIRONMENT}/db_password" \
    --with-decryption \
    --query 'Parameter.Value' \
    --output text)

export DB_NAME=$(aws ssm get-parameter \
    --name "/${PROJECT_NAME}/${ENVIRONMENT}/db_name" \
    --query 'Parameter.Value' \
    --output text 2>/dev/null || echo "fortunespinner")

export PORT=$(aws ssm get-parameter \
    --name "/${PROJECT_NAME}/${ENVIRONMENT}/port" \
    --query 'Parameter.Value' \
    --output text 2>/dev/null || echo "8080")

export REACT_APP_API_URL=$(aws ssm get-parameter \
    --name "/${PROJECT_NAME}/${ENVIRONMENT}/api_url" \
    --query 'Parameter.Value' \
    --output text 2>/dev/null || echo "/api")

# Validate required parameters
if [ -z "$DB_PASSWORD" ]; then
    echo "❌ Required parameter DB_PASSWORD not found in Parameter Store"
    echo "   Please create it with:"
    echo "   aws ssm put-parameter --name '/${PROJECT_NAME}/${ENVIRONMENT}/db_password' --value 'your-password' --type 'SecureString'"
    exit 1
fi

echo "✅ Environment variables loaded successfully"

# Navigate to project root
cd "$(dirname "$0")/.."

# Pull latest changes
echo "📥 Pulling latest code..."
git pull origin main

# Deploy with Docker Compose
echo "🐳 Starting Docker containers..."
docker-compose -f docker-compose.prod.yml up -d --build

# Wait for services to be healthy
echo "⏳ Waiting for services to be ready..."
sleep 10

# Run migrations if needed
if [ ! -f ".migrations_done_${ENVIRONMENT}" ]; then
    echo "🗄️  Running database migrations..."
    docker exec -i fortunespinner-mysql mysql -u${DB_USER} -p${DB_PASSWORD} ${DB_NAME} < migrations/001_initial_schema.sql
    touch ".migrations_done_${ENVIRONMENT}"
fi

# Health check
echo "🏥 Performing health check..."
if curl -f http://localhost/api/health > /dev/null 2>&1; then
    echo "✅ Deployment successful!"
    echo "🌐 Application is running at http://localhost"
else
    echo "❌ Health check failed. Check logs with: docker-compose -f docker-compose.prod.yml logs"
    exit 1
fi

echo "📊 To view logs: docker-compose -f docker-compose.prod.yml logs -f"