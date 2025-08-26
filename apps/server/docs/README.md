# Chess Game Server

A scalable Go server for the roguelike chess game using the Gin framework.

## Quick Start

```bash
# Install dependencies
go mod tidy

# Run in development mode
go run main.go

# Build for production
go build -o bin/server main.go
```

## Project Structure

```
apps/server/
├── main.go                    # Application entry point
├── cmd/server/               # Server startup logic
├── internal/
│   ├── api/
│   │   ├── handlers/         # HTTP request handlers
│   │   ├── middleware/       # Custom middleware
│   │   └── routes/          # Route definitions
│   ├── config/              # Configuration management
│   ├── models/              # Data models
│   └── services/            # Business logic services
├── pkg/                     # Reusable packages
└── docs/                    # Documentation
```

## API Endpoints

### Health Check
- `GET /api/v1/health` - Server health status

### Future Endpoints
- `/api/v1/auth/*` - Authentication routes
- `/api/v1/game/*` - Game management routes
- `/api/v1/shop/*` - Shop and upgrades routes
- `/api/v1/player/*` - Player management routes

## Environment Variables

Copy `.env.example` to `.env` and configure:

- `ENV` - Environment (development/production)
- `PORT` - Server port (default: 8080)
- `ALLOWED_ORIGINS` - CORS allowed origins
- `DATABASE_URL` - Database connection string

## Development

The server uses:
- **Gin** for HTTP routing and middleware
- **CORS** for cross-origin requests
- **godotenv** for environment configuration

## Integration with Turbo

This server integrates with the monorepo's Turbo setup:
- `pnpm dev` - Start development server
- `pnpm build` - Build the server
- `pnpm lint` - Run Go formatting checks