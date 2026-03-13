package domain

import (
	"time"

	"gorm.io/gorm"
)

type TableStatus string

const (
	TableAvailable TableStatus = "available"
	TableOccupied  TableStatus = "occupied"
)

type Table struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	TenantID    uint           `gorm:"not null;index" json:"tenant_id"`
	BranchID    uint           `gorm:"not null;index" json:"branch_id"`
	Branch      Branch         `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	TableNumber string         `gorm:"not null" json:"table_number"`
	Capacity    int            `json:"capacity"`
	Status      TableStatus    `gorm:"type:varchar(20);default:'available'" json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type OrderStatus string

const (
	OrderPending   OrderStatus = "pending"
	OrderPaid      OrderStatus = "paid"
	OrderCompleted OrderStatus = "completed"
	OrderCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	TenantID        uint           `gorm:"not null;index" json:"tenant_id"`
	BranchID        uint           `gorm:"not null;index" json:"branch_id"`
	Branch          Branch         `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	CustomerName    string         `json:"customer_name"`
	TableID         *uint          `json:"table_id"` // Nullable for Delivery
	Table           *Table         `json:"table,omitempty"`
	Status          OrderStatus    `gorm:"type:varchar(20);default:'pending'" json:"status"`
	IsDelivery      bool           `gorm:"default:false" json:"is_delivery"`
	DeliveryAddress string         `json:"delivery_address,omitempty"`
	DeliveryContact string         `json:"delivery_contact,omitempty"`
	TotalPrice      float64        `gorm:"type:decimal(10,2)" json:"total_price"`
	Items           []OrderItem    `json:"items,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}


type OrderItem struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	OrderID   uint           `gorm:"not null;index" json:"order_id"`
	ProductID uint           `gorm:"not null;index" json:"product_id"`
	Product   Product        `json:"product,omitempty"`
	Quantity  int            `gorm:"not null" json:"quantity"`
	Price     float64        `gorm:"type:decimal(10,2)" json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type OrderRepository interface {
	CreateTable(table *Table) error
	GetTableByID(id uint, branchID uint) (*Table, error)
	ListTables(branchID uint) ([]Table, error)
	UpdateTable(table *Table) error
	UpdateTableStatus(id uint, status TableStatus) error
	DeleteTable(id uint, branchID uint) error

	CreateOrder(order *Order) error
	GetOrderByID(id uint, branchID uint) (*Order, error)
	ListOrders(branchID uint) ([]Order, error)
	UpdateOrderStatus(id uint, branchID uint, status OrderStatus) error
	UpdateProductStock(productID uint, newStock float64) error
}

type OrderUsecase interface {
	CreateTable(table *Table) error
	ListTables(branchID uint) ([]Table, error)
	UpdateTable(table *Table) error
	DeleteTable(id uint, branchID uint) error
	
	CreateOrder(order *Order) error
	GetOrderByID(id uint, branchID uint) (*Order, error)
	ListOrders(branchID uint) ([]Order, error)
	CompleteOrder(orderID uint, branchID uint) error
}
