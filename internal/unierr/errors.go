package unierr

var (
	SuccessCode       = NewErrCore(0, "ok")
	InternalError     = NewErrCore(1, "Internal error")
	InvalidTokenError = NewErrCore(2, "Invalid JWT token")
	NoTokenError      = NewErrCore(3, "Missing JWT token")
	IllegalParams     = NewErrCore(4, "非法参数")
)

var (
	VideoReadError      = NewErrCore(20001, "视频数据读取失败")
	VideoConvertError   = NewErrCore(20002, "视频数据转换失败")
	VideoIsEmptyError   = NewErrCore(20003, "视频数据为空")
	VideoPublishFiled   = NewErrCore(20004, "视频上传失败")
	CoverGeneFiled      = NewErrCore(20005, "封面生成失败")
	CoverUploadFiled    = NewErrCore(20006, "封面上传失败")
	GetVideoListFiled   = NewErrCore(20007, "视频列表获取失败")
	GetFeedFiled        = NewErrCore(20008, "视频流获取失败")
	GetPublishListFiled = NewErrCore(20009, "发布视频获取失败")
	VideoNotFound       = NewErrCore(20010, "视频找不到")
)

var (
	UsernameOrPassword                    = NewErrCore(30001, "用户名或密码为空")
	UsernameOrPasswordLenMore32Characters = NewErrCore(30002, "用户名或密码长度不能大于32个字符")
	UsernameOrPasswordFailToUpload        = NewErrCore(30003, "用户名或密码上传失败")
	UsernameExist                         = NewErrCore(30004, "该用户名已存在")
	UsernameNotExist                      = NewErrCore(30005, "用户名不存在")
	PasswordWrong                         = NewErrCore(30006, "密码错误")
	UserNotExist                          = NewErrCore(30007, "用户不存在")
	TokenNotExist                         = NewErrCore(30008, "Token不存在")
	UserIdInvalid                         = NewErrCore(30009, "user_id不合法")
)

var (
	FavoriteActionError = NewErrCore(40001, "点赞或者取消点赞失败")
)

var (
	CommentNotFound    = NewErrCore(50001, "评论查询失败")
	CommentPostFiled   = NewErrCore(50002, "评论发表失败")
	CommentDeleteFiled = NewErrCore(50003, "评论删除失败")
	CommentEmpty       = NewErrCore(50004, "评论内容为空")
)
