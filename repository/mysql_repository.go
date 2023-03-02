package repository

import (
	"errors"
	"fmt"
	"seckill/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type mysqlRepository struct {
	DB *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) model.MysqlRepository {
	return &mysqlRepository{
		DB: db,
	}
}

func (m *mysqlRepository) CreateUser(ctx *gin.Context, u *model.User) error {
	if m.CheckUser(u.Username) {
		fmt.Println("[mysql] user already exists")
		return errors.New("user already exists")
	}
	err := m.DB.Create(&u).Error
	if err != nil {
		fmt.Println("[mysql] create user failed: ", err)
	}
	fmt.Println("[mysql] create user succeeded: ", u.Username)
	return nil
}

func (m *mysqlRepository) CheckUser(name string) bool {
	var user model.User
	m.DB.Select("id").Where("username = ?", name).First(&user)
	return user.Id > 0
}

func (m *mysqlRepository) SearchUser(ctx *gin.Context, name string) (*model.User, error) {
	var user model.User
	m.DB.Where("username = ?", name).First(&user)
	if user.Id > 0 {
		return &user, nil
	} else {
		fmt.Println("[mysql] no such user")
		return nil, errors.New("user not found")
	}
}

func (m *mysqlRepository) CreateOrder(o *model.Order) error {
	err := m.DB.Create(&o).Error
	if err != nil {
		fmt.Println("[mysql] create order failed: ", err)
	}
	fmt.Println("[mysql] create order succeeded: ", o.ID)
	return nil

}
func (m *mysqlRepository) CreatGoods(ctx *gin.Context, g *model.Goods) error {
	err := m.DB.Create(&g).Error
	if err != nil {
		fmt.Println("[mysql] create goods failed: ", err)
	}
	fmt.Println("[mysql] create goods succeeded: ", g.GoodsName)
	return nil

}
func (m *mysqlRepository) SearchGoods(ctx *gin.Context, name string) (*model.Goods, error) {
	var goods model.Goods
	m.DB.Where("goods_name = ?", name).First(&goods)
	if goods.Id > 0 {
		return &goods, nil
	} else {
		fmt.Println("[mysql] no such goods")
		return nil, errors.New("no such goods")
	}
}

func (m *mysqlRepository) SearchAllGoods(ctx *gin.Context, name string) (*model.Goods, error) {
	var goods model.Goods
	m.DB.Where("goods_name = ?", name).First(&goods)
	if goods.Id == 0 {
		return nil, errors.New("no such goods")
	} else {
		return &goods, nil
	}
}
