package model

import "time"

type Users struct {
	ID           int64     `gorm:"primaryKey;column:id" json:"id"`
	Username     string    `gorm:"column:username" json:"username"`
	Password     string    `gorm:"column:password" json:"password"`
	Nik          string    `gorm:"column:nik" json:"nik"`
	Name         string    `gorm:"column:name" json:"name"`
	Alamat       string    `gorm:"column:alamat" json:"alamat"`
	Phone_Number string    `gorm:"column:phone_number" json:"phone_number"`
	Limit        float64   `gorm:"column:limit" json:"limit"`
	RolesID      int64     `gorm:"column:id_role" json:"id_role"`
	Roles        Roles     `gorm:"foreignKey:RolesID" json:"-"`
	Created_At   time.Time `gorm:"column:created_at" json:"created_at"`
}

type FileKTP struct {
	Filename string
	Path     string
}
