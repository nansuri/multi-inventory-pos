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

func (r *productRepository) GetProductByID(id uint, tenantID uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.Preload("Recipes.Ingredient").Where("id = ? AND tenant_id = ?", id, tenantID).First(&product).Error
	return &product, err
}

func (r *productRepository) ListProducts(tenantID uint) ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Preload("Recipes.Ingredient").Where("tenant_id = ?", tenantID).Find(&products).Error
	return products, err
}

func (r *productRepository) UpdateProduct(product *domain.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) DeleteProduct(id uint, tenantID uint) error {
	return r.db.Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&domain.Product{}).Error
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

func (r *productRepository) ListProductionLogs(tenantID uint) ([]domain.ProductionLog, error) {
	var logs []domain.ProductionLog
	err := r.db.Preload("Product").Preload("Ingredients").Where("tenant_id = ?", tenantID).Order("created_at desc").Find(&logs).Error
	return logs, err
}
