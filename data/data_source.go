package data

import (
	"fmt"
	"seckill/config"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type DataSource struct {
	Db          *gorm.DB
	RedisClient *redis.Client
	Rabbitmq    *Rabbitmq
}

func InitData() (*DataSource, error) {
	config, err := config.GetAppConfig()
	if err != nil {
		panic("failed to load data config: " + err.Error())
	}
	db := initMysql(config)
	r := initRedis(config)
	mq := initRabbitMQ(config)

	return &DataSource{
		Db:          db,
		RedisClient: r,
		Rabbitmq:    mq,
	}, nil
}

func (d *DataSource) Close() error {
	dbInstance, _ := d.Db.DB()
	if err := dbInstance.Close(); err != nil {
		return fmt.Errorf("error closing Postgresql: %w", err)
	}

	if err := d.RedisClient.Close(); err != nil {
		return fmt.Errorf("error closing Redis Client: %w", err)
	}
	// if err := d.Rabbitmq.Close(); err != nil {
	// 	return fmt.Errorf("error closing Cloud Storage client: %w", err)
	// }

	return nil
}
