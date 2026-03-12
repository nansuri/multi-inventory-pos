package usecase

import (
	"errors"

	"github.com/username/multi-inventory-manager/internal/domain"
)

type inventoryUsecase struct {
	repo domain.InventoryRepository
}

func NewInventoryUsecase(repo domain.InventoryRepository) domain.InventoryUsecase {
	return &inventoryUsecase{repo: repo}
}

func (u *inventoryUsecase) CreateItem(item *domain.Item) error {
	return u.repo.CreateItem(item)
}

func (u *inventoryUsecase) GetItemByID(id uint, tenantID uint) (*domain.Item, error) {
	return u.repo.GetItemByID(id, tenantID)
}

func (u *inventoryUsecase) GetItemByBarcode(barcode string, tenantID uint) (*domain.Item, error) {
	return u.repo.GetItemByBarcode(barcode, tenantID)
}

func (u *inventoryUsecase) ListItems(tenantID uint) ([]domain.Item, error) {
	return u.repo.ListItems(tenantID)
}

func (u *inventoryUsecase) UpdateItem(item *domain.Item) error {
	return u.repo.UpdateItem(item)
}

func (u *inventoryUsecase) UpdateStock(itemID uint, tenantID uint, quantity float64) error {
	item, err := u.repo.GetItemByID(itemID, tenantID)
	if err != nil {
		return err
	}

	item.CurrentStock += quantity
	if item.CurrentStock < 0 {
		return errors.New("insufficient stock")
	}

	// We can use a surgical map update or just pass the object
	return u.repo.UpdateItem(&domain.Item{
		ID:           itemID,
		TenantID:     tenantID,
		CurrentStock: item.CurrentStock,
	})
}

func (u *inventoryUsecase) DeleteItem(id uint, tenantID uint) error {
	return u.repo.DeleteItem(id, tenantID)
}

func (u *inventoryUsecase) ToggleActive(id uint, tenantID uint) error {
	item, err := u.repo.GetItemByID(id, tenantID)
	if err != nil {
		return err
	}

	return u.repo.UpdateItem(&domain.Item{
		ID:       id,
		TenantID: tenantID,
		IsActive: !item.IsActive,
	})
}
