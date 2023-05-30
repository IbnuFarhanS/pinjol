package model

import "time"

type Payment struct {
	ID              uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	TransactionID   uint          `gorm:"column:id_transaction" json:"transaction_id"`
	Transaction     Transaction   `gorm:"foreignKey:TransactionID" json:"transaction"`
	PaymentAmount   float64       `gorm:"not null" json:"payment_amount"`
	PaymentDate     time.Time     `gorm:"not null;default:now()" json:"payment_date"`
	PaymentMethodID uint          `json:"payment_method_id"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID" json:"payment_method"`

	NextInstallment float64 `gorm:"-" json:"next_installment"`
}
