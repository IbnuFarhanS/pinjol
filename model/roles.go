package model

import "time"

type Roles struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Created_at time.Time `gorm:"column:created_at" json:"created_at"`
}
