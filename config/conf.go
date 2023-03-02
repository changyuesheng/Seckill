package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type AppConfig struct {
	AppMode        string `ini:"AppMode"`
	HttpPort       string `ini:"HttpPort"`
	DbConfig       `ini:"mysql"`
	RedisConfig    `ini:"redis"`
	RabbitMQConfig `ini:"rabbitmq"`
	// EtcdConfig     `ini:"etcd"`
}

type DbConfig struct {
	Dbtype   string `ini:"DbType"`
	DbName   string `ini:"DbName"`
	DbHost   string `ini:"DbHost"`
	DbPort   string `ini:"DbPort"`
	DbUser   string `ini:"DbUser"`
	Password string `ini:"DbPassword"`
}

type RedisConfig struct {
	RedisHost     string `ini:"RedisHost"`
	RedisPort     string `ini:"RedisPort"`
	RedisName     string `ini:"RedisName"`
	RedisPassowrd string `ini:"RedisPassowrd"`
	Databasenum   int    `ini:"RedisDbNum"`
}

type RabbitMQConfig struct {
	RabbitMQHost     string `ini:"RabbitMQHost"`
	RabbitMQPort     string `ini:"RabbitMQPort"`
	RabbitMQUser     string `int:"RabbitMQUser"`
	RabbitMQPaasword string `ini:"RabbitMQPassword"`
	QueueName        string `ini:"QueueName"`
}

// type EtcdConfig struct {
// 	EtcdHost string `ini:"EtcdHost"`
// 	EtcdPort string `ini:"EtcdPort"`
// }

// 获取系统配置文件
func GetAppConfig() (appConfig AppConfig, err error) {

	var configObj = AppConfig{
		RabbitMQConfig: RabbitMQConfig{},
		DbConfig:       DbConfig{},
		RedisConfig:    RedisConfig{},
		// EtcdConfig:     EtcdConfig{},
	}
	err = ini.MapTo(&configObj, "./config.ini")
	if err != nil {
		logrus.Error("Load config failed: ", err)
		panic(err)
	}
	return configObj, nil
}
