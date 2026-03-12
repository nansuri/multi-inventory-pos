package domain

type DashboardSummary struct {
	TotalItems     int64   `json:"total_items"`
	LowStockItems  int64   `json:"low_stock_items"`
	TotalProducts  int64   `json:"total_products"`
	TodayOrders    int64   `json:"today_orders"`
	TodayRevenue   float64 `json:"today_revenue"`
	TotalTables    int64   `json:"total_tables"`
	OccupiedTables int64   `json:"occupied_tables"`
}

type ReportUsecase interface {
	GetDashboardSummary(tenantID uint) (*DashboardSummary, error)
}
