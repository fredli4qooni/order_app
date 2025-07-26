package model

// Stats data statistik dashboard
type Stats struct {
	TotalOrders       int64            `json:"total_orders"`
	OrdersByStatus    map[string]int64 `json:"orders_by_status"`
	TopUser           TopUserStats     `json:"top_user"`
	AvgCompletionTime string           `json:"avg_completion_time"`
}

// TopUserStats berisi info user dengan order terbanyak
type TopUserStats struct {
	UserName   string `json:"user_name"`
	OrderCount int64  `json:"order_count"`
}
