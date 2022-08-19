package main

import (
	"myblog/model"
	"myblog/router"
)

func main() {
	model.InitDb()
	router.InitRouter()
}
