package usecase

import (
	"errors"

	"github.com/username/multi-inventory-manager/internal/domain"
)

type productUsecase struct {
	repo          domain.ProductRepository
	inventoryRepo domain.InventoryRepository
}

func NewProductUsecase(repo domain.ProductRepository, inventoryRepo domain.InventoryRepository) domain.ProductUsecase {
	return &productUsecase{repo: repo, inventoryRepo: inventoryRepo}
}

func (u *productUsecase) CreateProduct(product *domain.Product) error {
	return u.repo.CreateProduct(product)
}

func (u *productUsecase) GetProductByID(id uint, tenantID uint) (*domain.Product, error) {
	return u.repo.GetProductByID(id, tenantID)
}

func (u *productUsecase) ListProducts(tenantID uint) ([]domain.Product, error) {
	return u.repo.ListProducts(tenantID)
}

func (u *productUsecase) SetRecipe(productID uint, tenantID uint, recipes []domain.Recipe) error {
	// Check if product belongs to tenant
	_, err := u.repo.GetProductByID(productID, tenantID)
	if err != nil {
		return err
	}

	// Clear existing recipes
	if err := u.repo.ClearRecipes(productID); err != nil {
		return err
	}

	// Add new recipe items
	for _, r := range recipes {
		r.ProductID = productID
		if err := u.repo.AddRecipeItem(&r); err != nil {
			return err
		}
	}

	return nil
}

func (u *productUsecase) PrepareProduct(productID uint, tenantID uint, pax int) error {
	// 1. Get Product and Recipes
	product, err := u.repo.GetProductByID(productID, tenantID)
	if err != nil {
		return err
	}

	if len(product.Recipes) == 0 {
		return errors.New("product has no recipe defined")
	}

	// 2. Check if sufficient stock for ALL ingredients
	for _, recipe := range product.Recipes {
		totalNeeded := recipe.Quantity * float64(pax)
		if recipe.Ingredient.CurrentStock < totalNeeded {
			return errors.New("insufficient stock for ingredient: " + recipe.Ingredient.Name)
		}
	}

	// 3. Deduct stock
	for _, recipe := range product.Recipes {
		totalNeeded := recipe.Quantity * float64(pax)
		err := u.inventoryRepo.UpdateItem(&domain.Item{
			ID:           recipe.IngredientID,
			TenantID:     tenantID,
			CurrentStock: recipe.Ingredient.CurrentStock - totalNeeded,
		})
		// Note: The simple UpdateItem in inventoryRepo might need a specific UpdateStock method to be safe
		// Let's assume we use UpdateItem for now, but in real scenarios, use Atomic transactions.
		if err != nil {
			return err
		}
	}

	// 4. Increase Product stock
	product.Stock += float64(pax)
	if err := u.repo.UpdateProduct(product); err != nil {
		return err
	}

	// 5. Create log
	return u.repo.CreatePreparationLog(&domain.PreparationLog{
		TenantID:  tenantID,
		ProductID: productID,
		Pax:       pax,
	})
}

func (u *productUsecase) ListPreparationLogs(tenantID uint) ([]domain.PreparationLog, error) {
	return u.repo.ListPreparationLogs(tenantID)
}
