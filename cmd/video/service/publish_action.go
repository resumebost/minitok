package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	uuid "github.com/satori/go.uuid"
	"minitok/cmd/video/dal"
	"minitok/internal/constant"
	"minitok/internal/jwt"
	"minitok/internal/oss"
	"minitok/internal/util"
	"minitok/kitex_gen/video"
	"os"
	"sync"
)

var ossInfo = &constant.AllConstants.OSS

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
	videoPath := ossInfo.VideoResourceFolder + videoName

	err := os.WriteFile(videoPath, req.Data, os.ModePerm)
	if err != nil {
		//return unierr.VideoPublishFiled
		return err
	}

	coverName, err := util.GetVideoCover(videoPath, 3)
	if err != nil {
		//return unierr.CoverGeneFiled
		return err
	}
	coverPath := ossInfo.CoverResourceFolder + coverName

	var wg sync.WaitGroup
	wg.Add(2)

	//上传视频到OSS
	go func() {
		defer func() {
			if err := recover(); err != nil { //防止协程崩溃，保持健壮性
				klog.Fatalf("upload video to oss filed: %s", err)
			}
		}()
		defer wg.Done()
		err = oss.UploadLocalVideo(videoPath, videoName)
		if err != nil {
			klog.Fatalf("upload video to oss filed: %s", err)
		}
	}()

	//上传视频封面到OSS
	go func() {
		defer func() {
			if err := recover(); err != nil {
				klog.Fatalf("upload cover to oss filed: %s", err)
			}
		}()
		defer wg.Done()

		err = oss.UploadLocalCover(coverPath, coverName)
		if err != nil {
			klog.Fatalf("upload cover to oss filed: %s", err)
		}
	}()
	wg.Wait()

	wg.Add(2)
	//删除本地暂存的视频
	go func() {
		defer func() {
			if err := recover(); err != nil {
				klog.Fatalf("本地视频删除失败: %s", err)
			}
		}()
		defer wg.Done()
		err := os.Remove(videoPath)
		if err != nil {
			klog.Fatalf("本地视频删除失败: %s", err)
		}
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil { //防止协程崩溃，保持健壮性
				klog.Fatalf("本地封面删除失败: %s", err)
			}
		}()
		defer wg.Done()
		err := os.Remove(coverPath)
		if err != nil {
			klog.Fatalf("本地封面删除失败: %s", err)
		}
	}()
	wg.Wait()

	//存储到数据库
	//id := s.ctx.Value("id")
	claims, err := jwt.ParseToken(req.Token)
	id := claims.ID

	if err != nil {
		return err
	}

	UPLPrefix := "https://" + ossInfo.BucketName + "." + ossInfo.Endpoint + "/" + ossInfo.VideoBaseURL

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
