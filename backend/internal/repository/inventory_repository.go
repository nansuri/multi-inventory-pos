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

func (r *inventoryRepository) GetItemByID(id uint, branchID uint) (*domain.Item, error) {
	var item domain.Item
	err := r.db.Where("id = ? AND branch_id = ?", id, branchID).First(&item).Error
	return &item, err
}

func (r *inventoryRepository) GetItemByBarcode(barcode string, branchID uint) (*domain.Item, error) {
	var item domain.Item
	err := r.db.Where("barcode = ? AND branch_id = ?", barcode, branchID).First(&item).Error
	return &item, err
}

func (r *inventoryRepository) ListItems(branchID uint) ([]domain.Item, error) {
	var items []domain.Item
	err := r.db.Where("branch_id = ?", branchID).Find(&items).Error
	return items, err
}

func (r *inventoryRepository) UpdateItem(item *domain.Item) error {
	return r.db.Model(&domain.Item{}).Where("id = ? AND branch_id = ?", item.ID, item.BranchID).Updates(item).Error
}

func (r *inventoryRepository) UpdateStock(id uint, branchID uint, newStock float64) error {
	return r.db.Model(&domain.Item{}).Where("id = ? AND branch_id = ?", id, branchID).Update("current_stock", newStock).Error
}

func (r *inventoryRepository) UpdateField(id uint, branchID uint, field string, value interface{}) error {
	return r.db.Model(&domain.Item{}).Where("id = ? AND branch_id = ?", id, branchID).Update(field, value).Error
}

func (r *inventoryRepository) DeleteItem(id uint, branchID uint) error {
	return r.db.Where("id = ? AND branch_id = ?", id, branchID).Delete(&domain.Item{}).Error
}
