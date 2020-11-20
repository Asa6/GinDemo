package config

import (
	. "GinDemo/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

// 数据库信息
type DBInfo struct {
	Datatype string `yaml:"datatype"`
	Hostname string `yaml:"hostname"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
}

// 区分数据库环境
type Database struct {
	Debug   DBInfo
	Test    DBInfo
	Release DBInfo
}

//获取数据库配置
func (d *Database) GetInfo() DBInfo {
	// 获取操作系统类型
	ostype := runtime.GOOS

	// 获取当前目录的路径
	dir, _ := os.Getwd()

	// 定义filePath变量，存储数据库配置文件路径
	var filePath string

	//兼容Linux和windows的文件路径
	if ostype == "windows" {
		filePath = dir + "\\config\\database.yaml"
	} else if ostype == "linux" {
		filePath = dir + "/config/database.yaml"
	}

	// 读取数据库配置文件
	if fileData, err := ioutil.ReadFile(filePath); err != nil {
		fmt.Println(err)
	} else {
		//解析yaml文件内容
		err := yaml.Unmarshal(fileData, &d)
		if err != nil {
			log.Fatalf("cannot unmarshal data: %v", err)
		}

		// 返回当前环境的数据库信息
		if gin.Mode() == gin.DebugMode {
			return d.Debug
		} else if gin.Mode() == gin.TestMode {
			return d.Test
		} else {
			return d.Release
		}
	}

	// 默认返回debug数据库信息
	return d.Debug
}

func (d *DBInfo) GetConnect() *gorm.DB {
	// 数据库连接信息格式化
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", d.Username, d.Password, d.Hostname, d.Port, d.Database)

	// 连接数据库
	//var err error
	db, err := gorm.Open(d.Datatype, connArgs)
	if err != nil {
		log.Fatalf("数据库连接失败：%v", err)
	}

	// 全局禁用表复数
	db.SingularTable(true)
	return db
}

func CreateModel(db *gorm.DB) {
	// 检查模型User表是否存在，不存在则创建此表
	if !db.HasTable(&User{}) {
		err := db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{}).Error
		if err != nil {
			log.Fatalf("模型User表创建失败：%v", err)
		}
	}
}
