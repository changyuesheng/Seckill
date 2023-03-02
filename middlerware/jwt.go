package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"seckill/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	Username string
	Passowrd string
	Uid      int
	Role     int
	jwt.StandardClaims
}

var JwtKey []byte = []byte("89sh82js784254262sdgwsva")

// 生成用户token
func GenerateToken(ctx *gin.Context, u *model.User) (string, error) {
	expireTime := time.Now().Add(3 * time.Hour)
	claims := MyClaims{
		Username: u.Username,
		Passowrd: u.Password,
		Uid:      u.Id,
		Role:     u.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireTime.Unix(),
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		fmt.Println("generate token error: ", err)
		return "", err
	} else {
		return token, nil
	}
}

// 验证用户token
func ValidateToken(token string) (*MyClaims, error) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, nil
	} else {
		return nil, errors.New("token is invalid")
	}
}

// JWT 中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    "",
				"message": "token is null",
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"code":    "",
				"message": "token format is wrong",
			})
			c.Abort()
			return
		}
		key, err := ValidateToken(checkToken[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    "",
				"message": "token is invalid",
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			c.JSON(http.StatusOK, gin.H{
				"code":    "",
				"message": "token expired",
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Set("userid", key.Uid)
		c.Set("userrole", key.Role)
		c.Next()
	}
}
