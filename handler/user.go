package handler

import (
	"fmt"
	"net/http"
	"seckill/model"
	"seckill/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// 用户注册
func (h *Handler) Signup(ctx *gin.Context) {
	var user *model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR,
			"message": err.Error(),
		})
		return
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		fmt.Println("validate failed ", err)
		ctx.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR,
			"message": err,
		})
		return
	}
	err := h.UserService.Signup(ctx, user)
	if err != nil {
		fmt.Println("signup failed ", err)
		ctx.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  utils.SUCCESS,
		"message": "signup succeeded",
	})
}

// 用户登录
func (h *Handler) Signin(ctx *gin.Context) {
	var user *model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR,
			"message": err.Error(),
		})
		return
	}
	token, err := h.UserService.Signin(ctx, user)
	if err != nil {
		fmt.Println("signin failed: ", err)
		ctx.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": utils.SUCCESS,
		"tokens": token,
	})

}

// 用户退出
func (h *Handler) Signout(ctx *gin.Context) {

}
