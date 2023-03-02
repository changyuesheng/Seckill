package utils

const (
	SUCCESS = 200
	ERROR   = 500

	// 用户相关
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NO_RIGHT  = 1008
	ERROR_USER_NOT_EXIST = 1003

	// token相关
	ERROR_TOKEN_EXIST      = 2001
	ERROR_TOKEN_RUNTIME    = 2002
	ERROR_TOKEN_WRONG      = 2003
	ERROR_TOKEN_TYPE_WRONG = 2004

	// 优惠券相关
)

var codemsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_USER_NO_RIGHT:    "用户没有权限",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
