# Phase 1: Infrastructure & Project Setup

This plan covers the initialization of the backend and common infrastructure.

## Objective
Establish a solid foundation using Go (DDD) and Docker.

## Key Files & Context
- `go.mod`: Dependency management.
- `cmd/api/main.go`: Entry point for the server.
- `docker-compose.yml`: Infrastructure (Postgres, Redis).
- `.env`: Unified configuration.

## Implementation Steps
1. **Initialize Go Module**: `go mod init github.com/username/multi-inventory-manager`.
2. **Setup Project Directory Structure (DDD)**:
   - `cmd/`: Application entry points.
   - `internal/domain/`: Core entities and interfaces.
   - `internal/usecase/`: Business logic.
   - `internal/repository/`: Data persistence (GORM).
   - `internal/delivery/http/`: HTTP handlers and routing.
   - `pkg/`: Reusable utilities.
3. **Infrastructure Configuration**:
   - Create `docker-compose.yml` with PostgreSQL and Redis services.
   - Setup `config/` package to load environmental variables.
4. **Database Connectivity**:
   - Implement GORM connection with auto-migration.
   - Setup Redis client for caching/sessions.

## Verification & Testing
- Run `docker-compose up` to ensure services are running.
- Verify DB connection during application startup.
- Check if `.env` values are correctly loaded.
