package model

import (
	"time"
)

type Transactions struct {
	ID         int64     `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Products   Products  `gorm:"foreignKey;column:id_product" validate:"required" json:"product"`
	Users      Users     `gorm:"foreignKey;column:id_user" validate:"required" json:"users"`
	Status     bool      `gorm:"column:status" validate:"required" json:"status"`
	Created_At time.Time `gorm:"column:created_at" validate:"required" json:"created_at"`
	Due_Date   time.Time `gorm:"column:due_date" validate:"required" json:"due_date"`
}
