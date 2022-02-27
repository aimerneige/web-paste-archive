package dto

import "github.com/aimerneige/web-paste/model"

type RecordResponseDto struct {
	ID         uint   `json:"id"`
	Content    string `json:"content"`
	Title      string `json:"title"`
	CreateTime int64  `json:"create_time"`
	ExpireTime int64  `json:"expire_time"`
	ShortLink  string `json:"short_link"`
	RawLink    string `json:"raw_link"`
}

// CreateRecordResponseDto create record response dto with record
func CreateRecordResponseDto(record *model.Record) *RecordResponseDto {
	return &RecordResponseDto{
		ID:         record.ID,
		Content:    record.Content,
		Title:      record.Title,
		CreateTime: record.CreateTime,
		ExpireTime: record.ExpireTime,
		ShortLink:  record.ShortLink,
		RawLink:    record.RawLink,
	}
}
