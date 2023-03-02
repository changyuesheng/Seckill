package service

import (
	"errors"
	"fmt"
	middleware "seckill/middlerware"
	"seckill/model"
	"seckill/utils"

	"github.com/gin-gonic/gin"
)

type userService struct {
	MysqlRepository model.MysqlRepository
	RedisRepository model.RedisRepository
}

func NewUserService(m model.MysqlRepository, r model.RedisRepository) model.UserService {
	return &userService{
		MysqlRepository: m,
		RedisRepository: r,
	}
}

// 注册服务
func (s *userService) Signup(ctx *gin.Context, u *model.User) error {
	pw, err := utils.HashPassword(u.Password)
	if err != nil {
		fmt.Println("Unable to hash password ", err)
		return err
	}
	u.Password = pw
	if err := s.MysqlRepository.CreateUser(ctx, u); err != nil {
		return err
	}
	return nil
}

// 登录服务
func (s *userService) Signin(ctx *gin.Context, user *model.User) (string, error) {
	fuser, err := s.MysqlRepository.SearchUser(ctx, user.Username)
	if err != nil {
		return "", err
	}
	match, err := utils.ComparePasswords(fuser.Password, user.Password)
	if err != nil {
		return "", err
	}
	if !match {
		return "", errors.New("wrong password")
	}
	token, err := middleware.GenerateToken(ctx, fuser)
	if err != nil {
		return "", err
	}
	*user = *fuser
	return token, nil
}

// 注销服务
func (s *userService) Signout(ctx *gin.Context, u *model.User) error {
	return nil
}


