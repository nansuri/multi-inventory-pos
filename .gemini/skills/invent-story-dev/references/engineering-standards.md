# Invent Story: Engineering Standards

Foundational mandates for all development on the Invent Story project.

## 🔒 Security & Data Integrity
- **Multi-Tenancy:** Isolation is critical. Every database operation MUST be scoped to a `tenant_id`. NEVER assume a global context.
- **Credential Protection:** NEVER log secrets or commit `.env` files. Use the `.env.example` as a template.
- **Authentication:** All business endpoints MUST be protected by the `AuthMiddleware`.

## 🛠️ Environment Configuration
- Use the unified root `.env` file for configuration.
- Backend settings are prefixed with `DB_`, `REDIS_`, `APP_`, etc.
- Frontend settings are prefixed with `VITE_`.

## 📦 Containerization
- Always keep `Dockerfile` and `docker-compose.yml` updated when adding new dependencies or services.
- Verify changes by running `docker-compose up --build`.

## 📝 Coding Style
- **Backend (Go):** Standard `gofmt`. Use clear, descriptive variable names. Favor explicit error handling.
- **Frontend (TS):** Strict typing. Use `interface` for data models. Follow the composition API pattern in Vue 3.
- **Documentation:** Update `README.md` or create a new plan in `plans/` when implementing cross-cutting changes.
