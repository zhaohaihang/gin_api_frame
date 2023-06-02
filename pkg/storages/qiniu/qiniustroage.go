package qiniu

import (
	"context"
	"gin_api_frame/config"
	"mime/multipart"

	"github.com/google/wire"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiNiuStroage struct {
	config *config.Config
}

func NewQiNiuStroage(c *config.Config) *QiNiuStroage {
	return &QiNiuStroage{
		config: c,
	}
}

var QiNiuStroageProviderSet = wire.NewSet(NewQiNiuStroage)

func (q *QiNiuStroage) UploadToQiNiu(key string ,file multipart.File, fileSize int64) (path string, err error) {
	var AccessKey = q.config.QiNiu.AccessKey
	var SerectKey = q.config.QiNiu.SerectKey
	var Bucket = q.config.QiNiu.Bucket
	var ImgUrl = q.config.QiNiu.Domain

	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}

	mac := qbox.NewMac(AccessKey, SerectKey)
	upToken := putPlicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadongZheJiang2,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.Put(context.Background(), &ret, upToken, key,file, fileSize, &putExtra)

	if err != nil {
		return "", err
	}
	url := ImgUrl + ret.Key
	return url, nil
}
