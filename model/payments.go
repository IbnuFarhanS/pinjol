package model

import "time"

<<<<<<< HEAD
type Payments struct {
	ID              int64         `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	TransactionsID  int           `gorm:"column:id_transaction" validate:"required" json:"id_transaction"`
	Transactions    Transactions  `gorm:"foreignKey:TransactionsID" validate:"required" json:"-"`
	PaymentMethodID int           `gorm:"column:id_payment_method" validate:"required" json:"id_payment_method"`
	Payment_Method  PaymentMethod `gorm:"foreignKey:PaymentMethodID" validate:"required" json:"-"`
	Payment_Amount  float64       `gorm:"column:payment_amount" validate:"required" json:"payment_amount"`
	Payment_Date    time.Time     `gorm:"column:payment_date" validate:"required" json:"payment_date"`
=======
type Payment struct {
	ID              uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	TransactionID   uint          `gorm:"column:id_transaction" json:"id_transaction"`
	Transaction     Transaction   `gorm:"foreignKey:TransactionID" json:"-"`
	PaymentAmount   float64       `gorm:"not null" json:"payment_amount"`
	PaymentDate     time.Time     `gorm:"not null;default:now()" json:"payment_date"`
	PaymentMethodID uint          `gorm:"column:id_payment_method" json:"id_payment_method"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID" json:"-"`

	NextInstallment float64 `gorm:"-" json:"next_installment"`
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
