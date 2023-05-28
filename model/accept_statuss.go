package model

import "time"

type AcceptStatus struct {
	ID             int64        `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	TransactionsID int64        `gorm:"column:id_transaction" validate:"required" json:"id_transaction"`
	Transactions   Transactions `gorm:"foreignKey:TransactionsID" validate:"required" json:"-"`
	Status         bool         `gorm:"column:status" validate:"required" json:"status"`
	Created_At     time.Time    `gorm:"column:created_at" validate:"required" json:"created_at"`
}
