# Phase 6: Dashboard & Analytics

This plan focuses on high-level reporting and business health.

## Objective
Provide insights into stock status, revenue, and money movement.

## Key Files & Context
- `internal/usecase/report_usecase.go`: Summary logic.
- `internal/delivery/http/dashboard_handler.go`: UI-friendly endpoints.

## Implementation Steps
1. **Summary Metrics**:
   - Total items low in stock (below `MinStockAlert`).
   - Daily/Monthly revenue (from `Orders`).
   - Recent activity log (inventory added/deducted).
2. **Money Movement**:
   - `Transaction`: ID, TenantID, Type (In/Out), Amount, Description, Category (Stock Purchase, Sales).
   - Track stock purchases (manual entries or linked to Item registration).
3. **API Endpoints**:
   - `GET /dashboard/summary`: High-level numbers for cards.
   - `GET /dashboard/stock-alerts`: List of items to restock.
   - `GET /reports/revenue`: Revenue over time (weekly/monthly).

## Verification & Testing
- Mock data generation to test dashboard visualization logic.
- Verify "Total Stock Value" calculation.
