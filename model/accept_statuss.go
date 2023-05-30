package model

import "time"

type AcceptStatus struct {
<<<<<<< HEAD
	ID             int64        `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	TransactionsID int64        `gorm:"column:id_transaction" validate:"required" json:"id_transaction"`
	Transactions   Transactions `gorm:"foreignKey:TransactionsID" validate:"required" json:"-"`
	Status         bool         `gorm:"column:status" validate:"required" json:"status"`
	Created_At     time.Time    `gorm:"column:created_at" validate:"required" json:"created_at"`
=======
	ID            uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	TransactionID uint        `gorm:"column:id_transaction" json:"id_transaction"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID" json:"-"`
	Status        bool        `gorm:"not null" json:"status"`
	CreatedAt     time.Time   `gorm:"not null;default:now()" json:"created_at"`
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
