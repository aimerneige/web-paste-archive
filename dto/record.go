package dto

import (
	"github.com/aimerneige/web-paste/model"
	"github.com/spf13/viper"
)

type RecordResponseDto struct {
	ID         uint   `json:"id"`
	Content    string `json:"content"`
	Title      string `json:"title"`
	CreateTime int64  `json:"create_time"`
	ExpireTime int64  `json:"expire_time"`
	ShortLink  string `json:"short_link"`
	RawLink    string `json:"raw_link"`
}

type RecordRequestDto struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}

// CreateRecordResponseDto create record response dto with record
func CreateRecordResponseDto(record model.Record) RecordResponseDto {
	baseUrl := viper.GetString("common.base")
	shortLink := baseUrl + record.ShortLink
	rawLink := baseUrl + "raw/" + record.RawLink
	return RecordResponseDto{
		ID:         record.ID,
		Content:    record.Content,
		Title:      record.Title,
		CreateTime: record.CreateTime,
		ExpireTime: record.ExpireTime,
		ShortLink:  shortLink,
		RawLink:    rawLink,
	}
}
