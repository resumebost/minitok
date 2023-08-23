package oss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/kitex/pkg/klog"
	"minitok/internal/constant"
)

var (
	client       *oss.Client
	bucket       *oss.Bucket
	ossConstants = &constant.AllConstants.OSS
)

func InitOSS() {

	// 开启CRC加密
	crc := oss.EnableCRC(true)
	// 设置日志模式
	logLevel := oss.SetLogLevel(oss.LogOff)

	// 创建OSSClient实例
	client, err := oss.New(ossConstants.Endpoint,
		ossConstants.AccessKeyID,
		ossConstants.AccessKeySecret,

		crc,
		logLevel)
	if err != nil {
		klog.Fatalf("OSS instance creation failed: %v", err)
	}

	// 获取存储空间
	bucket, err = client.Bucket(ossConstants.BucketName)
	if err != nil {
		klog.Fatalf("Can not get bucket instance: %v", err)
	}

	// videoBaseURL = constant.VideoBaseURL
	// coverBaseURL = constant.CoverBaseURL
}

func UploadVideo(video []byte, objectName string) error {
	reader := bytes.NewReader(video)
	err := bucket.PutObject(objectName, reader)
	if err != nil {
		return err
	}
	return nil
}

func UploadCover(cover []byte, objectName string) error {
	reader := bytes.NewReader(cover)
	err := bucket.PutObject(objectName, reader)
	if err != nil {
		return err
	}
	return nil
}

func UploadLocalVideo(localVideoPath string, objectName string) error {
	err := bucket.PutObjectFromFile(objectName, localVideoPath)
	if err != nil {
		return err
	}
	return nil
}

func UploadLocalCover(localCoverPath string, objectName string) error {
	err := bucket.PutObjectFromFile(objectName, localCoverPath)
	if err != nil {
		return err
	}
	return nil
}
