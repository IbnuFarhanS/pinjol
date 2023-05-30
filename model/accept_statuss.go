package model

import "time"

type AcceptStatus struct {
	ID            uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	TransactionID uint        `gorm:"column:id_transaction" json:"id_transaction"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID" json:"-"`
	Status        bool        `gorm:"not null" json:"status"`
	CreatedAt     time.Time   `gorm:"not null;default:now()" json:"created_at"`
}
