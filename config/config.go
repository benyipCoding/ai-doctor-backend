package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 对应 config.yaml 的顶级结构
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

var GlobalConfig Config

func InitConfig() {
	viper.SetConfigName("config") // 文件名 (不需要后缀)
	viper.SetConfigType("yaml")   // 文件格式
	viper.AddConfigPath(".")      // 在当前目录查找

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %w", err))
	}

	// 将配置映射到结构体
	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		panic(fmt.Errorf("解析配置文件失败: %w", err))
	}
}
