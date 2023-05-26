package request

import (
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
)

type CreateUsersRequest struct {
	Username     string        `gorm:"column:username" validate:"required" json:"username"`
	Password     string        `gorm:"column:password" validate:"required" json:"password"`
	Nik          string        `gorm:"column:nik" validate:"required" json:"nik"`
	Name         string        `gorm:"column:name" validate:"required" json:"name"`
	Alamat       string        `gorm:"column:alamat" validate:"required" json:"alamat"`
	Phone_Number string        `gorm:"column:phone_number" validate:"required" json:"phone_number"`
	Limit        float64       `gorm:"column:limit" validate:"required" json:"limit"`
	RolesID      int           `gorm:"column:id_role" validate:"required" json:"id_role"`
	Roles        []model.Roles `gorm:"foreignKey:RolesID" validate:"required" json:"-"`
	Created_At   time.Time     `gorm:"column:created_at" validate:"required" json:"created_at"`
}

type UpdateUsersRequest struct {
	ID           int64   `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Username     string  `gorm:"column:username" validate:"required" json:"username"`
	Password     string  `gorm:"column:password" validate:"required" json:"password"`
	Nik          string  `gorm:"column:nik" validate:"required" json:"nik"`
	Name         string  `gorm:"column:name" validate:"required" json:"name"`
	Alamat       string  `gorm:"column:alamat" validate:"required" json:"alamat"`
	Phone_Number string  `gorm:"column:phone_number" validate:"required" json:"phone_number"`
	Limit        float64 `gorm:"column:limit" validate:"required" json:"limit"`
	// Roles        []model.Roles `gorm:"foreignKey:ID_Role" validate:"required" json:"roles"`
	Created_At time.Time `gorm:"column:created_at" validate:"required" json:"created_at"`
}

type LoginRequest struct {
	Username string `validate:"required,max=200,min=2" json:"username"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}
