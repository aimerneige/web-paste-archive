package database

import (
	"github.com/aimerneige/web-paste/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DatabaseInterface interface {
	InitDB(migrateDst ...interface{}) (*gorm.DB, error)
}

// InitDatabase init database
func InitDatabase(dbi DatabaseInterface) {
	db, err := dbi.InitDB(&model.Record{})
	if err != nil {
		panic(err)
	}

	DB = db
}

// GetDB get database
func GetDB() *gorm.DB {
	return DB
}
