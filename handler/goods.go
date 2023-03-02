package handler

import (
	"net/http"
	"seckill/model"
	"seckill/utils"

	"github.com/gin-gonic/gin"
)

// 查询商品
func (h *Handler) GetGoods(ctx *gin.Context) {

}

// 秒杀商品
func (h *Handler) Seckill(ctx *gin.Context) {
	var goods *model.Goods
	if err := ctx.ShouldBindJSON(&goods); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR,
			"message": err.Error(),
		})
		return
	}
	userRole := ctx.MustGet("userrole").(int)
	if userRole != 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR,
			"message": "Only consumer can get goods",
		})
		return
	}
	userName := ctx.MustGet("username").(string)
	err := h.GoodsService.Seckill(ctx, userName, goods)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR,
			"message": "seckill failed",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  utils.SUCCESS,
		"message": "seckill succeeded",
	})
}

// 添加商品
func (h *Handler) AddGoods(ctx *gin.Context) {
	var goods *model.Goods
	if err := ctx.ShouldBindJSON(&goods); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR,
			"message": err.Error(),
		})
		return
	}
	userRole := ctx.MustGet("userrole").(int)
	if userRole != 2 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR,
			"message": "Only sellers can create goods",
		})
		return
	}
	userId := ctx.MustGet("userid").(int)
	goods.MerchantId = userId
	if err := h.GoodsService.AddGoods(ctx, goods); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR,
			"message": "add goods failed",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  utils.SUCCESS,
		"message": "add goods succeeded",
	})
}

// 预热商品
func (h *Handler) Preheat(ctx *gin.Context) {
	var goods *model.Goods
	if err := ctx.ShouldBindJSON(&goods); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR,
			"message": err.Error(),
		})
		return
	}
	userRole := ctx.MustGet("userrole").(int)
	if userRole != 2 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR,
			"message": "Only sellers can preheat goods",
		})
		return
	}
	userId := ctx.MustGet("userid").(int)
	goods.MerchantId = userId
	if err := h.GoodsService.Preheat(ctx, goods); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR,
			"message": "add goods failed",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  utils.SUCCESS,
		"message": "preheat goods succeeded",
	})
}
