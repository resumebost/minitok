# Minitok

青训营后端项目: 简易抖音 minitok.

## 项目依赖

- API 框架: gin 
- RPC 框架: kitex
- 持久层框架: gorm
- 服务注册/发现: etcd

## 开发指南

### 项目部署

数据库 sql 文件在 `config` 下.

TODO

### 规范

- [提交规范](doc/commit.md)
- [命名规范](doc/naming.md)

### 如何实现功能

以实现 video 功能为例. 有些文件夹可能没有创建, 可自行创建.

- 在 `idl/video/service.proto` 下定义 rpc 接口和相应的 message
- 在根目录运行 `./scripts/kitex-all.sh` 脚本, 生成所有框架文件
- 在 `cmd/video/handler.go` 下实现接口, 可以考虑将逻辑放到一个单独的 service 文件夹内
- 在 `cmd/video/dal/video.go` 下封装对数据层的操作
- 如果需要用到 redis 可以仿照函数 SetVideoDB 创立一个 SetVideoRD
- 在 `cmd/video/rpc/` 下建立需要用到的 rpc client:
  - 将 init 函数调用放到 `cmd/video/rpc/init.go` 里的 InitForVideo 中
  - 可以仿照 `cmd/api/rpc/` 下的 rpc client 建立方法

### 错误处理

API 调用和 RPC 调用的返回格式相似, 所以使用相似的错误格式.

错误的基本定义在 `internal/unierr/errcore.go` 内, 同文件夹下的 `errors.go` 是预定义的一些错误.

如果要增加错误的话, 在 `errors.go` 中添加.

### 如何测试 RPC 接口

以测试 video 为例子.

在根目录的 `test` 文件夹下建立 `video_test.go`, 复用 `cmd/api/rpc` 下建立好的 rpc client 进行测试.

例如测试 Feed 接口:

```go
package test

func TestFeed(t *testing.T) {
    doFeed(t)
}

func doFeed(t assert.TestingT) {
    resp, err := rpc.Feed(ctx, &feed.FeedRequest{
        LatestTime: nil,
        Token:      nil,
    })
    assert.NoError(t, err)
    fmt.Printf("%v", resp)
}
```

命令行的方式待调研, 可以看一下 grpcui.

### 日志

暂时用 klog.

### 其它

- 经过 JWT 中间件验证之后, ctx 会以键值对的形式保存用户的 username 和 id, key 就是 `username` 和 `id`.
