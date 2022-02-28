package controller

import (
	"path/filepath"
	"time"

	"github.com/aimerneige/web-paste/dto"
	"github.com/aimerneige/web-paste/model"
	"github.com/aimerneige/web-paste/response"
	"github.com/aimerneige/web-paste/service"
	"github.com/aimerneige/web-paste/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// CreateNewRecord create new record
func CreateNewRecord(c *gin.Context) {
	content := c.GetString("content")
	title := c.GetString("title")
	createTime := time.Now().Unix()
	expireTime := createTime + 3600*24
	userIP := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")
	shortLink := utils.GetRandomShortLink(6)
	rawLink := shortLink + ".txt"

	dirPath := viper.GetString("common.path")
	filePath := filepath.Join(dirPath, rawLink)
	err := utils.SaveContent(content, filePath)
	if err != nil {
		response.InternalServerError(c, err, "Fail to save content to file.")
		return
	}

	record := model.Record{
		Content:    content,
		Title:      title,
		CreateTime: createTime,
		ExpireTime: expireTime,
		UserIP:     userIP,
		UserAgent:  userAgent,
		ShortLink:  shortLink,
		RawLink:    rawLink,
	}
	err = service.InsertNewRecord(record)
	if err != nil {
		response.InternalServerError(c, err, "Fail to create new record.")
		return
	}

	dto := dto.CreateRecordResponseDto(record)
	response.Created(c, dto, "Success to create new record.")
}

// GetRecordByID get record by id
func GetRecordByID(c *gin.Context) {
	id := c.GetUint("id")

	record, err := service.GetRecordByID(id)
	if err != nil {
		response.NotFound(c, err, "Fail to get record by id.")
		return
	}

	dto := dto.CreateRecordResponseDto(record)
	response.OK(c, dto, "Success to get record by id.")
}

// DeleteRecordByID delete record by id
func DeleteRecordByID(c *gin.Context) {
	id := c.GetUint("id")

	err := service.DeleteRecordByID(id)
	if err != nil {
		response.InternalServerError(c, err, "Fail to delete record by id.")
		return
	}

	response.Deleted(c)
}
