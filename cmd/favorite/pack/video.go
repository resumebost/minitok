package pack

//命名有冲突
import (
	v "minitok/kitex_gen/video"
)

func ConvertToFavoriteVideo(video *v.Video) *v.Video {
	if video == nil {
		return nil
	}

	return &v.Video{
		Id:            video.Id,
		Author:        video.Author,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    video.IsFavorite,
		Title:         video.Title,
	}
}

// 解析并封装rpc返回的结果
func ConvertToFavoriteVideos(videoMap map[int64]*v.Video, likedVideoIDs []int64) []*v.Video {
	likedVideos := make([]*v.Video, 0)
	for _, videoID := range likedVideoIDs {
		if video, ok := videoMap[videoID]; ok {
			likedVideos = append(likedVideos, ConvertToFavoriteVideo(video))
		}
	}
	return likedVideos
}
