package usecase

import (
	"errors"

	"github.com/username/multi-inventory-manager/internal/domain"
)

type orderUsecase struct {
	repo          domain.OrderRepository
	inventoryRepo domain.InventoryRepository
}

func NewOrderUsecase(repo domain.OrderRepository, inventoryRepo domain.InventoryRepository) domain.OrderUsecase {
	return &orderUsecase{repo: repo, inventoryRepo: inventoryRepo}
}

func (u *orderUsecase) CreateTable(table *domain.Table) error {
	return u.repo.CreateTable(table)
}

func (u *orderUsecase) ListTables(tenantID uint) ([]domain.Table, error) {
	return u.repo.ListTables(tenantID)
}

func (u *orderUsecase) CreateOrder(order *domain.Order) error {
	// If it's a table order, mark table as occupied
	if order.TableID != nil {
		if err := u.repo.UpdateTableStatus(*order.TableID, domain.TableOccupied); err != nil {
			return err
		}
	}
	return u.repo.CreateOrder(order)
}

func (u *orderUsecase) GetOrderByID(id uint, tenantID uint) (*domain.Order, error) {
	return u.repo.GetOrderByID(id, tenantID)
}

func (u *orderUsecase) ListOrders(tenantID uint) ([]domain.Order, error) {
	return u.repo.ListOrders(tenantID)
}

func (u *orderUsecase) CompleteOrder(orderID uint, tenantID uint) error {
	// 1. Get Order with Items
	order, err := u.repo.GetOrderByID(orderID, tenantID)
	if err != nil {
		return err
	}

	if order.Status == domain.OrderCompleted {
		return errors.New("order already completed")
	}

	// 2. Deduct product stock for each item in the order
	for _, item := range order.Items {
		newStock := item.Product.Stock - float64(item.Quantity)
		if newStock < 0 {
			return errors.New("insufficient stock for product: " + item.Product.Name)
		}

		if err := u.repo.UpdateProductStock(item.ProductID, newStock); err != nil {
			return err
		}
	}

	// 3. Update order status
	if err := u.repo.UpdateOrderStatus(orderID, tenantID, domain.OrderCompleted); err != nil {
		return err
	}

	// 4. If table order, mark table as available
	if order.TableID != nil {
		if err := u.repo.UpdateTableStatus(*order.TableID, domain.TableAvailable); err != nil {
			return err
		}
	}

	return nil
}
