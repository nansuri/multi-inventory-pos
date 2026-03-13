package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	TenantID  uint           `gorm:"not null;index" json:"tenant_id"`
	BranchID  uint           `gorm:"not null;index" json:"branch_id"`
	Branch    Branch         `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	Name      string         `gorm:"not null" json:"name"`
	Price     float64        `gorm:"type:decimal(10,2)" json:"price"`
	Stock     float64        `gorm:"type:decimal(10,2);default:0" json:"stock"`
	Recipes   []Recipe       `json:"recipes,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Recipe struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	ProductID    uint           `gorm:"not null;index" json:"product_id"`
	IngredientID uint           `gorm:"not null;index" json:"ingredient_id"` // Item ID
	Ingredient   Item           `gorm:"foreignKey:IngredientID" json:"ingredient,omitempty"`
	Quantity     float64        `gorm:"type:decimal(10,2);not null" json:"quantity"`
	Unit         string         `gorm:"not null" json:"unit"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type ProductionLog struct {
	ID          uint                     `gorm:"primaryKey" json:"id"`
	TenantID    uint                     `gorm:"not null;index" json:"tenant_id"`
	BranchID    uint                     `gorm:"not null;index" json:"branch_id"`
	Branch      Branch                   `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	ProductID   uint                     `gorm:"not null;index" json:"product_id"`
	Product     Product                  `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Pax         int                      `gorm:"not null" json:"pax"`
	Ingredients []ProductionIngredientLog `json:"ingredients,omitempty"`
	CreatedAt   time.Time                `json:"created_at"`
}

type ProductionIngredientLog struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	ProductionLogID uint    `gorm:"not null;index" json:"production_log_id"`
	IngredientName  string  `gorm:"not null" json:"ingredient_name"`
	Quantity        float64 `gorm:"type:decimal(10,2);not null" json:"quantity"`
	Unit            string  `gorm:"not null" json:"unit"`
}

type ProductRepository interface {
	CreateProduct(product *Product) error
	GetProductByID(id uint, branchID uint) (*Product, error)
	ListProducts(branchID uint) ([]Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id uint, branchID uint) error
	
	// Recipe operations
	AddRecipeItem(recipe *Recipe) error
	GetRecipesByProductID(productID uint) ([]Recipe, error)
	ClearRecipes(productID uint) error

	// Production Log
	CreateProductionLog(log *ProductionLog) error
	ListProductionLogs(branchID uint) ([]ProductionLog, error)
}

type ProductUsecase interface {
	CreateProduct(product *Product) error
	GetProductByID(id uint, branchID uint) (*Product, error)
	ListProducts(branchID uint) ([]Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id uint, branchID uint) error
	SetRecipe(productID uint, branchID uint, recipes []Recipe) error
	PrepareProduct(productID uint, branchID uint, pax int) error
	ListProductionLogs(branchID uint) ([]ProductionLog, error)
}
