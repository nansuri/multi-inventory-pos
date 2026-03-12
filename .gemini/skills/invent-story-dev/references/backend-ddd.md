# Backend Development: Domain-Driven Design (DDD)

The Invent Story backend is organized into layers to separate concerns and ensure maintainability.

## 📁 Directory Structure
- `internal/domain/`: Core business entities and repository/usecase interfaces. **No external dependencies.**
- `internal/usecase/`: Pure business logic. Orchestrates entities and repositories.
- `internal/repository/`: Data persistence (GORM). Implements domain interfaces.
- `internal/delivery/http/`: Request handling (Gin). Entry point for the API.

## 🧱 Development Guardrails

### 1. Domain Entities
- Define entities in `internal/domain/*.go`.
- Use GORM tags for database mapping.
- All entities MUST have a `TenantID uint` for multi-tenant isolation.
- Interfaces for Repositories and Usecases live here.

### 2. Usecase (Business Logic)
- Implementation lives in `internal/usecase/*.go`.
- Usecases MUST NOT interact with the database directly; use Repository interfaces.
- Handle complex business rules (e.g., "Deduct inventory when cooking").
- **Multi-tenancy:** Always pass `tenantID` to repository methods.

### 3. Repository (Persistence)
- Implementation lives in `internal/repository/*.go`.
- Every query MUST include a filter for `tenant_id` to prevent data leaks.
- Example: `r.db.Where("tenant_id = ?", tenantID)...`

### 4. Delivery (HTTP)
- Implementation lives in `internal/delivery/http/*.go`.
- Responsible for request binding, validation, and JSON response formatting.
- Extract `tenant_id` and `user_id` from Gin context (set by middleware).

## 🚀 New Feature Workflow
1. Define the entity and interfaces in `domain/`.
2. Update `main.go` to include the new entity in `AutoMigrate`.
3. Implement the `repository`.
4. Implement the `usecase`.
5. Implement the `handler` and register routes in `main.go`.
