package api

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/utils/errmsg"
	"net/http"
	"strconv"
)

func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := model.CreateCategory(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"message:": errmsg.GetErrMsg(code),
	})
}

func DelCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DelteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code := model.EditCategory(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetCategoryId(c *gin.Context) {
	name := c.Param("name")
	code, id := model.GetCategoryId(name)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"id":      id,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetCategory(c *gin.Context) {
	data, code := model.GetCategory()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
