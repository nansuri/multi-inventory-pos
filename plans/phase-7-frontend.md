# Phase 7: Frontend Development (VueJS)

This plan covers the client-side implementation.

## Objective
Build a responsive dashboard using Vue 3 and Vite.

## Key Files & Context
- `frontend/`: Vue project root.
- `frontend/src/store/`: Pinia state management.
- `frontend/src/views/`: Main dashboard components.

## Implementation Steps
1. **Initial Setup**:
   - `npm create vite@latest frontend -- --template vue`.
   - Install dependencies: `vue-router`, `pinia`, `axios`, `tailwindcss`.
2. **Authentication Flow**:
   - Register/Login views.
   - Axios interceptor to add JWT to requests.
   - Auth guard for routes.
3. **Module Views**:
   - **Inventory**: Table listing items, scan barcode button (using camera API).
   - **Recipes**: Product list with ingredient details.
   - **POS (Orders)**: Table grid for ordering.
   - **Dashboard**: Charts and alerts.
4. **Integration**:
   - Connect all views to Backend API.
   - Unified error handling.

## Verification & Testing
- Component testing (Vue Test Utils).
- End-to-end flow: Login -> Register Item -> Create Product -> Order Product -> Check Dashboard.
