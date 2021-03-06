package service

import (
	"time"

	"github.com/aimerneige/web-paste/database"
	"github.com/aimerneige/web-paste/model"
)

// InsertNewRecord insert new record
func InsertNewRecord(record model.Record) (err error) {
	db := database.GetDB()
	err = db.Create(&record).Error

	return
}

// GetRecordByShortLink get record by short link
func GetRecordByShortLink(shortLink string) (record model.Record, err error) {
	db := database.GetDB()
	err = db.Where("short_link = ?", shortLink).First(&record).Error

	return
}

// DeleteRecordByID delete record by id
func DeleteRecordByID(id uint) (err error) {
	db := database.GetDB()
	err = db.Delete(&model.Record{}, id).Error

	return
}

// DeleteAllExpiredRecords delete all expired records
func DeleteAllExpiredRecords() (err error) {
	db := database.GetDB()
	err = db.Where("expire_time < ?", time.Now().Unix()).Delete(&model.Record{}).Error

	return
}
