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

func (r *orderRepository) GetTableByID(id uint, branchID uint) (*domain.Table, error) {
	var table domain.Table
	err := r.db.Where("id = ? AND branch_id = ?", id, branchID).First(&table).Error
	return &table, err
}

func (r *orderRepository) ListTables(branchID uint) ([]domain.Table, error) {
	var tables []domain.Table
	err := r.db.Where("branch_id = ?", branchID).Order("table_number asc").Find(&tables).Error
	return tables, err
}

func (r *orderRepository) UpdateTable(table *domain.Table) error {
	return r.db.Save(table).Error
}

func (r *orderRepository) UpdateTableStatus(id uint, status domain.TableStatus) error {
	return r.db.Model(&domain.Table{}).Where("id = ?", id).Update("status", status).Error
}

func (r *orderRepository) DeleteTable(id uint, branchID uint) error {
	return r.db.Where("id = ? AND branch_id = ?", id, branchID).Delete(&domain.Table{}).Error
}

func (r *orderRepository) CreateOrder(order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) GetOrderByID(id uint, branchID uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Items.Product.Recipes.Ingredient").Where("id = ? AND branch_id = ?", id, branchID).First(&order).Error
	return &order, err
}

func (r *orderRepository) ListOrders(branchID uint) ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Preload("Table").Preload("Items.Product").Where("branch_id = ?", branchID).Order("created_at desc").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) UpdateOrderStatus(id uint, branchID uint, status domain.OrderStatus) error {
	return r.db.Model(&domain.Order{}).Where("id = ? AND branch_id = ?", id, branchID).Update("status", status).Error
}

func (r *orderRepository) UpdateProductStock(productID uint, newStock float64) error {
	return r.db.Model(&domain.Product{}).Where("id = ?", productID).Update("stock", newStock).Error
}
