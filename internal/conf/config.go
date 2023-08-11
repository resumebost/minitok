package conf

// TODO: 用 etcd + viper 做远程配置管理

const (
	UserServiceName     = "user_service"
	VideoServiceName    = "video_service"
	CommentServiceName  = "comment_service"
	FavoriteServiceName = "favorite_service"

	ETCDAddress = "127.0.0.1:2379"

	APIServiceAddress      = "127.0.0.1:8080"
	UserServiceAddress     = "127.0.0.1:8000"
	VideoServiceAddress    = "127.0.0.1:8001"
	CommentServiceAddress  = "127.0.0.1:8002"
	FavoriteServiceAddress = "127.0.0.1:8003"

	MySQLDefaultDSN = "root:123456@tcp(localhost:3306)/minitok?charset=utf8&parseTime=True&loc=Local"
	RedisAddress    = "127.0.0.1:6379"
	RedisPassword   = "123456"
)
