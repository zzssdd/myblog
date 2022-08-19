package api

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/utils/errmsg"
	"net/http"
	"strconv"
)

func AddProject(c *gin.Context) {
	var data model.Project
	_ = c.ShouldBindJSON(&data)
	code := model.CreateProject(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetProjectInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code, data := model.GetProjectInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func DeleteProject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteProject(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func EditProject(c *gin.Context) {
	var data model.Project
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.EditProject(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetProject(c *gin.Context) {
	data, code := model.GetProject()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
