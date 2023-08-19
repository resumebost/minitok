package oss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	uuid "github.com/satori/go.uuid"
	"minitok/internal/conf"
)

var (
	client       *oss.Client
	bucket       *oss.Bucket
	videoBaseURL string
	coverBaseURL string
)

func InitOSS() {
	// 设置连接数为10，每个主机的最大闲置连接数为20，每个主机的最大连接数为20。
	//conn := oss.MaxConns(10,20,20)
	// 设置HTTP连接超时时间为20秒，HTTP读取或写入超时时间为60秒。
	//time := oss.Timeout(20,60)

	// 开启CRC加密。
	crc := oss.EnableCRC(true)
	// 设置日志模式。
	logLevel := oss.SetLogLevel(oss.LogOff)

	// 创建OSSClient实例
	client, err := oss.New(conf.Endpoint, conf.AccessKeyID, conf.AccessKeySecret, crc, logLevel)
	if err != nil {
		panic(err)
	}
	// 获取存储空间
	bucket, err = client.Bucket(conf.BucketName)
	if err != nil {
		panic(err)
	}

	videoBaseURL = conf.VideoBaseURL
	coverBaseURL = conf.CoverBaseURL
}

func UploadVideo(video []byte) (string, error) {
	reader := bytes.NewReader(video)
	objectName := videoBaseURL + uuid.NewV4().String() + ".mp4"
	err := bucket.PutObject(objectName, reader)
	if err != nil {
		return "", err
	}
	return objectName, nil
}

func UploadCover(cover []byte) (string, error) {
	reader := bytes.NewReader(cover)
	objectName := coverBaseURL + uuid.NewV4().String() + ".png"
	err := bucket.PutObject(objectName, reader)
	if err != nil {
		return "", err
	}
	return objectName, nil
}

func UploadLocalVideo(localVideoPath string) (string, error) {
	objectName := videoBaseURL + uuid.NewV4().String() + ".mp4"
	err := bucket.PutObjectFromFile(objectName, localVideoPath)
	if err != nil {
		return "", err
	}
	return objectName, nil
}

func UploadLocalCover(localCoverPath string) (string, error) {
	objectName := coverBaseURL + uuid.NewV4().String() + ".png"
	err := bucket.PutObjectFromFile(objectName, localCoverPath)
	if err != nil {
		return "", err
	}
	return objectName, nil
}
