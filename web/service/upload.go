package service

import (
	"github.com/gin-gonic/gin"
	"go_api_frame/web/common/errcode"
	"go_api_frame/web/common/upload"
	"go_api_frame/web/middleware/log"
)

func UploadImage(ctx *gin.Context) interface{} {
	file, image, err := ctx.Request.FormFile("image")
	if err != nil || image == nil {
		log.SugarLogger.Error("err:", err)
		panic(errcode.ParamsErr)
	}
	fileName := upload.GetImageName(image.Filename)
	savePath := upload.GetImagePath()
	//检测文件格式和大小
	if !upload.CheckImageExt(fileName) || !upload.CheckImageSize(file) {
		panic(errcode.ErrorUploadCheckImageFormat)
	}
	//检测路径是否存在和权限
	err = upload.CheckImage(savePath)
	if err != nil {
		log.SugarLogger.Error("err:", err)
		panic(errcode.ErrorUploadCheckImageFail)
	}
	//保存文件
	err = ctx.SaveUploadedFile(image, savePath+fileName)
	if err != nil {
		log.SugarLogger.Error("err:", err)
		panic(errcode.ErrorUploadSaveImageFail)
	}
	return upload.GetImageFullUrl(fileName)
}
