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
	// Full update for general editing, but scoped to ID and TenantID
	return r.db.Model(&domain.Item{}).Where("id = ? AND tenant_id = ?", item.ID, item.TenantID).Updates(item).Error
}

func (r *inventoryRepository) DeleteItem(id uint, tenantID uint) error {
	return r.db.Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&domain.Item{}).Error
}
