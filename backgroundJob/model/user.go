package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"varchar(25);not null"`
	Number string `gorm:"varchar(110);not null;unique"`
	Password string `gorm:"size:255;not null"`
}


