package model

import (
	"fmt"
	"myblog/utils/errmsg"
)

type Category struct {
	ID   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	Num  int    `gorm:"type:int;default:0"json:"num"`
}

func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DelteCategory(id int) int {
	var category Category
	err = db.Where("id=?", id).Delete(&category).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func EditCategory(id int, data *Category) int {
	var category Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Model(&category).Where("id=?", id).Updates(maps).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetCategoryId(name string) (int, int) {
	var category Category

	err = db.Select("ID").Where("name=?", name).First(&category).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR, 0
	}
	return errmsg.SUCCESS, category.ID
}

func GetCategory() ([]Category, int) {
	var category []Category
	err = db.Find(&category).Error
	if err != nil {
		fmt.Println(err)
		return nil, errmsg.ERROR
	}
	return category, errmsg.SUCCESS
}
