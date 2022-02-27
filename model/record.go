package model

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	Content    string `gorm:"type:text"`
	Title      string `gorm:"type:varchar(255)"`
	CreateTime int64  `gorm:"type:bigint"`
	ExpireTime int64  `gorm:"type:bigint"`
	UserIP     string `gorm:"type:varchar(255)"`
	UserAgent  string `gorm:"type:varchar(255)"`
	ShortLink  string `gorm:"type:varchar(255),unique"`
	RawLink    string `gorm:"type:varchar(255),unique"`
}
