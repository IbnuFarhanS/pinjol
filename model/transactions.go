package model

import (
	"time"
)

type Transactions struct {
	ID         int64     `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	ProductsID int       `gorm:"column:id_product" validate:"required" json:"id_product"`
	Products   Products  `gorm:"foreignKey:ProductsID" validate:"required" json:"-"`
	UsersID    int       `gorm:"column:id_user" validate:"required" json:"id_user"`
	Users      Users     `gorm:"foreignKey:UsersID" validate:"required" json:"-"`
	Status     bool      `gorm:"column:status" validate:"required" json:"status"`
	Created_At time.Time `gorm:"column:created_at" validate:"required" json:"created_at"`
	Due_Date   time.Time `gorm:"column:due_date" validate:"required" json:"due_date"`
}
