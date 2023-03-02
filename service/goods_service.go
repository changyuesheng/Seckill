package service

import (
	"errors"
	"fmt"
	"seckill/model"
	"strings"

	"github.com/gin-gonic/gin"
)

type goodsService struct {
	MysqlRepository  model.MysqlRepository
	RedisRepository  model.RedisRepository
	RabbitRepository model.RabbitMQRepository
}

func NewGoodsService(m model.MysqlRepository, r model.RedisRepository, mq model.RabbitMQRepository) model.GoodsService {
	return &goodsService{
		MysqlRepository:  m,
		RedisRepository:  r,
		RabbitRepository: mq,
	}
}

func (s *goodsService) AddGoods(ctx *gin.Context, goods *model.Goods) error {
	_, err := s.MysqlRepository.SearchGoods(ctx, goods.GoodsName)
	if err == nil {
		return errors.New("goods already exists")
	}
	err = s.MysqlRepository.CreatGoods(ctx, goods)
	if err != nil {
		return err
	}
	return nil
}
func (s *goodsService) ListGoods(ctx *gin.Context) error {
	return nil
}

// 预热商品
func (s *goodsService) Preheat(ctx *gin.Context, goods *model.Goods) error {
	fgoods, err := s.MysqlRepository.SearchGoods(ctx, goods.GoodsName)
	if err != nil {
		fmt.Println("no such goods")
		return err
	}
	userId := ctx.MustGet("userid").(int)
	if userId != fgoods.MerchantId {
		fmt.Println("this goods does not belong to you")
		return errors.New("this goods does not belong to you")
	}
	err = s.RedisRepository.PreHeat(ctx, fgoods)
	if err != nil {
		fmt.Println("preheat error: ", err)
		return err
	}
	return nil
}

func (s *goodsService) Seckill(ctx *gin.Context, username string, goods *model.Goods) error {

	// ---用户抢优惠券。后面需要高并发处理---
	// 先在缓存执行原子性的秒杀操作。将原子性地完成"判断能否秒杀-执行秒杀"的步骤
	err := s.RedisRepository.Seckill(ctx, username, goods.GoodsName)
	if err != nil {
		return err
	}
	_ = s.RabbitRepository.Push(username + "." + goods.GoodsName)
	return nil
}

func (s *goodsService) StartSeckillConsumer() {
	c := s.RabbitRepository.Getch()
	qname := s.RabbitRepository.Getqname()

	msgs, err := c.Consume(
		qname,     // queue
		"toutiao", // consumer
		true,      // auto-ack,true消费了就消失
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		fmt.Println("Failed to register a consumer")
	}
	for d := range msgs {
		temp := strings.Split(string(d.Body), ".")
		order := &model.Order{
			Username:  temp[0],
			GoodsName: temp[1],
		}
		err := s.MysqlRepository.CreateOrder(order)
		if err != nil {
			fmt.Println("write data to database failed")
		}

	}
}
