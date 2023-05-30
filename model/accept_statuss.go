package model

import "time"

type AcceptStatus struct {
	ID             int64        `gorm:"primaryKey;column:id" json:"id"`
	TransactionsID int64        `gorm:"column:id_transaction" json:"id_transaction"`
	Transactions   Transactions `gorm:"foreignKey:TransactionsID" json:"-"`
	Status         bool         `gorm:"column:status" json:"status"`
	Created_At     time.Time    `gorm:"column:created_at" json:"created_at"`
}
