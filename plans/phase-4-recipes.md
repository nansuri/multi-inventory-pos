# Phase 4: Food & Recipe Management

This plan links food products to inventory through recipes.

## Objective
Define finished products (Food) and their ingredient requirements (Recipes).

## Key Files & Context
- `internal/domain/product.go`: Food entity.
- `internal/domain/recipe.go`: Recipe/Ingredient relationship.
- `internal/usecase/recipe_usecase.go`: Logic to calculate pax-based inventory impact.

## Implementation Steps
1. **Define Entities**:
   - `Product`: ID, TenantID, Name, Price.
   - `Recipe`: ID, ProductID, IngredientID (ItemID), Quantity, Unit.
2. **Setup Recipe Creation**:
   - Link multiple Items (ingredients) to one Product.
3. **Pax-based Preparation Logic**:
   - `POST /production/prepare`: Input ProductID and Pax (count).
   - Logic: Iterate over Recipe, calculate total needed quantity, and deduct from Item Stock.
4. **API Endpoints**:
   - `GET /products`: List available food items.
   - `POST /products/{id}/recipe`: Set ingredients for a food item.

## Verification & Testing
- Unit test for pax-based ingredient calculation.
- Transaction test: ensure inventory is only deducted if *all* ingredients have sufficient stock (or warn if not).
