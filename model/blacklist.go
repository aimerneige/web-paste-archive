package model

import "gorm.io/gorm"

type Blacklist struct {
	gorm.Model
	UserIP    string `gorm:"type:varchar(255)"`
	UserAgent string `gorm:"type:varchar(255)"`
}
