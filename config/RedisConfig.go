package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
)

// redis信息
type RedisInfo struct {
	Hostname string `yaml:"hostname"`
	Database string `yaml:"database"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
}

// 区分redis环境
type RedisENV struct {
	Debug   RedisInfo
	Test    RedisInfo
	Release RedisInfo
}

func (rn *RedisENV) GetInfo() RedisInfo {
	// 获取操作系统类型
	ostype := runtime.GOOS

	// 获取当前目录的路径
	dir, _ := os.Getwd()

	// 定义filePath变量，存储数据库配置文件路径
	var filePath string

	//兼容Linux和windows的文件路径
	if ostype == "windows" {
		filePath = dir + "\\config\\redis.yaml"
	} else if ostype == "linux" {
		filePath = dir + "/config/redis.yaml"
	}

	// 读取数据库配置文件
	if fileData, err := ioutil.ReadFile(filePath); err != nil {
		fmt.Println(err)
	} else {
		//解析yaml文件内容
		err := yaml.Unmarshal(fileData, &rn)
		if err != nil {
			log.Fatalf("cannot unmarshal data: %v", err)
		}

		// 返回当前环境的数据库信息
		if gin.Mode() == gin.DebugMode {
			return rn.Debug
		} else if gin.Mode() == gin.TestMode {
			return rn.Test
		} else {
			return rn.Release
		}
	}

	// 默认返回debug数据库信息
	return rn.Debug
}

var RDB *redis.Client

func (ri *RedisInfo) GetRedisClient() {
	// ri.Database为string类型，需要转换为int类型
	db, err := strconv.Atoi(ri.Database)
	if err == nil {
		fmt.Println("err", err)
	}

	// 连接redis
	RDB = redis.NewClient(&redis.Options{
		Addr:     ri.Hostname + ":" + ri.Port,
		Password: ri.Password, // no password set
		DB:       db,          // use default DB
	})

}
