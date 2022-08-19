package api

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/utils/errmsg"
	"net/http"
	"strconv"
)

func AddBoard(c *gin.Context) {
	var data model.Board
	_ = c.ShouldBindJSON(&data)
	code := model.CreatBoard(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func DeleteBoard(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteBoard(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetBoard(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	data, code := model.GetBoard(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
