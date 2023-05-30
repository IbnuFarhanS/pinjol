package model

import "time"

<<<<<<< HEAD
type Products struct {
	ID          int64     `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Name        string    `gorm:"column:name" validate:"required" json:"name"`
	Installment int64     `gorm:"column:installment" validate:"required" json:"installment"`
	Bunga       float64   `gorm:"column:bunga" validate:"required" json:"bunga"`
	Created_At  time.Time `gorm:"column:created_at" validate:"required" json:"created_at"`
	// Amount      float64   `gorm:"column:amount" validate:"required" json:"amount"`
=======
type Product struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Installment int       `gorm:"not null" json:"installment"`
	Interest    float64   `gorm:"not null" json:"interest"`
	CreatedAt   time.Time `gorm:"not null;default:now()" json:"created_at"`
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
