package model

import "time"

type PaymentMethod struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Created_At time.Time `gorm:"column:created_at" json:"created_at"`
}
