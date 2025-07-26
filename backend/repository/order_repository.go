package repository

import (
	"gorm.io/gorm"
	"it-order-app/backend/model"
)

type OrderRepository interface {
	Save(order model.Order) (model.Order, error)
	FindByUserID(userID uint) ([]model.Order, error)
	FindAll() ([]model.Order, error)
	FindByID(orderID uint) (model.Order, error)
	UpdateStatus(orderID uint, status string) (model.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

// Save menyimpan order baru ke database
func (r *orderRepository) Save(order model.Order) (model.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}
	// Preload User setelah membuat agar response berisi data user juga
	err = r.db.Preload("User").First(&order, order.ID).Error
	return order, err
}

// FindByUserID mengambil semua order milik satu user
func (r *orderRepository) FindByUserID(userID uint) ([]model.Order, error) {
	var orders []model.Order
	// Preload("User") untuk mengambil data user terkait
	// Mengurutkan berdasarkan yang terbaru
	err := r.db.Preload("User").Where("user_id = ?", userID).Order("created_at DESC").Find(&orders).Error
	if err != nil {
		return orders, err
	}
	return orders, nil
}

// FindAll mengambil semua order dari semua user, dengan data user-nya
func (r *orderRepository) FindAll() ([]model.Order, error) {
	var orders []model.Order
	// Preload("User") akan otomatis melakukan JOIN ke tabel users
	err := r.db.Preload("User").Order("created_at DESC").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// FindByID mencari satu order berdasarkan ID
func (r *orderRepository) FindByID(orderID uint) (model.Order, error) {
	var order model.Order
	// Preload("User") untuk mengambil data user terkait
	err := r.db.Preload("User").Where("id = ?", orderID).First(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

// UpdateStatus mengubah status sebuah order
func (r *orderRepository) UpdateStatus(orderID uint, status string) (model.Order, error) {
	// Pertama, panggil FindByID untuk mendapatkan order beserta data user-nya
	order, err := r.FindByID(orderID)
	if err != nil {
		return order, err // Return error jika order tidak ditemukan
	}

	// Update kolom status
	order.Status = status

	// Simpan perubahan ke database
	if err := r.db.Save(&order).Error; err != nil {
		return order, err
	}

	// `order` sudah berisi data User dari FindByID, jadi bisa langsung dikembalikan
	return order, nil
}
