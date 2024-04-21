// 全局常量
package consts

// 未来扩充可用 ERR_*** 的形式命名
const (
	SUCCESS              = 0      //成功
	FAIL                 = -1     //失败
	PARAM_ERROR          = -2     //参数错误
	UNKNOWN_ERROR_STATUS = -1211  //未知错误,一般都是数据库死锁
	IN_MAINTENANCE       = -10000 //系统维护中
	NOT_IN_SAFE_IP       = -10100 //IP不在白名单内,无法操作
	NO_TOKEN             = -4001  //缺少token
	TOKEN_AUTH_FAIL      = -4002  //无此用户, 未登录或登录超时
	TOKEN_EXPIRED        = -4003  //存取token 过期
	TOKEN_INVALID        = -4005  //token 无效
)

var errMsg = map[int]string{
	SUCCESS:              "成功",
	FAIL:                 "失败",
	PARAM_ERROR:          "参数错误",
	UNKNOWN_ERROR_STATUS: "未知错误",
	IN_MAINTENANCE:       "系统维护中",
	NOT_IN_SAFE_IP:       "IP不在白名单内,无法操作",
	NO_TOKEN:             "缺少token",
	TOKEN_AUTH_FAIL:      "无此用户, 未登录或登录超时",
	TOKEN_EXPIRED:        "token过期",
	TOKEN_INVALID:        "token无效",
}

// 返回错误信息
func RetErrMsg(code int) string {
	str, ok := errMsg[code]
	if ok {
		return str
	}
	return RetErrMsg(UNKNOWN_ERROR_STATUS)
}
