package upload

import (
	"fmt"
	//"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// 获取文件大小的接口
type Size interface {
	Size() int64
}

// 获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}

func GetSize(f multipart.File) int64 {
	//con, err := ioutil.ReadFile(f)
	//return len(con), err
	if staInterface, ok := f.(Stat); ok {
		fileInfo, _ := staInterface.Stat()
		fmt.Println("文件信息 size:", fileInfo.Size())
		return fileInfo.Size()
	}
	if sizeInterFace, ok := f.(Size); ok {
		fmt.Println("文件大小 size:", sizeInterFace.Size())
		return sizeInterFace.Size()
	}
	return 0
}

func GetExt(fileName string) string {
	return path.Ext(fileName)
}

func CheckExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

func IsNotExistMkDir(src string) error {
	notExist := CheckExist(src)
	if notExist {
		err := MkDir(src)
		if err != nil {
			return err
		}
	}
	return nil
}

func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
