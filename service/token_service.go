package service

import (
	"seckill/model"

	"github.com/gin-gonic/gin"
)

type tokenService struct {
	TokenRepository model.RedisRepository
}

func NewTokenService(r model.RedisRepository) model.TokenService {
	return &tokenService{
		TokenRepository: r,
	}
}

func (s *tokenService) NewTokenForUser(ctx *gin.Context, u *model.User) (*model.Token, error) {
	return nil, nil
}
func (s *tokenService) DeletUserToken(ctx *gin.Context, u *model.User) error {
	return nil
}
func (s *tokenService) ValidateIDToken(tokenString string) (*model.User, error) {
	return nil, nil
}
