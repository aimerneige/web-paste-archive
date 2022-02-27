package service

import (
	"github.com/aimerneige/web-paste/database"
	"github.com/aimerneige/web-paste/model"
)

// GetAllBlacklist get all blacklist
func GetAllBlacklist() (blacklist []model.Blacklist, err error) {
	db := database.GetDB()
	err = db.Find(&blacklist).Error

	return
}

// InsertNewBlacklist insert new blacklist
func InsertNewBlacklist(blacklist model.Blacklist) (err error) {
	db := database.GetDB()
	err = db.Create(blacklist).Error

	return
}
