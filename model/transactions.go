package model

import "time"

type Transactions struct {
	ID int64 `gorm:"primaryKey;column:id" json:"id"`

	ProductsID int64    `gorm:"column:id_product" json:"id_product"`
	Products   Products `gorm:"foreignKey:ProductsID" json:"-"`

	UsersID int64 `gorm:"column:id_user" json:"id_user"`
	Users   Users `gorm:"foreignKey:UsersID" json:"-"`

	Status bool    `gorm:"column:status" json:"status"`
	Amount float64 `gorm:"column:amount" json:"amount"`

	Created_At time.Time `gorm:"column:created_at" json:"created_at"`
	Due_Date   time.Time `gorm:"column:due_date" json:"due_date"`

	Totalmounth float64 `gorm:"-" json:"total_mounth"`
	TotalTax    float64 `gorm:"-" json:"total_tax"`
	Total       float64 `gorm:"-" json:"total"`
}
