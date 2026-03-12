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

func (u *reportUsecase) GetDashboardSummary(tenantID uint) (*domain.DashboardSummary, error) {
	var summary domain.DashboardSummary
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// 1. Inventory stats
	u.db.Model(&domain.Item{}).Where("tenant_id = ?", tenantID).Count(&summary.TotalItems)
	u.db.Model(&domain.Item{}).Where("tenant_id = ? AND current_stock <= min_stock_alert", tenantID).Count(&summary.LowStockItems)

	// 2. Product stats
	u.db.Model(&domain.Product{}).Where("tenant_id = ?", tenantID).Count(&summary.TotalProducts)

	// 3. Order stats (Today)
	u.db.Model(&domain.Order{}).Where("tenant_id = ? AND created_at >= ?", tenantID, startOfDay).Count(&summary.TodayOrders)
	u.db.Model(&domain.Order{}).Where("tenant_id = ? AND created_at >= ? AND status = ?", tenantID, startOfDay, domain.OrderCompleted).
		Select("COALESCE(SUM(total_price), 0)").Scan(&summary.TodayRevenue)

	// 4. Table stats
	u.db.Model(&domain.Table{}).Where("tenant_id = ?", tenantID).Count(&summary.TotalTables)
	u.db.Model(&domain.Table{}).Where("tenant_id = ? AND status = ?", tenantID, domain.TableOccupied).Count(&summary.OccupiedTables)

	return &summary, nil
}
