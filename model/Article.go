package model

import (
	"fmt"
	"gorm.io/gorm"
	"myblog/utils/errmsg"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title     string `gorm:"type:varchar(100);not null" json:"title"`
	Cid       int    `gorm:"type:int;not null" json:"cid"`
	Desc      string `gorm:"type:varchar(200);not null" json:"desc"`
	Content   string `gorm:"type:longtext;not null" json:"content"`
	Img       string `gorm:"type:varchar(100);not null" json:"img"`
	ReadCount int    `gorm:"type:int;not null;default:0" json:"readCount"`
}

func CreateArticle(data *Article) int {
	var category Category
	err := db.Create(&data).Error
	err = db.Model(&category).Where("id=?", data.Cid).UpdateColumn("num", gorm.Expr("num+1")).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteArticle(id int) int {
	var art Article
	err := db.Where("id=?", id).Delete(&art).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func UpdateArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&article).Where("id=?", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetArticleInfo(id int) (Article, int) {
	var article Article
	err = db.Preload("Category").Where("id=?", id).First(&article).Error
	err = db.Model(&article).Where("id=?", id).UpdateColumn("read_count", gorm.Expr("read_count+1")).Error
	if err != nil {
		return article, errmsg.ERROR_ARI_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

func GetArticleList() ([]Article, int) {
	var articleList []Article

	err = db.Select("article.id", "title", "img", "updated_at", "created_at", "desc", "read_count", "Category.name").Order("article.created_at DESC").Joins("Category").Find(&articleList).Error
	if err != nil {
		fmt.Println(err)
		return nil, errmsg.ERROR
	}
	return articleList, errmsg.SUCCESS
}

func SearchArticle(title string) ([]Article, int) {
	var articleList []Article
	err = db.Select("title", "img", "created_at", `desc`, "read_count", "Category.name").Order("Created_At").Joins("Category").Where("title LIKE ?", title+"%").Find(&articleList).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return articleList, errmsg.SUCCESS
}

func GetCateArt(name string, pageSize int, pageNum int) ([]Article, int) {
	var cateArtList []Article

	err = db.Joins("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("Category.name=?", name).Find(&cateArtList).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST
	}
	return cateArtList, errmsg.SUCCESS
}
