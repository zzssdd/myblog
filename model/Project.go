package model

import (
	"gorm.io/gorm"
	"myblog/utils/errmsg"
)

type Project struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"json:"name"`
	Desc string `gorm:"type:varchar(50);not null"json:"desc"`
	Url  string `gorm:"type:varchar(50);not null"json:"url"`
	Img  string `gorm:"type:varchar(100);not null"json:"img"`
}

func CreateProject(data *Project) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetProjectInfo(id int) (int, Project) {
	var project Project
	err := db.Where("id=?", id).First(&project).Error
	if err != nil {
		return errmsg.ERROR, project
	}
	return errmsg.SUCCESS, project
}

func DeleteProject(id int) int {
	var project Project
	err := db.Where("id=?", id).Delete(&project).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func EditProject(id int, data *Project) int {
	var project Project
	var maps = make(map[string]interface{})
	maps["Name"] = data.Name
	maps["Desc"] = data.Desc
	maps["Url"] = data.Url
	maps["Img"] = data.Img
	err := db.Model(&project).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetProject() ([]Project, int) {
	var project []Project
	err := db.Find(&project).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return project, errmsg.SUCCESS
}
