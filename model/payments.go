package model

import "time"

type Payments struct {
	ID int64 `gorm:"primaryKey;column:id" json:"id"`

	TransactionsID int64        `gorm:"column:id_transaction" json:"id_transaction"`
	Transactions   Transactions `gorm:"foreignKey:TransactionsID" json:"-"`

	PaymentMethodID int64         `gorm:"column:id_payment_method" json:"id_payment_method"`
	Payment_Method  PaymentMethod `gorm:"foreignKey:PaymentMethodID" json:"-"`

	Payment_Amount float64   `gorm:"column:payment_amount" json:"payment_amount"`
	Payment_Date   time.Time `gorm:"column:payment_date" json:"payment_date"`

	NextInstallment float64 `gorm:"-" json:"netx_installment"`
}
