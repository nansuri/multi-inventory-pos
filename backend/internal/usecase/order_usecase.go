package usecase

import (
	"errors"

	"github.com/username/multi-inventory-manager/internal/domain"
)

type orderUsecase struct {
	repo          domain.OrderRepository
	inventoryRepo domain.InventoryRepository
	productRepo   domain.ProductRepository
}

func NewOrderUsecase(repo domain.OrderRepository, inventoryRepo domain.InventoryRepository, productRepo domain.ProductRepository) domain.OrderUsecase {
	return &orderUsecase{repo: repo, inventoryRepo: inventoryRepo, productRepo: productRepo}
}

func (u *orderUsecase) CreateTable(table *domain.Table) error {
	return u.repo.CreateTable(table)
}

func (u *orderUsecase) ListTables(branchID uint) ([]domain.Table, error) {
	return u.repo.ListTables(branchID)
}

func (u *orderUsecase) UpdateTable(table *domain.Table) error {
	return u.repo.UpdateTable(table)
}

func (u *orderUsecase) DeleteTable(id uint, branchID uint) error {
	return u.repo.DeleteTable(id, branchID)
}

func (u *orderUsecase) CreateOrder(order *domain.Order) error {
	// 1. If it's a table order, mark table as occupied
	if order.TableID != nil {
		if err := u.repo.UpdateTableStatus(*order.TableID, domain.TableOccupied); err != nil {
			return err
		}
	}

	// 2. Deduct product stock for each item in the order
	for i, item := range order.Items {
		product, err := u.productRepo.GetProductByID(item.ProductID, order.BranchID)
		if err != nil {
			return err
		}

		newStock := product.Stock - float64(item.Quantity)
		if newStock < 0 {
			return errors.New("insufficient stock for product: " + product.Name)
		}

		if err := u.repo.UpdateProductStock(item.ProductID, newStock); err != nil {
			return err
		}
		
		// Ensure the price is set from the product's current price
		order.Items[i].Price = product.Price
	}

	return u.repo.CreateOrder(order)
}

func (u *orderUsecase) GetOrderByID(id uint, branchID uint) (*domain.Order, error) {
	return u.repo.GetOrderByID(id, branchID)
}

func (u *orderUsecase) ListOrders(branchID uint) ([]domain.Order, error) {
	return u.repo.ListOrders(branchID)
}

func (u *orderUsecase) CompleteOrder(orderID uint, branchID uint) error {
	// 1. Get Order
	order, err := u.repo.GetOrderByID(orderID, branchID)
	if err != nil {
		return err
	}

	if order.Status == domain.OrderCompleted {
		return errors.New("order already completed")
	}

	// 2. Update order status
	if err := u.repo.UpdateOrderStatus(orderID, branchID, domain.OrderCompleted); err != nil {
		return err
	}

	// 3. If table order, mark table as available
	if order.TableID != nil {
		if err := u.repo.UpdateTableStatus(*order.TableID, domain.TableAvailable); err != nil {
			return err
		}
	}

	return nil
}
