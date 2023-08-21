package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	uuid "github.com/satori/go.uuid"
	"minitok/cmd/video/dal"
	"minitok/internal/conf"
	"minitok/internal/jwt"
	"minitok/internal/oss"
	"minitok/internal/unierr"
	"minitok/internal/util"
	"minitok/kitex_gen/video"
	"os"
)

type PublishActionService struct {
	ctx context.Context
}

func NewUploadVideoService(ctx context.Context) *PublishActionService {
	return &PublishActionService{
		ctx: ctx,
	}
}

func (s *PublishActionService) PublishVideo(req *video.PublishActionRequest) error {
	//存储视频信息到本地，截取封面
	videoName := uuid.NewV4().String() + ".mp4"
	videoPath := conf.VideoResourceFolder + videoName
	err := os.WriteFile(videoPath, req.Data, os.ModePerm)
	if err != nil {
		return unierr.VideoPublishFiled
	}

	coverName, err := util.GetVideoCover(videoPath, 3)
	if err != nil {
		return unierr.CoverGeneFiled
	}

	//上传视频到OSS
	err = oss.UploadLocalVideo(videoPath, videoName)
	if err != nil {
		return unierr.VideoPublishFiled
	}

	//上传视频封面到OSS
	coverPath := conf.CoverResourceFolder + coverName
	err = oss.UploadLocalCover(coverPath, coverName)
	if err != nil {
		return unierr.CoverUploadFiled
	}

	go func() {
		err := os.Remove(videoPath)
		if err != nil {
			klog.Fatal("本地视频删除失败", err)
		}
	}()
	go func() {
		err := os.Remove(coverPath)
		if err != nil {
			klog.Fatal("本地封面删除失败", err)
		}
	}()
	//存储到数据库
	//id := s.ctx.Value("id")
	claims, err := jwt.ParseToken(req.Token)
	id := claims.ID

	if err != nil {
		return err
	}

	UPLPrefix := "https://" + conf.BucketName + "." + conf.Endpoint + "/" + conf.VideoBaseURL

	_, err = dal.CreateVideo(&dal.Video{
		AuthorID: id,
		PlayURL:  UPLPrefix + videoName,
		CoverURL: UPLPrefix + coverName,
		Title:    req.Title,
	}, s.ctx)
	if err != nil {
		return err
	}
	return nil
}
