package repository

import (
	"errors"
	"seckill/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type redisRepository struct {
	RClient   *redis.Client
	scriptSHA string
}

func NewRedisRepository(client *redis.Client) model.RedisRepository {
	res := &redisRepository{
		RClient: client,
	}
	res.scriptSHA = res.PreLoadScript(secKillScript)
	return res
}

func (r *redisRepository) StoreToken(ctx *gin.Context, token *model.Token, expiresIn time.Duration) error {
	return nil

}
func (r *redisRepository) DeleteToken(ctx *gin.Context, userID string) error {
	return nil
}

func (r *redisRepository) PreHeat(ctx *gin.Context, goods *model.Goods) error {
	key := goods.GoodsName
	fields := map[string]interface{}{
		"id":        goods.Id,
		"merchant":  goods.MerchantId,
		"goodsname": goods.GoodsName,
		"stock":     goods.Stock,
		"State":     goods.State,
	}
	err := r.RClient.HMSet(key, fields).Err()
	return err
}

// func CacheCoupon(coupon model.Coupon) (string, error) {
// 	key := getCouponKeyByCoupon(coupon)
// 	fields := map[string]interface{}{
// 		"id": coupon.Id,
// 		"username": coupon.Username,
// 		"couponName": coupon.CouponName,
// 		"amount": coupon.Amount,
// 		"left": coupon.Left,
// 		"stock": coupon.Stock,
// 		"description": coupon.Description,
// 	}
// 	val, err := data.SetMapForever(key, fields)
// 	return val, err
// }

func (r *redisRepository) Seckill(ctx *gin.Context, userName string, goodsName string) error {
	evalres, err := r.RClient.EvalSha(r.scriptSHA, []string{userName, goodsName, goodsName}).Result()
	if err != nil {
		return errors.New("script eval failed1")
	}
	res, ok := evalres.(int64)
	if !ok {
		return errors.New("script eval failed2")
	}
	switch {
	case res == -1:
		return errors.New("user already has goods")
	case res == -2:
		return errors.New("no such goods")
	case res == -3:
		return errors.New("no goods left ")
	case res == 1: // left为0时, 就是存量为0, 那就是没抢到, 也可能原本为1, 抢完变成了0.
		return nil
	}
	return err
}

// lua脚本
const secKillScript = `
    -- Check if User has coupon
    -- KEYS[1]: hasCouponKey "{username}-has"
    -- KEYS[2]: couponName   "{couponName}"
    -- KEYS[3]: couponKey    "{couponName}-info"
    -- 返回值有-1, -2, -3, 都代表抢购失败
    -- 返回值为1代表抢购成功
    -- Check if coupon exists and is cached
	local couponLeft = redis.call("hget", KEYS[3], "stock");
	if (couponLeft == false)
	then
		return -2;  -- No such coupon
	end
	if (tonumber(couponLeft) == 0)  --- couponLeft是字符串类型
    then
		return -3;  --  No Coupon Left.
	end
    -- Check if the user has got the coupon --
	local userHasCoupon = redis.call("SISMEMBER", KEYS[1], KEYS[2]);
	if (userHasCoupon == 1)
	then
		return -1;
	end
    -- User gets the coupon --
	redis.call("hset", KEYS[3], "stock", couponLeft - 1);
	redis.call("SADD", KEYS[1], KEYS[2]);
	return 1;
`

// 预加载lua脚本
func (r *redisRepository) PreLoadScript(script string) string {
	// sha := sha1.Sum([]byte(script))
	scriptsExists, err := r.RClient.ScriptExists(script).Result()
	if err != nil {
		panic("Failed to check if script exists: " + err.Error())
	}
	if !scriptsExists[0] {
		scriptSHA, err := r.RClient.ScriptLoad(script).Result()
		if err != nil {
			panic("Failed to load script " + script + " err: " + err.Error())
		}
		return scriptSHA
	}
	print("Script Exists.")
	return ""
}
