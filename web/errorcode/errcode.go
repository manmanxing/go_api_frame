package errorcode

//自定义错误返回信息
//code=1 表示成功，code=-1 表示失败。
//msg 表示提示信息。
//data 表示要返回的数据。
type ErrInfo struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetErr(code string) ErrInfo {
	error, ok := ErrMap[code]
	if ok {
		return ErrInfo{
			Code: code,
			Msg:  error,
			Data: nil,
		}

	} else {
		return ErrInfo{
			Code: "-1",
			Msg:  "code码异常",
			Data: nil,
		}
	}
}

var ErrMap = map[string]string{
	"0":      "",
	"100002": "参数错误",
}

var Params_err = "100002" //参数错误
