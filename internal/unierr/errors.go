package unierr

var (
	SuccessCode       = NewErrCore(0, "ok")
	InternalError     = NewErrCore(1, "Internal error")
	InvalidTokenError = NewErrCore(2, "Invalid JWT token")
	NoTokenError      = NewErrCore(3, "Missing JWT token")
	IllegalParams     = NewErrCore(4, "非法参数")
)

var (
	VideoReadError    = NewErrCore(20001, "视频数据读取失败")
	VideoConvertError = NewErrCore(20002, "视频数据转换失败")
	VideoIsEmptyError = NewErrCore(20003, "视频数据为空")
	VideoPublishFiled = NewErrCore(20004, "视频上传失败")
	CoverGeneFiled    = NewErrCore(20005, "封面生成失败")
	CoverUploadFiled  = NewErrCore(20006, "封面上传失败")

	GetVideoListFiled   = NewErrCore(20007, "视频列表获取失败")
	GetFeedFiled        = NewErrCore(20008, "视频流获取失败")
	GetPublishListFiled = NewErrCore(20009, "发布视频获取失败")
)
