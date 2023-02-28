package util

import (
	"Raising/conf"
	"fmt"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func DeleteFile(filename string) error {
	var AccessKey = conf.AccessKey
	var SerectKey = conf.SerectKey
	var Bucket = conf.Bucket

	mac := qbox.NewMac(AccessKey, SerectKey)
	fmt.Println("mmaacc", mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Region=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)
	key := filename

	err := bucketManager.Delete(Bucket, key)
	if err != nil {
		panic(err)

	}
	return nil
}
func DeleteFiles(filename []string) error {
	var AccessKey = conf.AccessKey
	var SerectKey = conf.SerectKey
	var Bucket = conf.Bucket
	mac := qbox.NewMac(AccessKey, SerectKey)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Region=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)
	keys := filename
	deleteOps := make([]string, 0, len(keys))
	for _, key := range keys {
		deleteOps = append(deleteOps, storage.URIDelete(Bucket, key))
	}
	rets, err := bucketManager.Batch(deleteOps)
	if err != nil {
		if _, ok := err.(*storage.ErrorInfo); ok {
			for _, ret := range rets {
				if ret.Code != 200 {
					fmt.Println("ret.Data.Error:", ret.Data.Error)
					return fmt.Errorf(ret.Data.Error)
				}
			}
		} else {
			fmt.Println("err:", err)
			return err
		}

	}
	return nil

}
