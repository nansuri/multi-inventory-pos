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

func (u *inventoryUsecase) GetItemByID(id uint, branchID uint) (*domain.Item, error) {
	return u.repo.GetItemByID(id, branchID)
}

func (u *inventoryUsecase) GetItemByBarcode(barcode string, branchID uint) (*domain.Item, error) {
	return u.repo.GetItemByBarcode(barcode, branchID)
}

func (u *inventoryUsecase) ListItems(branchID uint) ([]domain.Item, error) {
	return u.repo.ListItems(branchID)
}

func (u *inventoryUsecase) UpdateItem(item *domain.Item) error {
	return u.repo.UpdateItem(item)
}

func (u *inventoryUsecase) UpdateStock(itemID uint, branchID uint, quantity float64) error {
	item, err := u.repo.GetItemByID(itemID, branchID)
	if err != nil {
		return err
	}

	newStock := item.CurrentStock + quantity
	if newStock < 0 {
		return errors.New("insufficient stock")
	}

	return u.repo.UpdateStock(itemID, branchID, newStock)
}

func (u *inventoryUsecase) DeleteItem(id uint, branchID uint) error {
	return u.repo.DeleteItem(id, branchID)
}

func (u *inventoryUsecase) ToggleActive(id uint, branchID uint) error {
	item, err := u.repo.GetItemByID(id, branchID)
	if err != nil {
		return err
	}

	return u.repo.UpdateField(id, branchID, "is_active", !item.IsActive)
}
