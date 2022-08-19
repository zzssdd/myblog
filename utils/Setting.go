package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var Config = new(AppConf)

type AppConf struct {
	ServerConfig `ini:"server"`
	DbConfig     `ini:"database"`
}
type ServerConfig struct {
	AppMode  string `ini:"AppMode"`
	HttpPort string `ini:"HttpPort"`
	JwtKey   string `ini:"JwtKey"`
	Username string `ini:"Username"`
	Password string `ini:"Password""`
}

type DbConfig struct {
	Db         string `ini:"Db"`
	DbHost     string `ini:"DbHost"`
	DbPort     string `ini:"DbPort"`
	DbUser     string `ini:"DbUser"`
	DbPassword string `ini:"DbPassword"`
	DbName     string `ini:"DbName"`
}

func init() {
	err := ini.MapTo(Config, "config/config.ini")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
