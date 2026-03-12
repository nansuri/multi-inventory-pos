# Phase 5: Order & Table Management

This plan implements front-of-house operations: orders and seating.

## Objective
Manage restaurant tables, handle customer orders, and integrate with inventory/delivery.

## Key Files & Context
- `internal/domain/table.go`: Table entity.
- `internal/domain/order.go`: Order entity.
- `internal/usecase/order_usecase.go`: Business logic for orders.

## Implementation Steps
1. **Table Management**:
   - `Table`: ID, TenantID, TableNumber, Capacity, Status (Available/Occupied).
   - CRUD for tables per tenant.
2. **Order Management**:
   - `Order`: ID, TenantID, TableID (null for delivery), OrderStatus (Pending/Paid/Completed), TotalPrice.
   - `OrderItem`: OrderID, ProductID, Quantity, Price.
3. **Delivery Integration**:
   - Add flag to `Order` for delivery.
   - Store delivery address and tracking info.
4. **Order Impact**:
   - When an order is "Completed", trigger ingredient deduction from inventory (similar to "Prepare" logic in Phase 4).
5. **API Endpoints**:
   - `GET /tables`: List all tables.
   - `POST /orders`: Place a new order for a table or delivery.
   - `PATCH /orders/{id}/pay`: Process payment and complete order.

## Verification & Testing
- Integration test for "Order -> Complete -> Inventory Deduction".
- Concurrent test: multiple orders for the same table.
