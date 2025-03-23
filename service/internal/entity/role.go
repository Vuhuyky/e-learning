package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name"`
	Code ROLE   `json:"code" gorm:"unique"`

	Profiles        []Profile        `json:"profiles" gorm:"foreignKey:RoleId;"`
	RolePermissions []RolePermission `json:"rolePermissions" gorm:"foreignKey:RoleId;"`
}

type ROLE string

const (
	ADMIN ROLE = "ADMIN"
	USER  ROLE = "USER"
)
