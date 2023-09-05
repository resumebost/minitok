package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"minitok/cmd/api/rpc"
	"minitok/internal/jwt"
	"minitok/kitex_gen/video"
	"os"
	"testing"
	"time"
)

var TOKEN, _ = jwt.GenToken(2, "xxhhy")

func TestActionPublish(t *testing.T) {
	doActionPublish(t)
}

func TestFeed(t *testing.T) {
	doActionFeedWithTime(t)
}

func TestPublishList(t *testing.T) {
	doActionPublishList(t)
}

func TestPublishIdList(t *testing.T) {
	doActionPublishIdList(t)
}

func TestVideoList(t *testing.T) {
	doActionVideoList(t)
}

func BenchmarkActionPublish(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doActionPublish(b)
		}
	})
}

// func InitTest()
func doActionPublish(t assert.TestingT) {
	filePath := "../webresource/video/testvideo.mp4"
	bytes, err := os.ReadFile(filePath)
	assert.NoError(t, err)

	resp, err := rpc.VideoPublishAction(ctx, &video.PublishActionRequest{
		Token: TOKEN,
		Data:  bytes,
		Title: "testVideo",
	})
	assert.NoError(t, err)
	fmt.Println(resp)
}

func doActionFeedWithTime(t *testing.T) {
	resp, err := rpc.Feed(ctx, &video.FeedRequest{
		LatestTime: time.Now().Unix(),
		Token:      TOKEN})
	assert.NoError(t, err)
	fmt.Println(resp)
}

func doActionPublishList(t *testing.T) {
	resp, err := rpc.VideoPublishList(ctx, &video.PublishListRequest{
		UserId: 1,
		Token:  TOKEN,
	})
	assert.NoError(t, err)
	fmt.Println(resp)
}

func doActionPublishIdList(t *testing.T) {
	resp, err := rpc.VideoIdList(ctx, &video.PublishListIdsRequest{UserId: 1})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp)
}

func doActionVideoList(t *testing.T) {
	videoIdList := []int64{1, 2, 3, 4, 5}
	fmt.Println(len(videoIdList))
	fmt.Println(videoIdList)
	resp, err := rpc.VideoList(ctx, &video.GetVideosRequest{Token: TOKEN, VideoIds: videoIdList})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp)
}
