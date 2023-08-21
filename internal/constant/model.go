package constant

import (
	"fmt"
)

type ServiceInfo struct {
	Host string `mapstructure:"host"`
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}

func (s ServiceInfo) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type EtcdInfo struct {
	Host string `mapstructure:"host"`
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}

type RedisInfo struct {
	Host     string `mapstructure:"host"`
	Name     string `mapstructure:"name"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

func (r RedisInfo) Addr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

type OSSInfo struct {
	Endpoint            string `mapstructure:"endpoint"`
	AccessKeyID         string `mapstructure:"access-key-id"`
	AccessKeySecret     string `mapstructure:"access-key-secret"`
	BucketName          string `mapstructure:"bucket-name"`
	CoverBaseURL        string `mapstructure:"cover-base-url"`
	VideoBaseURL        string `mapstructure:"video-base-url"`
	CoverResourceFolder string `mapstructure:"cover-resource-folder"`
	VideoResourceFolder string `mapstructure:"video-resource-folder"`
}

type DatasourceInfo struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Settings string `mapstructure:"settings"`
}

func (d DatasourceInfo) DSNString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Database,
		d.Settings)
}

type Constant struct {
	APIService      ServiceInfo    `mapstructure:"api-service"`
	UserService     ServiceInfo    `mapstructure:"user-service"`
	VideoService    ServiceInfo    `mapstructure:"video-service"`
	CommentService  ServiceInfo    `mapstructure:"comment-service"`
	FavoriteService ServiceInfo    `mapstructure:"favorite-service"`
	Etcd            EtcdInfo       `mapstructure:"etcd"`
	Redis           RedisInfo      `mapstructure:"redis"`
	OSS             OSSInfo        `mapstructure:"oss"`
	Datasource      DatasourceInfo `mapstructure:"datasource"`
}
