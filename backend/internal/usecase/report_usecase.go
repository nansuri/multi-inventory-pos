package usecase

import (
	"time"

	"github.com/username/multi-inventory-manager/internal/domain"
	"gorm.io/gorm"
)

type reportUsecase struct {
	db *gorm.DB
}

func NewReportUsecase(db *gorm.DB) domain.ReportUsecase {
	return &reportUsecase{db: db}
}

func (u *reportUsecase) GetDashboardSummary(tenantID uint, branchID uint) (*domain.DashboardSummary, error) {
	var summary domain.DashboardSummary
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// Helper to handle optional branch filtering
	scopedDB := func(model interface{}) *gorm.DB {
		db := u.db.Model(model).Where("tenant_id = ?", tenantID)
		if branchID != 0 {
			db = db.Where("branch_id = ?", branchID)
		}
		return db
	}

	// 1. Inventory stats
	scopedDB(&domain.Item{}).Count(&summary.TotalItems)
	scopedDB(&domain.Item{}).Where("current_stock <= min_stock_alert").Count(&summary.LowStockItems)

	// 2. Product stats
	scopedDB(&domain.Product{}).Count(&summary.TotalProducts)

	// 3. Order stats (Today)
	scopedDB(&domain.Order{}).Where("created_at >= ?", startOfDay).Count(&summary.TodayOrders)
	scopedDB(&domain.Order{}).Where("created_at >= ? AND status = ?", startOfDay, domain.OrderCompleted).
		Select("COALESCE(SUM(total_price), 0)").Scan(&summary.TodayRevenue)

	// 4. Table stats
	scopedDB(&domain.Table{}).Count(&summary.TotalTables)
	scopedDB(&domain.Table{}).Where("status = ?", domain.TableOccupied).Count(&summary.OccupiedTables)
	
	// Fetch actual tables for floor plan view
	if branchID != 0 {
		u.db.Where("tenant_id = ? AND branch_id = ?", tenantID, branchID).Order("table_number asc").Find(&summary.Tables)
	} else {
		// For aggregated, maybe don't show specific tables or show all
		u.db.Where("tenant_id = ?", tenantID).Order("branch_id, table_number asc").Find(&summary.Tables)
	}

	return &summary, nil
}

func (u *reportUsecase) GetOrderHistory(tenantID uint, branchID uint, period string) ([]domain.Order, float64, int64, error) {
	var orders []domain.Order
	var totalRevenue float64
	var orderCount int64

	now := time.Now()
	var startTime time.Time

	switch period {
	case "day":
		startTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	case "week":
		offset := int(now.Weekday()) - 1
		if offset < 0 {
			offset = 6
		}
		startTime = time.Date(now.Year(), now.Month(), now.Day()-offset, 0, 0, 0, 0, now.Location())
	case "month":
		startTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	default:
		startTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	}

	query := u.db.Where("tenant_id = ? AND created_at >= ? AND status = ?", tenantID, startTime, domain.OrderCompleted)
	if branchID != 0 {
		query = query.Where("branch_id = ?", branchID)
	}
	
	err := query.Preload("Table").Preload("Items.Product").Order("created_at desc").Find(&orders).Error
	if err != nil {
		return nil, 0, 0, err
	}

	countQuery := u.db.Model(&domain.Order{}).Where("tenant_id = ? AND created_at >= ? AND status = ?", tenantID, startTime, domain.OrderCompleted)
	if branchID != 0 {
		countQuery = countQuery.Where("branch_id = ?", branchID)
	}
	countQuery.Count(&orderCount)
	countQuery.Select("COALESCE(SUM(total_price), 0)").Scan(&totalRevenue)

	return orders, totalRevenue, orderCount, nil
}
