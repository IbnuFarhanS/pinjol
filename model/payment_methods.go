package model

import "time"

type PaymentMethod struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `gorm:"not null;default:now()" json:"created_at"`
}
