package repository

import (
	"github.com/username/multi-inventory-manager/internal/domain"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetProductByID(id uint, branchID uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.Preload("Recipes.Ingredient").Where("id = ? AND branch_id = ?", id, branchID).First(&product).Error
	return &product, err
}

func (r *productRepository) ListProducts(branchID uint) ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Preload("Recipes.Ingredient").Where("branch_id = ?", branchID).Find(&products).Error
	return products, err
}

func (r *productRepository) UpdateProduct(product *domain.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) DeleteProduct(id uint, branchID uint) error {
	return r.db.Where("id = ? AND branch_id = ?", id, branchID).Delete(&domain.Product{}).Error
}

func (r *productRepository) AddRecipeItem(recipe *domain.Recipe) error {
	return r.db.Create(recipe).Error
}

func (r *productRepository) GetRecipesByProductID(productID uint) ([]domain.Recipe, error) {
	var recipes []domain.Recipe
	err := r.db.Preload("Ingredient").Where("product_id = ?", productID).Find(&recipes).Error
	return recipes, err
}

func (r *productRepository) ClearRecipes(productID uint) error {
	return r.db.Where("product_id = ?", productID).Delete(&domain.Recipe{}).Error
}

func (r *productRepository) CreateProductionLog(log *domain.ProductionLog) error {
	return r.db.Create(log).Error
}

func (r *productRepository) ListProductionLogs(branchID uint) ([]domain.ProductionLog, error) {
	var logs []domain.ProductionLog
	err := r.db.Preload("Product").Preload("Ingredients").Where("branch_id = ?", branchID).Order("created_at desc").Find(&logs).Error
	return logs, err
}
