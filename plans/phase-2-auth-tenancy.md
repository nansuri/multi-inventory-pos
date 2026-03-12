# Phase 2: Tenant & User Management

This plan focuses on core multi-tenancy and user security.

## Objective
Implement authentication, RBAC, and multi-tenant data isolation.

## Key Files & Context
- `internal/domain/tenant.go`: Tenant entity.
- `internal/domain/user.go`: User entity (Owner/Employee).
- `internal/delivery/http/middleware/auth.go`: JWT & Multi-tenant middleware.

## Implementation Steps
1. **Define Core Entities**:
   - `Tenant`: ID, Name, Currency, etc.
   - `User`: ID, TenantID, Username, PasswordHash, Role (Owner, Employee).
2. **Setup Authentication Layer**:
   - JWT integration for token generation and validation.
   - Password hashing (bcrypt).
3. **Multi-tenant Logic**:
   - Add `TenantID` to all relevant database models.
   - Implement GORM global scope or middleware to inject `TenantID` in every query for isolation.
4. **API Endpoints**:
   - `POST /auth/register`: Create a new tenant and owner user.
   - `POST /auth/login`: Authenticate user and return JWT.
   - `GET /tenant/me`: Get current tenant settings.

## Verification & Testing
- Unit tests for password hashing and JWT signing.
- Integration tests to ensure data from Tenant A is not accessible by Tenant B.
- Test RBAC: Ensure Employees cannot access Owner-only settings.
