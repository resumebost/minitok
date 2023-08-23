package unierr

var (
	SuccessCode       = NewErrCore(0, "ok")
	InternalError     = NewErrCore(1, "Internal error")
	InvalidTokenError = NewErrCore(2, "Invalid JWT token")
	NoTokenError      = NewErrCore(3, "Missing JWT token")
	IllegalParams     = NewErrCore(4, "非法参数")
)

var (
	VideoConvertError = NewErrCore(20001, "视频数据转换失败")
	VideoPublishFiled = NewErrCore(20002, "视频上传失败")

	GetVideoListFiled   = NewErrCore(20003, "视频列表获取失败")
	GetFeedFiled        = NewErrCore(20004, "视频流获取失败")
	GetPublishListFiled = NewErrCore(20005, "发布视频获取失败")
	VideoNotFound       = NewErrCore(20006, "未找到视频")
)
