package util

import (
	"Raising/conf"
	"Raising/pkg/e"
	"context"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> fd910d7 (golang)
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// 封装上传图片到七牛云然后返回状态和图片的url
func UploadToQiNiu(file multipart.File, fileheader *multipart.FileHeader, fileSize int64, pid string) (int, string) {
	var AccessKey = conf.AccessKey
	var SerectKey = conf.SerectKey
	var Bucket = conf.Bucket
	var ImgUrl = conf.QiniuServer
<<<<<<< HEAD
	fmt.Println(AccessKey)
	fmt.Println(SerectKey)
=======

>>>>>>> fd910d7 (golang)
	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SerectKey)
<<<<<<< HEAD
	fmt.Println("mac:", mac)
	upToken := putPlicy.UploadToken(mac)
	fmt.Println("uptoken:", upToken)
=======
	upToken := putPlicy.UploadToken(mac)
>>>>>>> fd910d7 (golang)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	var filebox string
	fileheader.Filename = pid + fileheader.Filename
	key := filebox + fileheader.Filename
	err := formUploader.Put(context.Background(), &ret, upToken, key, file, fileSize, &putExtra)
	if err != nil {
		code := e.ErrorUploadFile
		return code, err.Error()
	}

	url := ImgUrl + ret.Key
	return 200, url
}
