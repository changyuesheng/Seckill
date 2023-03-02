package router

import (
	"seckill/data"
	"seckill/handler"
	middleware "seckill/middlerware"

	"github.com/gin-gonic/gin"
)

func InitRouter(d *data.DataSource) (*gin.Engine, error) {
	r := gin.Default()
	//r.Use(gin.Recovery())
	//r.Use(middleware.Logger())
	h := handler.NewHandler(d)
	auth := r.Group("v1")
	auth.Use(middleware.JwtToken())
	{
		// TODO:用户模块路由接口
		auth.POST("user/signout", h.Signout)
		// TODO:商品模块路由接口
		auth.PATCH("/goods/seckill", h.Seckill)
		auth.POST("/goods/heat", h.Preheat)
		auth.POST("/goods/add", h.AddGoods)
		auth.PATCH("/goods/preheat", h.Preheat)
	}
	router := r.Group("v1")
	{
		// 用户模块
		router.POST("user/signup", h.Signup)
		router.POST("user/signin", h.Signin)
		router.GET("/test", h.Welcome)
		// 商品模块
		router.GET("/goods/list", h.GetGoods)

	}
	return r, nil
}
