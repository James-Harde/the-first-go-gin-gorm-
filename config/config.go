// 用来读取config.yml配置文件中的内容的
package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Host     string
		Port     int
		Name     string
		Username string
		Password string
	}
}

// 为什么这么写？*Config是一个指针类型（指针类型返回的是地址，要用*解引用才能得到地址存放的值：eg *p=10,i=10,p为 *int类型放的是int类型的地址，*解引用之后即得到存放的10而i为int类型，存放10 ）；
// 指向Config结构体的实例，这样可以在程序中方便地访问和修改配置文件中的内容。
var AppConfig *Config // AppConfig是一个全局变量，用来存储配置文件中的内容。

func InitConfig() {
	//用viper去读取config.yml配置文件中的内容
	// 1. 设置配置文件基础信息
	viper.SetConfigName("config")   // 配置文件名
	viper.SetConfigType("yml")      // 配置文件类型
	viper.AddConfigPath("./config") // 配置文件路径
	// 读文件、打开文件、viper 路径：./ 基于【终端执行命令的文件夹】（运行目录）
	// import 导入本地包：./ 基于【当前.go 源码文件所在文件夹】

	// 2. 读取磁盘上的yml配置文件
	// 读取配置文件，如果读取失败，打印错误信息
	if err := viper.ReadInConfig(); err != nil {
		// panic(err)
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 3. 初始化结构体指针，接收配置映射数据
	AppConfig = &Config{} // 初始化 AppConfig 变量，创建一个 Config 结构体的实例，并将其地址赋值给 AppConfig.空的
	//把读取到的配置文件内容，反序列化到AppConfig结构体中
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("反序列化配置文件失败: %v", err)
	}
}
