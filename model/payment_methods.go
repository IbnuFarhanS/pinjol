package model

<<<<<<< HEAD
import (
	"time"
)

type PaymentMethod struct {
	ID         int64     `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Name       string    `gorm:"column:name" validate:"required" json:"name"`
	Created_At time.Time `gorm:"column:created_at" validate:"required" json:"created_at"`
=======
import "time"

type PaymentMethod struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `gorm:"not null;default:now()" json:"created_at"`
>>>>>>> 79e83b473a1c0aca2de729b88ccc29fed5de00a9
}
