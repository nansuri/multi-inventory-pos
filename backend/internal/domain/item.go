package domain

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	TenantID      uint           `gorm:"not null;index" json:"tenant_id"`
	BranchID      uint           `gorm:"not null;index" json:"branch_id"`
	Branch        Branch         `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	Name          string         `gorm:"not null" json:"name"`
	Barcode       string         `gorm:"index" json:"barcode"`
	SKU           string         `gorm:"index" json:"sku"`
	CostPrice     float64        `gorm:"column:cost_price;type:decimal(10,2)" json:"cost_price"`
	CurrentStock  float64        `gorm:"type:decimal(10,2);default:0" json:"current_stock"`
	MinStockAlert float64        `gorm:"type:decimal(10,2);default:0" json:"min_stock_alert"`
	Unit          string         `gorm:"not null" json:"unit"` // KG, GR, Piece, etc.
	IsActive      bool           `gorm:"default:true" json:"is_active"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type InventoryRepository interface {
	CreateItem(item *Item) error
	GetItemByID(id uint, branchID uint) (*Item, error)
	GetItemByBarcode(barcode string, branchID uint) (*Item, error)
	ListItems(branchID uint) ([]Item, error)
	UpdateItem(item *Item) error
	UpdateStock(id uint, branchID uint, newStock float64) error
	UpdateField(id uint, branchID uint, field string, value interface{}) error
	DeleteItem(id uint, branchID uint) error
}

type InventoryUsecase interface {
	CreateItem(item *Item) error
	GetItemByID(id uint, branchID uint) (*Item, error)
	GetItemByBarcode(barcode string, branchID uint) (*Item, error)
	ListItems(branchID uint) ([]Item, error)
	UpdateItem(item *Item) error
	UpdateStock(itemID uint, branchID uint, quantity float64) error
	DeleteItem(id uint, branchID uint) error
	ToggleActive(id uint, branchID uint) error
}
