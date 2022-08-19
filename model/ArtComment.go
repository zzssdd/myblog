package model

import (
	"fmt"
	"gorm.io/gorm"
	"myblog/utils/errmsg"
)

type ArtComment struct {
	Article Article `gorm:"foreignkey:Aid"`
	gorm.Model
	Name    string `gorm:"type:varchar(20);not null" json:"name"`
	Content string `gorm:"type:varchar(500);not null" json:"content"`
	Aid     int    `gorm:"type:int;not null" json:"aid"`
}

func CreateComment(data *ArtComment) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteComment(id int) int {
	var comment ArtComment
	err := db.Where("id=?", id).Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetCommentList(pageSize int, pageNum int, id int) ([]ArtComment, int) {
	var artcomment []ArtComment
	err = db.Joins("Article").Select("name", "art_comment.created_at", "art_comment.content").Where("aid=?", id).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Find(&artcomment).Error
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return nil, errmsg.ERROR
	}
	return artcomment, errmsg.SUCCESS
}

func GetArtCommentList() ([]ArtComment, int) {
	var artcomment []ArtComment
	err = db.Select("art_comment.id", "name", "art_comment.content", "Article.title").Joins("Article").Order("art_comment.created_at DESC").Find(&artcomment).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return artcomment, errmsg.SUCCESS
}
