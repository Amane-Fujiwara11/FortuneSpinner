#!/bin/bash
# Script to setup AWS Parameter Store parameters for FortuneSpinner
# Usage: ./setup-aws-parameters.sh [environment]

set -e

ENVIRONMENT=${1:-production}
PROJECT_NAME="fortunespinner"

echo "🔧 Setting up AWS Parameter Store for ${PROJECT_NAME} (${ENVIRONMENT})"

# Check AWS CLI
if ! command -v aws &> /dev/null; then
    echo "❌ AWS CLI is not installed. Please install it first."
    exit 1
fi

# Function to create parameter
create_parameter() {
    local name=$1
    local value=$2
    local type=${3:-String}
    local description=$4
    
    echo "Creating parameter: ${name}"
    aws ssm put-parameter \
        --name "${name}" \
        --value "${value}" \
        --type "${type}" \
        --description "${description}" \
        --overwrite \
        2>/dev/null || echo "  Parameter might already exist, skipping..."
}

# Create parameters
echo "📝 Creating parameters..."

create_parameter \
    "/${PROJECT_NAME}/${ENVIRONMENT}/db_host" \
    "mysql" \
    "String" \
    "Database host endpoint"

create_parameter \
    "/${PROJECT_NAME}/${ENVIRONMENT}/db_port" \
    "3306" \
    "String" \
    "Database port"

create_parameter \
    "/${PROJECT_NAME}/${ENVIRONMENT}/db_user" \
    "root" \
    "String" \
    "Database username"

create_parameter \
    "/${PROJECT_NAME}/${ENVIRONMENT}/db_name" \
    "fortunespinner" \
    "String" \
    "Database name"

create_parameter \
    "/${PROJECT_NAME}/${ENVIRONMENT}/port" \
    "8080" \
    "String" \
    "Backend API port"

create_parameter \
    "/${PROJECT_NAME}/${ENVIRONMENT}/api_url" \
    "/api" \
    "String" \
    "Frontend API URL"

# Prompt for secure password
echo ""
echo "⚠️  IMPORTANT: You need to manually set the database password"
echo "Run the following command with your secure password:"
echo ""
echo "aws ssm put-parameter \\"
echo "  --name '/${PROJECT_NAME}/${ENVIRONMENT}/db_password' \\"
echo "  --value 'YOUR_SECURE_PASSWORD_HERE' \\"
echo "  --type 'SecureString' \\"
echo "  --description 'Database password for ${PROJECT_NAME}' \\"
echo "  --overwrite"
echo ""

# List all parameters
echo "📋 Current parameters:"
aws ssm describe-parameters \
    --parameter-filters "Key=Name,Option=BeginsWith,Values=/${PROJECT_NAME}/${ENVIRONMENT}/" \
    --query 'Parameters[*].[Name,Type]' \
    --output table

echo ""
echo "✅ Setup complete! Don't forget to set the database password."