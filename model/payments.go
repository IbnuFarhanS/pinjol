package model

import "time"

type Payments struct {
	ID              int64         `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	TransactionsID  int           `gorm:"column:id_transaction" validate:"required" json:"id_transaction"`
	Transactions    Transactions  `gorm:"foreignKey:TransactionsID" validate:"required" json:"-"`
	PaymentMethodID int           `gorm:"column:id_payment_method" validate:"required" json:"id_payment_method"`
	Payment_Method  PaymentMethod `gorm:"foreignKey:PaymentMethodID" validate:"required" json:"-"`
	Payment_Amount  float64       `gorm:"column:payment_amount" validate:"required" json:"payment_amount"`
	Payment_Date    time.Time     `gorm:"column:payment_date" validate:"required" json:"payment_date"`
}
