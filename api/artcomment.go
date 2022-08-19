package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/utils/errmsg"
	"net/http"
	"strconv"
)

func AddComment(c *gin.Context) {
	var data model.ArtComment
	_ = c.ShouldBindJSON(&data)
	fmt.Println(data)
	code := model.CreateComment(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteComment(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetCommentList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetCommentList(pageSize, pageNum, id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetArtCommentList(c *gin.Context) {
	data, code := model.GetArtCommentList()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetArtComment(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi("pagenum")
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetCommentList(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
