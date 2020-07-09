/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package db

import (
	"time"
)

type TBShortLink struct {
	ShortId     uint64     `gorm:"column:short_id;not null;primary_key" json:"shortId" description:"短码id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"-"`
	ShortCode   string     `gorm:"column:short_code;not null;index:idx_short_code"  json:"shortCode" description:"短码"`
	LongUrl     string     `gorm:"column:long_url;null;" json:"longLink" description:"长连接"`
	LongLinkMd5 string     `gorm:"column:long_link_md5;null;" json:"-" description:"长连接md5, 用于查询"`
	Type        int8       `gorm:"column:type;null;" json:"type" description:"类型 0-系统生成 1-自定义"`
}

// TableName 设置表名字
func (t *TBShortLink) TableName() string {
	return "tb_short_link"
}

// AddTBShortLink 插入记录
func AddTBShortLink(t *TBShortLink) error {
	return _DB.Create(t).Error
}

// DeleteTBShortLink 删除记录
func DeleteTBShortLink(query map[string]interface{}, model interface{}) error {
	return _DB.Where(query).Delete(model).Error
}

// GetTBShortLink 查找记录
func GetTBShortLink(query map[string]interface{}) (*TBShortLink, error) {
	var sortLink TBShortLink
	err := _DB.Where(query).First(&sortLink).Error
	return &sortLink, err
}

// GetAllTBShortLinks 查找所有记录
func GetAllTBShortLinks(query map[string]interface{}, offset int32, limit int32) (*[]TBShortLink, error) {
	var sortLinks []TBShortLink
	err := _DB.Where(query).Offset(offset).Limit(limit).Find(&sortLinks).Error
	return &sortLinks, err
}
