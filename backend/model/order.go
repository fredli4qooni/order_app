package model

import "time"

type Order struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	UserID      uint      `json:"user_id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Category    string    `gorm:"type:varchar(50)" json:"category"`
	Status      string    `gorm:"type:enum('Pending','In Progress','Done');default:'Pending'" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	User User `gorm:"foreignKey:UserID"`
}

func (Order) TableName() string {
	return "orders"
}
