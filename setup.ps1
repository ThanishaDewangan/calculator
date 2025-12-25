# Setup script for Windows PowerShell

Write-Host "Setting up Go Backend Project..." -ForegroundColor Green

# Check if Go is installed
Write-Host "`nChecking Go installation..." -ForegroundColor Yellow
$goVersion = go version
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Go is not installed or not in PATH" -ForegroundColor Red
    exit 1
}
Write-Host "Go found: $goVersion" -ForegroundColor Green

# Install dependencies
Write-Host "`nInstalling Go dependencies..." -ForegroundColor Yellow
go mod tidy
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to install dependencies" -ForegroundColor Red
    exit 1
}
Write-Host "Dependencies installed successfully" -ForegroundColor Green

# Check if sqlc is installed
Write-Host "`nChecking SQLC installation..." -ForegroundColor Yellow
$sqlcVersion = sqlc version
if ($LASTEXITCODE -ne 0) {
    Write-Host "SQLC not found. Installing..." -ForegroundColor Yellow
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    if ($LASTEXITCODE -ne 0) {
        Write-Host "Error: Failed to install SQLC" -ForegroundColor Red
        exit 1
    }
    Write-Host "SQLC installed successfully" -ForegroundColor Green
} else {
    Write-Host "SQLC found: $sqlcVersion" -ForegroundColor Green
}

# Generate SQLC code
Write-Host "`nGenerating SQLC code..." -ForegroundColor Yellow
sqlc generate
if ($LASTEXITCODE -ne 0) {
    Write-Host "Warning: SQLC generation failed. Make sure database migrations are set up." -ForegroundColor Yellow
    Write-Host "You can generate SQLC code later with: sqlc generate" -ForegroundColor Yellow
} else {
    Write-Host "SQLC code generated successfully" -ForegroundColor Green
}

Write-Host "`nSetup complete!" -ForegroundColor Green
Write-Host "`nNext steps:" -ForegroundColor Cyan
Write-Host "1. Set up PostgreSQL database"
Write-Host "2. Run migrations: psql -U postgres -d users_db -f db/migrations/0001_create_users.sql"
Write-Host "3. Set environment variables (see README.md)"
Write-Host "4. Run the server: go run ./cmd/server"


