package lib

import (
	"fmt"

	"github.com/Qesy/qesydb"
	"github.com/Qesy/qesygo"
)

// Config 系统配置文件
type Config struct {
	Db        Db
	Cache     Cache
	Framework Framework
}

// Db 数据库配置
type Db struct {
	//Name    string `json:"username"`
	Driver   string
	Host     string
	Name     string
	User     string
	Password string
	Port     string
	Charset  string
}

// Cache Redis配置
type Cache struct {
	Driver string
	Host   string
	Key    string
	Num    string
	Port   string
	Auth   string
}

// Framework 框架配置
type Framework struct {
	Name    string
	Version string
	Auth    string
	Qq      string
	Email   string
	Key     string
}

var confRs Config

//RedisCr Redis进程池
var RedisCr qesygo.CacheRedis

//Key 网站加密种子
var Key string

func init() {
	conf := getConf()
	Key = conf.Framework.Key
	RedisCr.Conninfo = conf.Cache.Host + ":" + conf.Cache.Port
	RedisCr.Auth = conf.Cache.Auth
	if err := RedisCr.Connect(); err != nil {
		qesygo.Die("Redis Connect Err : " + err.Error())
	}
	fmt.Println("Redis Connection Success !")
	if err := qesydb.Connect(conf.Db.User + ":" + conf.Db.Password + "@tcp(" + conf.Db.Host + ":" + conf.Db.Port + ")/" + conf.Db.Name + "?charset=" + conf.Db.Charset); err != nil {
		qesygo.Die("Mysql Connect Err : " + err.Error())
	}
	fmt.Println("Mysql Connection Success !")
	fmt.Println("Service Start Success !")
}

// getConf 获取配置文件
func getConf() Config {
	if str, err := qesygo.ReadFile("conf.ini"); err == nil {
		if err := qesygo.JsonDecode(str, &confRs); err != nil {
			fmt.Printf("config reload error: %s", err)
		}
	} else {
		qesygo.Die(err)
	}
	return confRs

}
