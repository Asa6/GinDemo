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
	"time"
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
type DBEnv struct {
	Debug   DBInfo
	Test    DBInfo
	Release DBInfo
}

//获取数据库配置
func (d *DBEnv) GetInfo() DBInfo {
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

// 声明公共的数据库连接句柄
var DB *gorm.DB

func (d *DBInfo) GetConnect() {
	// 数据库连接信息格式化
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", d.Username, d.Password, d.Hostname, d.Port, d.Database)

	var err error
	// 连接数据库
	DB, err = gorm.Open(d.Datatype, connArgs)
	if err != nil {
		log.Fatalf("数据库连接失败：%v", err)
	}

	// 报错： invalid connection问题
	// 解决： 设置一个连接被使用的最长时间，即过了一段时间后会被强制回收，理论上这可以有效减少不可用连接出现的概率。当数据库方面也设置了连接的超时时间时(mysql默认8小时)，这个值应当不超过数据库的超时参数值。
	DB.DB().SetConnMaxIdleTime(time.Minute)

	// 全局禁用表复数
	DB.SingularTable(true)

	// 自动迁移
	DB.AutoMigrate(&User{})
}

// 已开启自动迁移，无需调用此func
func CreateModel() {
	// 检查模型User表是否存在，不存在则创建此表
	if !DB.HasTable(&User{}) {
		err := DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{}).Error
		if err != nil {
			log.Fatalf("模型User表创建失败：%v", err)
		}
	}
}
