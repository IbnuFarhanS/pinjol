package model

import "time"

type Payments struct {
	ID             int64         `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Transactions   Transactions  `gorm:"foreignKey;column:id_transaction" validate:"required" json:"transactions"`
	Payment_Method PaymentMethod `gorm:"foreignKey;column:id_payment_method" validate:"required" json:"payment_method"`
	Payment_Amount float64       `gorm:"column:payment_amount" validate:"required" json:"payment_amount"`
	Payment_Date   time.Time     `gorm:"column:payment_date" validate:"required" json:"payment_date"`
}
