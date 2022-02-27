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

// GetRecordByID get record by id
func GetRecordByID(id uint) (record model.Record, err error) {
	db := database.GetDB()
	err = db.First(&record, id).Error

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
