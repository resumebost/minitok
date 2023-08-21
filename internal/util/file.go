package util

import (
	"bytes"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/disintegration/imaging"
	uuid "github.com/satori/go.uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"mime/multipart"
	"minitok/internal/constant"
	"os"
)

func File2Bytes(file *multipart.FileHeader) ([]byte, error) {
	content, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer content.Close()

	bytes, err := io.ReadAll(content)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

var ossConstants = &constant.AllConstants.OSS

func GetVideoCover(videoPath string, frameNum int64) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		klog.Fatal(err)
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		klog.Fatal(err)
	}

	coverName := uuid.NewV4().String() + ".png"

	err = imaging.Save(img, ossConstants.CoverResourceFolder+coverName)
	if err != nil {
		klog.Fatal(err)
	}
	return coverName, err
}
