package model

import "time"

type Products struct {
	ID          int64     `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Name        string    `gorm:"column:name" validate:"required" json:"name"`
	Amount      float64   `gorm:"column:amount" validate:"required" json:"amount"`
	Installment int64     `gorm:"column:installment" validate:"required" json:"installment"`
	Bunga       float64   `gorm:"column:bunga" validate:"required" json:"bunga"`
	Created_At  time.Time `gorm:"column:created_at" validate:"required" json:"created_at"`
}
