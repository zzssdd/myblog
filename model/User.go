package model

import (
	"myblog/utils"
	"myblog/utils/errmsg"
)

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func CheckLogin(username string, password string) int {
	if username != utils.Config.Username {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if password != utils.Config.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCESS
}
