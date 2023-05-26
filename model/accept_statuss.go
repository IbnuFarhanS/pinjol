package model

import "time"

type AcceptStatus struct {
	ID           int64        `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Transactions Transactions `gorm:"foreignKey;column:id_transaction" validate:"required" json:"transaction"`
	Status       bool         `gorm:"column:status" validate:"required" json:"status"`
	Created_At   time.Time    `gorm:"column:created_at" validate:"required" json:"created_at"`
}
