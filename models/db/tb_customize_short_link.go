/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package db

import "github.com/jinzhu/gorm"

type TBCustomizeShortLink struct {
	gorm.Model
	ShortId   uint64 `gorm:"column:short_id;not null;primary_key" json:"shortId" description:"短码id"`
	ShortCode string `gorm:"column:short_code;not null;index:idx_short_code" json:"chipId" description:"短码"`
	LongUrl   string `gorm:"column:long_url;null;" json:"longUrl" description:"长连接"`
}

// TableName 设置表名字
func (t *TBCustomizeShortLink) TableName() string {
	return "tb_customize_short_link"
}

// AddTBCustomizeShortLink 插入记录
func (t *TBCustomizeShortLink) AddTBCustomizeShortLink() error {
	return _DB.Create(t).Error
}

// DeleteTBCustomizeShortLink 删除记录
func DeleteTBCustomizeShortLink(query map[string]interface{}, model interface{}) error {
	return _DB.Where(query).Delete(model).Error
}

// GetTBCustomizeShortLink 查找记录
func GetTBCustomizeShortLink(query map[string]interface{}) (*TBCustomizeShortLink, error) {
	var customizeShortLink TBCustomizeShortLink
	err := _DB.Where(query).First(&customizeShortLink).Error
	return &customizeShortLink, err
}

// GetAllTBCustomizeShortLinks 查找所有记录
func GetAllTBCustomizeShortLinks(query map[string]interface{}, offset int32, limit int32) (*[]TBCustomizeShortLink, error) {
	var customizeShortLinks []TBCustomizeShortLink
	err := _DB.Where(query).Offset(offset).Limit(limit).Find(&customizeShortLinks).Error
	return &customizeShortLinks, err
}
