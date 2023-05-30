package model

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Installment int       `gorm:"not null" json:"installment"`
	Interest    float64   `gorm:"not null" json:"interest"`
	CreatedAt   time.Time `gorm:"not null;default:now()" json:"created_at"`
}
