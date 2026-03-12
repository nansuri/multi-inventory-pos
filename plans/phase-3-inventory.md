# Phase 3: Stock/Inventory Management

This plan focuses on managing basic ingredients (items) and their stock levels.

## Objective
Enable item registration (Barcode/QR) and manual/automated stock updates.

## Key Files & Context
- `internal/domain/item.go`: Item entity.
- `internal/usecase/inventory_usecase.go`: Stock tracking logic.

## Implementation Steps
1. **Define Item Entity**:
   - `Item`: ID, TenantID, Name, Barcode/QR, SKU, Price, CurrentStock, MinStockAlert, Unit (KG, GR, Piece).
2. **Setup Scanning Feature**:
   - Backend endpoint to search/register items by barcode.
3. **Inventory Endpoints**:
   - `POST /inventory/items`: Create a new ingredient.
   - `GET /inventory/items`: List all items for the tenant.
   - `PATCH /inventory/items/{id}/stock`: Manually add or subtract stock.
4. **Barcode/QR Integration**:
   - Ensure the database schema supports unique barcodes per tenant.

## Verification & Testing
- Test barcode search with sample codes.
- Validate stock updates: ensure stock cannot go below zero (or handle negative stock if business allows).
- Confirm low stock notification logic works when `CurrentStock` < `MinStockAlert`.
