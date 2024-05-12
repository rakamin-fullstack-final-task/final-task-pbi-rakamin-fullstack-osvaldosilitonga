package domain

import "time"

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Username  string    `json:"username"`
	Email     string    `gorm:"not null" json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"omitempty" json:"created_at"`
	UpdatedAt time.Time `gorm:"omitempty" json:"updated_at"`
}
