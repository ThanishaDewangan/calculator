# Important Notes

## SQLC Code Generation

**IMPORTANT**: Before building or running the application, you must generate the SQLC code:

```bash
sqlc generate
```

This will create the `db/sqlc/` directory with the generated database access code.

The code will not compile until SQLC code is generated because the repository layer depends on it.

## Build Order

1. Install dependencies: `go mod tidy`
2. Generate SQLC code: `sqlc generate`
3. Build: `go build ./cmd/server`
4. Run: `go run ./cmd/server`

## Docker Build

The Dockerfile automatically generates SQLC code during the build process, so Docker builds will work without manual SQLC generation.

## Database Setup

Make sure PostgreSQL is running and the database exists before generating SQLC code, as SQLC validates the schema against the database.

