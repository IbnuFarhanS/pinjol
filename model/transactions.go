package model

<<<<<<< HEAD
import (
	"time"
)

type Transactions struct {
	ID         int64     `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	ProductsID int64     `gorm:"column:id_product" validate:"required" json:"id_product"`
	Products   Products  `gorm:"foreignKey:ProductsID" validate:"required" json:"-"`
	UsersID    int64     `gorm:"column:id_user" validate:"required" json:"id_user"`
	Users      Users     `gorm:"foreignKey:UsersID" validate:"required" json:"-"`
	Status     bool      `gorm:"column:status" validate:"required" json:"status"`
	Created_At time.Time `gorm:"column:created_at" validate:"required" json:"created_at"`
	Due_Date   time.Time `gorm:"column:due_date" validate:"required" json:"due_date"`
=======
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
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
