package model

import "time"

type Products struct {
	ID          int64     `gorm:"primaryKey;column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Installment int64     `gorm:"column:installment" json:"installment"`
	Bunga       float64   `gorm:"column:bunga" json:"bunga"`
	Created_At  time.Time `gorm:"column:created_at" json:"created_at"`
}
