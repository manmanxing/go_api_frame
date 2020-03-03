package errcode

var ErrMap = map[string]string{
	"100001": "参数错误",
	"100002": "数据库错误",
	"100003": "验证token失败",
	"100004": "验证token超时",
}

var Params_err = "100001"                     //参数错误
var Database_err = "100002"                   //数据库错误
var ERROR_AUTH_CHECK_TOKEN_FAIL = "100003"    //验证token失败
var ERROR_AUTH_CHECK_TOKEN_TIMEOUT = "100004" //验证token超时
