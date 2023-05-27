package model

import "time"

type Users struct {
	ID           int64     `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Username     string    `gorm:"column:username" validate:"required" json:"username"`
	Password     string    `gorm:"column:password" validate:"required,min=5" json:"password"`
	Nik          string    `gorm:"column:nik" validate:"required" json:"nik"`
	Name         string    `gorm:"column:name" validate:"required" json:"name"`
	Alamat       string    `gorm:"column:alamat" validate:"required" json:"alamat"`
	Phone_Number string    `gorm:"column:phone_number" validate:"required" json:"phone_number"`
	Limit        float64   `gorm:"column:limit" validate:"gte=0" json:"limit"`
	RolesID      int       `gorm:"column:id_role" validate:"required" json:"id_role"`
	Roles        Roles     `gorm:"foreignKey:RolesID" validate:"required" json:"-"`
	Created_At   time.Time `gorm:"column:created_at" validate:"required" json:"created_at"`
}
