package model

import "github.com/jinzhu/gorm"

type Wirte struct {
	gorm.Model
	Name string `gorm:"varchar(25);not null"`
	Title string `gorm:"varchar(255);not null"`
	News_tpye string `gorm:"varchar(25);not null"`
	Author string `gorm:"varchar(25);not null"`
	Content string `gorm:"text;not null"`
	Time string `gorm:"varchar(25);not null"`
	Image string `gorm:"varchar(1024);not null"`

}
