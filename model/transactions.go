package model

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"column:id_user" json:"id_user"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	ProductID uint      `gorm:"column:id_product" json:"id_product"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"-"`
	Amount    float64   `gorm:"not null" json:"amount"`
	Status    bool      `gorm:"not null" json:"status"`
	CreatedAt time.Time `gorm:"not null;default:now()" json:"created_at"`
	DueDate   time.Time `gorm:"not null;default:now()" json:"due_date"`

	Total      float64 `gorm:"-" json:"total"`
	TotalMonth float64 `gorm:"-" json:"total_month"`
	TotalTax   float64 `gorm:"-" json:"total_tax"`
}
