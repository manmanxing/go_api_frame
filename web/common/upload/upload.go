package upload

import (
	"errors"
	"fmt"
	"go_api_frame/web/common/config"
	"go_api_frame/web/common/util"
	"go_api_frame/web/middleware/log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

func GetImageFullUrl(name string) string {
	return config.MyConfig.Host + ":" + strconv.Itoa(config.MyConfig.Port) + "/" + config.MyConfig.ImageSavePath + name
}

//获取文件名，MD5加密后
func GetImageName(name string) string {
	ext := GetExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

//检测文件后缀
func CheckImageExt(fileName string) bool {
	ext := GetExt(fileName)
	exts := strings.Split(config.MyConfig.ImageAllowExts, ",")
	for i := range exts {
		if strings.ToUpper(exts[i]) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

//检测文件大小
func CheckImageSize(f multipart.File) bool {
	size := GetSize(f)
	fmt.Println("文件 size：", size, " 要求大小为：", config.MyConfig.ImageMaxSize)
	return size <= (config.MyConfig.ImageMaxSize * 1024 * 1024)
}

//检测文件是否存在
//相对路径
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		log.SugarLogger.Error("err:", err)
		return err
	}
	err = IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		log.SugarLogger.Error("err:", err)
		return err
	}
	ok := CheckPermission(dir + "/" + src)
	if ok {
		return errors.New("file.CheckPermission Permission denied src: " + src)
	}
	return nil
}

//相对路径
func GetImagePath() string {
	return config.MyConfig.ImageSavePath
}
