package response

import (
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
)

type UsersResponse struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username    string     `gorm:"unique;not null" json:"username"`
	Password    string     `gorm:"not null" json:"password"`
	NIK         string     `gorm:"not null" json:"nik"`
	Name        string     `gorm:"not null" json:"name"`
	Address     string     `gorm:"not null" json:"address"`
	PhoneNumber string     `gorm:"not null" json:"phone_number"`
	Limit       float64    `gorm:"not null" json:"limit"`
	RoleID      uint       `gorm:"column:id_role" json:"id_role"`
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	Role        model.Role `gorm:"foreignKey:RoleID" json:"role"`
}

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
