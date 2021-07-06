package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name       string `gorm:"type:varchar(50);"`
	LastName   string `gorm:"type:varchar(50);"`
	Occupation string `gorm:"type:varchar(255);"`
	Email      string `gorm:"type:varchar(255);"`
	Password   string `gorm:"type:varchar(255);"`
}
