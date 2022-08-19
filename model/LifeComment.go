package model

import (
	"fmt"
	"gorm.io/gorm"
	"myblog/utils/errmsg"
)

type LifeComment struct {
	gorm.Model
	Life    Life   `gorm:"foreignkey:Lid"`
	Lid     int    `gorm:"type:int;not null" json:"lid"`
	Name    string `gorm:"type:varchar(20);not null" json:"name"`
	Content string `gorm:"type:longtext;not null" json:"content"`
}

func CreateLComment(data *LifeComment) int {
	err := db.Create(&data).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteLcomment(id int) int {
	var lifecomment LifeComment
	err := db.Where("id=?", id).Delete(&lifecomment).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetLcomment(pageSize int, pageNum int) ([]LifeComment, int) {
	var lifecomment []LifeComment
	err := db.Joins("Life").Select("life_comment.id", "Life.title", "name", "life_comment.created_at", "life_comment.content").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("life_comment.created_at DESC").Find(&lifecomment).Error
	if err != nil {
		fmt.Println(err)
		return nil, errmsg.ERROR
	}
	return lifecomment, errmsg.SUCCESS
}

func GetLifeComment(pageSize int, pageNum int, id int) ([]LifeComment, int) {
	var lifecomment []LifeComment
	err := db.Joins("Life").Select("Life.title", "name", "life_comment.created_at", "life_comment.content").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("lid=?", id).Order("created_at Desc").Find(&lifecomment).Error
	if err != nil {
		fmt.Println(err)
		return nil, errmsg.ERROR
	}
	return lifecomment, errmsg.SUCCESS
}
