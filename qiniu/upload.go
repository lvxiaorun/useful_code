package main

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"fmt"
	"context"
)

func main(){
	accessKey := "OfWDisyHuTZ3DRf1tjIN57WIFKTBG58_06i4FmMN"
	secretKey := "6iOZfhDsKOtZ3Of-pTKrHr2_eSJTOON9OJAJgfdf"
	mac := qbox.NewMac(accessKey, secretKey)
	bucket:="lvxiaorun"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	localFile := "D:\\BugReport.txt"
	key := "test"
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key,ret.Hash)
}
