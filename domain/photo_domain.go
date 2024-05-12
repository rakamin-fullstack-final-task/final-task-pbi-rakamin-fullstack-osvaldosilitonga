package domain

import "time"

type Photo struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserID    uint64
	Title     string
	Caption   string
	PhotoUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"foreignKey:user_id;references:id"`
}
