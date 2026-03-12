package repository

import (
	"github.com/username/multi-inventory-manager/internal/domain"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateTable(table *domain.Table) error {
	return r.db.Create(table).Error
}

func (r *orderRepository) GetTableByID(id uint, tenantID uint) (*domain.Table, error) {
	var table domain.Table
	err := r.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&table).Error
	return &table, err
}

func (r *orderRepository) ListTables(tenantID uint) ([]domain.Table, error) {
	var tables []domain.Table
	err := r.db.Where("tenant_id = ?", tenantID).Order("table_number asc").Find(&tables).Error
	return tables, err
}

func (r *orderRepository) UpdateTableStatus(id uint, status domain.TableStatus) error {
	return r.db.Model(&domain.Table{}).Where("id = ?", id).Update("status", status).Error
}

func (r *orderRepository) CreateOrder(order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) GetOrderByID(id uint, tenantID uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Items.Product.Recipes.Ingredient").Where("id = ? AND tenant_id = ?", id, tenantID).First(&order).Error
	return &order, err
}

func (r *orderRepository) ListOrders(tenantID uint) ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Preload("Table").Preload("Items.Product").Where("tenant_id = ?", tenantID).Order("created_at desc").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) UpdateOrderStatus(id uint, tenantID uint, status domain.OrderStatus) error {
	return r.db.Model(&domain.Order{}).Where("id = ? AND tenant_id = ?", id, tenantID).Update("status", status).Error
}

func (r *orderRepository) UpdateProductStock(productID uint, newStock float64) error {
	return r.db.Model(&domain.Product{}).Where("id = ?", productID).Update("stock", newStock).Error
}
