package domain

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	TenantID      uint           `gorm:"not null;index" json:"tenant_id"`
	Name          string         `gorm:"not null" json:"name"`
	Barcode       string         `gorm:"index" json:"barcode"`
	SKU           string         `gorm:"index" json:"sku"`
	Price         float64        `gorm:"type:decimal(10,2)" json:"price"`
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
	GetItemByID(id uint, tenantID uint) (*Item, error)
	GetItemByBarcode(barcode string, tenantID uint) (*Item, error)
	ListItems(tenantID uint) ([]Item, error)
	UpdateItem(item *Item) error
	DeleteItem(id uint, tenantID uint) error
}

type InventoryUsecase interface {
	CreateItem(item *Item) error
	GetItemByID(id uint, tenantID uint) (*Item, error)
	GetItemByBarcode(barcode string, tenantID uint) (*Item, error)
	ListItems(tenantID uint) ([]Item, error)
	UpdateItem(item *Item) error
	UpdateStock(itemID uint, tenantID uint, quantity float64) error
	DeleteItem(id uint, tenantID uint) error
	ToggleActive(id uint, tenantID uint) error
}
