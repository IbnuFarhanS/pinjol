package model

import "time"

type Roles struct {
	ID         int64     `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Name       string    `gorm:"column:name" validate:"required" json:"name"`
	Created_at time.Time `gorm:"column:created_at" validate:"required" json:"created_at"`
}
