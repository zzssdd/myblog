package model

import (
	"gorm.io/gorm"
	"myblog/utils/errmsg"
)

type Life struct {
	gorm.Model
	Title     string `gorm:"type:varchar(20);not null;" json:"title"`
	Desc      string `gorm:"type:varchar(20);not null;" json:"desc"`
	Content   string `gorm:"type:longtext;" json:"content"`
	Img       string `gorm:"type:varchar(100);" json:"img"`
	ReadCount int    `gorm:"type:int;default:0"json:"num"`
}

func CreateLife(data *Life) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteLife(id int) int {
	var life Life
	err := db.Where("id=?", id).Delete(&life).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func EditLife(id int, data *Life) int {
	var life Life
	var maps = make(map[string]interface{})
	maps["Title"] = data.Title
	maps["Desc"] = data.Desc
	maps["Content"] = data.Content
	maps["Img"] = data.Img
	err := db.Model(&life).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetLife(pageSize int, pageNum int) ([]Life, int) {
	var life []Life
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&life).Order("created_at DESC").Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return life, errmsg.SUCCESS
}

func GetLifeInfo(id int) (Life, int) {
	var life Life
	err := db.Where("id=?", id).First(&life).Error
	err = db.Model(&life).Where("id=?", id).UpdateColumn("read_count", gorm.Expr("read_count+1")).Error
	if err != nil {
		return life, errmsg.ERROR_LIFE_NOT_EXIST
	}
	return life, errmsg.SUCCESS
}
