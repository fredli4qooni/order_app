package model

import "time"

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Email     string    `gorm:"type:varchar(255);unique" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"-"`
	Role      string    `gorm:"type:enum('user','admin');default:'user'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
