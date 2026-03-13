# GEMINI Mandates: Invent Story

This document defines the core architectural principles, engineering standards, and mandatory workflows for the Invent Story project. These mandates take absolute precedence over general defaults.

## 🏛 Architectural Core

### 1. Multi-Tenancy Hierarchy
- **Structure**: `Tenant` > `Branch` > `Operational Data`.
- **Isolation**: Every operational entity (Inventory, Products, Orders, etc.) **must** be scoped by `BranchID`.
- **Context**: The backend uses `BranchMiddleware` to extract `X-Branch-ID` from headers. Repositories must explicitly filter by this ID.

### 2. Backend (Go/DDD)
- **Pattern**: Domain-Driven Design.
- **Layers**: Domain (Interfaces/Entities) -> Repository (GORM) -> Usecase (Logic) -> Delivery (Gin).
- **Validation**: All repository methods must accept a `branchID` or `tenantID` to ensure data leakage prevention.

### 3. Frontend (Vue 3 / Tailwind v4)
- **State**: Pinia (`authStore`, `configStore`).
- **Styling**: Tailwind CSS v4. Use `@variant dark (&:where(.dark, .dark *))` for dark mode.
- **Interceptors**: Axios must automatically attach `Authorization` and `X-Branch-ID` headers.
- **Reactivity**: Use component keys (e.g., `mainKey`) to force re-renders when switching branch context rather than full page reloads.

## 🛠 Mandatory Workflows

### 1. Verification & Dry Runs
- **Mandatory**: Before declaring any task complete, you **must** run:
  - **Backend**: `cd backend && go build ./cmd/api/main.go`
  - **Frontend**: `cd frontend && npm run build`
- **Zero Warnings**: Production builds must succeed without TypeScript errors or unused variable warnings.

### 2. Internationalization (i18n)
- All user-facing strings must use `t('key')`.
- Updates to `en.ts` must be mirrored in `id.ts` and `ja.ts`.

### 3. Database Migrations
- Use GORM `AutoMigrate`. 
- New models must be registered in `backend/cmd/api/main.go`.

## 🔒 Security & Performance
- **Secrets**: Never log or commit `.env` values.
- **Query Efficiency**: Always use `Preload` for relationships to avoid N+1 queries in the repository layer.
- **Asset Usage**: Prefer Lucide icons and CSS-based glassmorphism over external image assets for UI components.
