---
name: invent-story-dev
description: Specialized development workflows and guardrails for the Invent Story project. Use when adding new features, refactoring existing logic, or managing the Go DDD backend and Vue 3 frontend.
---

# Invent Story Development

This skill provides the necessary context and procedural knowledge to develop the Invent Story multi-tenant system effectively.

## 🏁 Quick Start

To begin developing a new feature:
1. Identify the domain and scope.
2. Read the appropriate reference module below.
3. Create a plan in the `plans/` directory.

## 📚 Reference Modules

- **Backend (Go/DDD)**: See [references/backend-ddd.md](references/backend-ddd.md) for architecture, multi-tenancy, and GORM patterns.
- **Frontend (Vue 3)**: See [references/frontend-best-practices.md](references/frontend-best-practices.md) for component design, Pinia, i18n, and styling.
- **Engineering Standards**: See [references/engineering-standards.md](references/engineering-standards.md) for security mandates, environment config, and general coding style.

## 🛠️ Core Workflows

### Adding a New Data Model
1. Define the entity in `backend/internal/domain/`.
2. Add to `AutoMigrate` in `backend/cmd/api/main.go`.
3. Implement Repository, Usecase, and Handler.
4. Add frontend views and routes.

### Local Development
- **Backend**: `cd backend && go run cmd/api/main.go`
- **Frontend**: `cd frontend && npm run dev`
- **Infrastructure**: `docker-compose up -d db redis`
