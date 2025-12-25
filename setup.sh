#!/bin/bash

# Setup script for Linux/Mac

echo "Setting up Go Backend Project..."

# Check if Go is installed
echo -e "\nChecking Go installation..."
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed or not in PATH"
    exit 1
fi
echo "Go found: $(go version)"

# Install dependencies
echo -e "\nInstalling Go dependencies..."
go mod tidy
if [ $? -ne 0 ]; then
    echo "Error: Failed to install dependencies"
    exit 1
fi
echo "Dependencies installed successfully"

# Check if sqlc is installed
echo -e "\nChecking SQLC installation..."
if ! command -v sqlc &> /dev/null; then
    echo "SQLC not found. Installing..."
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    if [ $? -ne 0 ]; then
        echo "Error: Failed to install SQLC"
        exit 1
    fi
    echo "SQLC installed successfully"
else
    echo "SQLC found: $(sqlc version)"
fi

# Generate SQLC code
echo -e "\nGenerating SQLC code..."
sqlc generate
if [ $? -ne 0 ]; then
    echo "Warning: SQLC generation failed. Make sure database migrations are set up."
    echo "You can generate SQLC code later with: sqlc generate"
else
    echo "SQLC code generated successfully"
fi

echo -e "\nSetup complete!"
echo -e "\nNext steps:"
echo "1. Set up PostgreSQL database"
echo "2. Run migrations: psql -U postgres -d users_db -f db/migrations/0001_create_users.sql"
echo "3. Set environment variables (see README.md)"
echo "4. Run the server: go run ./cmd/server"


