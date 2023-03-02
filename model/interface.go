package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

// 服务层
type UserService interface {
	Signup(ctx *gin.Context, u *User) error
	Signin(ctx *gin.Context, u *User) (string, error)
	Signout(ctx *gin.Context, u *User) error
}

type TokenService interface {
	NewTokenForUser(ctx *gin.Context, u *User) (*Token, error)
	DeletUserToken(ctx *gin.Context, u *User) error
	ValidateIDToken(tokenString string) (*User, error)
}

type GoodsService interface {
	AddGoods(ctx *gin.Context, g *Goods) error
	ListGoods(ctx *gin.Context) error
	Preheat(ctx *gin.Context, g *Goods) error
	Seckill(ctx *gin.Context, username string, goods *Goods) error
	StartSeckillConsumer()
}

// 仓库层
type MysqlRepository interface {
	CreateUser(ctx *gin.Context, u *User) error
	CreateOrder(o *Order) error
	CreatGoods(ctx *gin.Context, g *Goods) error
	SearchGoods(ctx *gin.Context, name string) (*Goods, error)
	SearchUser(ctx *gin.Context, name string) (*User, error)
	SearchAllGoods(ctx *gin.Context, name string) (*Goods, error)
}

type RedisRepository interface {
	StoreToken(ctx *gin.Context, token *Token, expiresIn time.Duration) error
	DeleteToken(ctx *gin.Context, userID string) error
	PreHeat(ctx *gin.Context, g *Goods) error
	Seckill(ctx *gin.Context, userName string, goodsName string) error
}

type RabbitMQRepository interface {
	Push(data string) error
	Getch() *amqp.Channel
	Getqname() string
}
