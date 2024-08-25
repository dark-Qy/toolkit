package main

import (
	"fmt"
	"os"
	"toolkit/config"
	"toolkit/dao"
	"toolkit/models"
	"toolkit/routers"
	"toolkit/setting"
)

// 设置基础配置
const defaultConfFile = "./config/config.ini"

func main() {
	fmt.Println("hello world")
	confFile := defaultConfFile
	if len(os.Args) > 2 {
		fmt.Println("use specified conf file: ", os.Args[1])
		confFile = os.Args[1]
	} else {
		fmt.Println("no configuration file was specified, use ./conf/config.ini")
	}
	// 首先获取配置文件
	config.InitConfig()
	if err := setting.Init(confFile); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	// 连接数据库
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
		return
	}
	defer dao.Close() // 程序退出关闭数据库连接
	// 根据模型创建数据库表项
	dao.DB.AutoMigrate(&models.User{})

	// 注册路由
	r := routers.SetupRouter()

	// 启动服务器
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
