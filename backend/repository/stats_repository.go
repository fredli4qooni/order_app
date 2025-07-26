package repository

import (
	"fmt"
	"gorm.io/gorm"
	"it-order-app/backend/model"
	"time"
)

type StatsRepository interface {
	GetStats() (model.Stats, error)
}

type statsRepository struct {
	db *gorm.DB
}

func NewStatsRepository(db *gorm.DB) StatsRepository {
	return &statsRepository{db}
}

func (r *statsRepository) GetStats() (model.Stats, error) {
	var stats model.Stats
	var err error

	// 1. Jumlah Total Order
	if err = r.db.Model(&model.Order{}).Count(&stats.TotalOrders).Error; err != nil {
		return stats, err
	}

	// 2. Order per Status
	type StatusCount struct {
		Status string
		Count  int64
	}
	var statusCounts []StatusCount
	if err = r.db.Model(&model.Order{}).Select("status, count(*) as count").Group("status").Scan(&statusCounts).Error; err != nil {
		return stats, err
	}
	stats.OrdersByStatus = make(map[string]int64)
	for _, sc := range statusCounts {
		stats.OrdersByStatus[sc.Status] = sc.Count
	}

	// 3. Order Terbanyak per User
	type UserOrderCount struct {
		UserID     uint
		Name       string
		OrderCount int64
	}
	var topUserCount UserOrderCount
	err = r.db.Model(&model.Order{}).
		Select("orders.user_id, users.name, COUNT(orders.id) as order_count").
		Joins("join users on users.id = orders.user_id").
		Group("orders.user_id, users.name").
		Order("order_count DESC").
		Limit(1).
		Scan(&topUserCount).Error
	if err == nil { // Hanya isi jika tidak ada error (misal, jika ada order)
		stats.TopUser = model.TopUserStats{
			UserName:   topUserCount.Name,
			OrderCount: topUserCount.OrderCount,
		}
	}

	// 4. Waktu Rata-rata Penyelesaian Order (dari 'Pending' ke 'Done')
	type AvgTimeResult struct {
		AvgSeconds float64
	}
	var avgResult AvgTimeResult
	err = r.db.Model(&model.Order{}).
		Select("AVG(TIMESTAMPDIFF(SECOND, created_at, updated_at)) as avg_seconds").
		Where("status = ?", "Done").
		Scan(&avgResult).Error
	if err == nil && avgResult.AvgSeconds > 0 {
		avgDuration := time.Duration(avgResult.AvgSeconds) * time.Second
		stats.AvgCompletionTime = formatDuration(avgDuration)
	} else {
		stats.AvgCompletionTime = "N/A"
	}

	return stats, nil
}

// Fungsi helper
func formatDuration(d time.Duration) string {
	d = d.Round(time.Second)
	days := d / (24 * time.Hour)
	d -= days * 24 * time.Hour
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute

	return fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
}
