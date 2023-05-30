package model

import "time"

type Payment struct {
	ID              uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	TransactionID   uint          `gorm:"column:id_transaction" json:"id_transaction"`
	Transaction     Transaction   `gorm:"foreignKey:TransactionID" json:"-"`
	PaymentAmount   float64       `gorm:"not null" json:"payment_amount"`
	PaymentDate     time.Time     `gorm:"not null;default:now()" json:"payment_date"`
	PaymentMethodID uint          `gorm:"column:id_payment_method" json:"id_payment_method"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID" json:"-"`

	NextInstallment float64 `gorm:"-" json:"next_installment"`
}
