package models

type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;unique" json:"-"`
	Books    []Book
}
