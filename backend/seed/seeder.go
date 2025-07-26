package main

import (
	"fmt"
	"it-order-app/backend/config"
	"it-order-app/backend/model"
	"log"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	// Muat .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Koneksi DB
	config.ConnectDatabase()
	db := config.DB

	// Panggil fungsi untuk seed users dan orders
	seedUsers(db)
	seedOrders(db)

	fmt.Println("Seeding finished successfully.")
}

// Fungsi untuk membuat user
func seedUsers(db *gorm.DB) {
	users := []model.User{
		{Name: "Admin IT", Email: "admin@it.com", Password: "password", Role: "admin"},
		{Name: "Budi Karyawan", Email: "budi@example.com", Password: "password", Role: "user"},
		{Name: "Ani Marketing", Email: "ani@example.com", Password: "password", Role: "user"},
		{Name: "Cici Finansial", Email: "cici@example.com", Password: "password", Role: "user"},
	}

	for _, u := range users {
		var existingUser model.User
		if db.Where("email = ?", u.Email).First(&existingUser).Error == gorm.ErrRecordNotFound {
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
			u.Password = string(hashedPassword)
			db.Create(&u)
			fmt.Printf("User created: %s\n", u.Email)
		}
	}
}

// Fungsi untuk membuat order
func seedOrders(db *gorm.DB) {
	var orderCount int64
	db.Model(&model.Order{}).Count(&orderCount)
	if orderCount > 0 {
		fmt.Println("Orders table is not empty. Skipping order seeding.")
		return
	}

	// Data dummy orders
	orders := []model.Order{
		// Orders untuk Budi (user_id = 2)
		{UserID: 2, Title: "Laptop lambat setelah update", Description: "Windows 10 terasa sangat lambat setelah update terakhir.", Category: "Hardware", Status: "Pending", CreatedAt: time.Now().AddDate(0, 0, -5)},
		{UserID: 2, Title: "Request akses ke folder Marketing", Description: "Saya butuh akses read/write ke folder shared Marketing.", Category: "Network", Status: "In Progress", CreatedAt: time.Now().AddDate(0, 0, -3)},
		{UserID: 2, Title: "Install Adobe Premiere Pro", Description: "Untuk kebutuhan editing video campaign.", Category: "Software", Status: "Done", CreatedAt: time.Now().AddDate(0, 0, -10)},

		// Orders untuk Ani (user_id = 3)
		{UserID: 3, Title: "Printer tidak bisa scan", Description: "Fungsi print berjalan normal, tapi scan tidak terdeteksi.", Category: "Hardware", Status: "Pending", CreatedAt: time.Now().AddDate(0, 0, -1)},
		{UserID: 3, Title: "Email tidak bisa mengirim lampiran >10MB", Description: "Setiap mencoba mengirim lampiran besar, email gagal terkirim.", Category: "Software", Status: "Done", CreatedAt: time.Now().AddDate(0, 0, -15)},

		// Orders untuk Cici (user_id = 4)
		{UserID: 4, Title: "Koneksi WiFi sering putus", Description: "Setiap 15-30 menit, koneksi WiFi di laptop saya terputus.", Category: "Network", Status: "Done", CreatedAt: time.Now().AddDate(0, 0, -7)},
		{UserID: 4, Title: "Butuh Keyboard Eksternal", Description: "Keyboard laptop kurang nyaman untuk input data banyak.", Category: "Hardware", Status: "Pending", CreatedAt: time.Now()},
	}

	if err := db.Create(&orders).Error; err != nil {
		log.Fatalf("Failed to seed orders: %v", err)
	}
	fmt.Printf("%d orders created.\n", len(orders))

	// Update manual untuk order yang 'Done' agar punya updated_at yang berbeda
	updateDoneOrders(db)
}

// Fungsi untuk update order yang sudah 'Done'
func updateDoneOrders(db *gorm.DB) {
	fmt.Println("Updating 'Done' orders with completion times...")
	// Order ID 3 (Install Premiere Pro), selesai dalam 2 hari
	db.Model(&model.Order{}).Where("id = ?", 3).Update("updated_at", time.Now().AddDate(0, 0, -8))

	// Order ID 5 (Email Attachment), selesai dalam 5 hari
	db.Model(&model.Order{}).Where("id = ?", 5).Update("updated_at", time.Now().AddDate(0, 0, -10))

	// Order ID 6 (WiFi Putus), selesai dalam 1 hari
	db.Model(&model.Order{}).Where("id = ?", 6).Update("updated_at", time.Now().AddDate(0, 0, -6))
}
