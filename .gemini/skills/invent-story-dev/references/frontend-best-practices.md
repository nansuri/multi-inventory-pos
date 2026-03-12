# Frontend Development: Best Practices

The Invent Story frontend is a modern Vue 3 application focused on a clean, component-based architecture.

## 🏗️ Architecture
- `src/api/`: Axios instances and API service definitions.
- `src/components/`: Reusable UI components.
- `src/layouts/`: Global page layouts (e.g., DashboardLayout).
- `src/stores/`: Pinia stores for state management (auth, config, etc.).
- `src/views/`: Main page components.
- `src/locales/`: Translation files (JSON/TS).

## 🧱 Development Guardrails

### 1. Component-First Approach
- Keep views lean. Move reusable UI logic into `src/components/`.
- Use **Tailwind CSS 4** for styling. Adhere to the `rounded-3xl` and `shadow-sm` design language.
- Use **Lucide Vue Next** for consistent iconography.

### 2. State Management (Pinia)
- Use Pinia stores for any state that spans multiple components.
- `auth.ts`: Handles JWT, user profile, and session persistence.
- `config.ts`: Handles global settings (language, currency formatting).

### 3. Internationalization (i18n)
- NEVER hardcode strings in templates.
- Use `{{ $t('key.path') }}` in templates or `const { t } = useI18n()` in scripts.
- Add new keys to `src/locales/en.ts`, `id.ts`, and `ja.ts` simultaneously.

### 4. Currency & Formatting
- Use `configStore.formatCurrency(amount)` for all price displays.
- This ensures currency symbols and formatting match the tenant's settings.

### 5. API Integration
- Use the shared `Axios` instance in `src/api/axios.ts`.
- The interceptor automatically attaches the Bearer token.
- Handle loading states and errors gracefully using professional modals or alerts.

## 🚀 New View Workflow
1. Create the `.vue` file in `src/views/`.
2. Add the route to `src/router/index.ts` (usually under `meta: { requiresAuth: true }`).
3. Add the navigation item to `DashboardLayout.vue`.
4. Define any necessary state in a Pinia store.
