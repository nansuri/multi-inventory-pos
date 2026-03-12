package repository

import (
	"github.com/username/multi-inventory-manager/internal/domain"
	"gorm.io/gorm"
)

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) domain.InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) CreateItem(item *domain.Item) error {
	return r.db.Create(item).Error
}

func (r *inventoryRepository) GetItemByID(id uint, tenantID uint) (*domain.Item, error) {
	var item domain.Item
	err := r.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&item).Error
	return &item, err
}

func (r *inventoryRepository) GetItemByBarcode(barcode string, tenantID uint) (*domain.Item, error) {
	var item domain.Item
	err := r.db.Where("barcode = ? AND tenant_id = ?", barcode, tenantID).First(&item).Error
	return &item, err
}

func (r *inventoryRepository) ListItems(tenantID uint) ([]domain.Item, error) {
	var items []domain.Item
	err := r.db.Where("tenant_id = ?", tenantID).Find(&items).Error
	return items, err
}

func (r *inventoryRepository) UpdateItem(item *domain.Item) error {
	// If only updating stock, we can be more surgical
	return r.db.Model(item).Updates(map[string]interface{}{
		"current_stock": item.CurrentStock,
	}).Error
}
