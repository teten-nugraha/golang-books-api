package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserId      uint
	User        User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
