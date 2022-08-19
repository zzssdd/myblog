package api

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/utils/errmsg"
	"net/http"
	"strconv"
)

func AddLcomment(c *gin.Context) {
	var data model.LifeComment
	_ = c.ShouldBindJSON(&data)
	code := model.CreateLComment(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func DeleteLcomment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteLcomment(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetLcomment(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	data, code := model.GetLcomment(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetLifecomment(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetLifeComment(pageSize, pageNum, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
