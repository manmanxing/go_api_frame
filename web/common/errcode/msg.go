package errcode

var ErrMap = map[string]string{
	"100001": "参数错误",
	"100002": "数据库错误",
	"100003": "验证token失败",
	"100004": "验证token超时",
	"100005": "保存图片失败",
	"100006": "检查图片失败",
	"100007": "校验图片格式或大小问题",
	"100008": "获取缓存失败",
	"100009": "Json转化失败",
	"100010": "设置缓存失败",
}

const (
	ParamsErr                   = "100001" //参数错误
	DatabaseErr                 = "100002" //数据库错误
	ErrorAuthCheckTokenFail     = "100003" //验证token失败
	ErrorAuthCheckTokenTimeout  = "100004" //验证token超时
	ErrorUploadSaveImageFail    = "100005" //保存图片失败
	ErrorUploadCheckImageFail   = "100006" //检查图片失败
	ErrorUploadCheckImageFormat = "100007" //校验图片格式或大小问题
	GetCacheFail                = "100008" //获取缓存失败
	JsonUnmarshalFail           = "100009" //Json转化失败
	SetCacheFail                = "100009" //设置缓存失败
)
