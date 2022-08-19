package model

import (
	"gorm.io/gorm"
	"myblog/utils/errmsg"
)

type Board struct {
	gorm.Model
	Name    string `gorm:"varchar(20);not null;" json:"name"`
	Content string `gorm:"varchar(500);not null;" json:"content"`
}

func CreatBoard(data *Board) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteBoard(id int) int {
	var board Board
	err := db.Where("id=?", id).Delete(&board).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetBoard(pageSize int, pageNum int) ([]Board, int) {
	var board []Board
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_at DESC").Find(&board).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return board, errmsg.SUCCESS
}
