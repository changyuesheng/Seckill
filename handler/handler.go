package handler

import (
	"fmt"
	"seckill/data"
	"seckill/model"
	"seckill/repository"
	"seckill/service"
)

type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
	GoodsService model.GoodsService
}

func NewHandler(d *data.DataSource) *Handler {
	fmt.Println("---Create New Handler---")

	mysqlReposiroty := repository.NewMysqlRepository(d.Db)
	redisRepository := repository.NewRedisRepository(d.RedisClient)
	rabbitReposirtoy := repository.NewRabbitRepository(d.Rabbitmq)

	userService := service.NewUserService(mysqlReposiroty, redisRepository)
	goodsService := service.NewGoodsService(mysqlReposiroty, redisRepository, rabbitReposirtoy)
	tokenService := service.NewTokenService(redisRepository)

	go goodsService.StartSeckillConsumer()

	return &Handler{
		UserService:  userService,
		TokenService: tokenService,
		GoodsService: goodsService,
	}
}
