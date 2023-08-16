// Code generated by Kitex v0.6.2. DO NOT EDIT.

package favoriteservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	favorite "minitok/kitex_gen/favorite"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

var favoriteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*favorite.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Action": kitex.NewMethodInfo(actionHandler, newActionArgs, newActionResult, false),
		"List":   kitex.NewMethodInfo(listHandler, newListArgs, newListResult, false),
		"Judge":  kitex.NewMethodInfo(judgeHandler, newJudgeArgs, newJudgeResult, false),
		"Count":  kitex.NewMethodInfo(countHandler, newCountArgs, newCountResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "favorite",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func actionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.ActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).Action(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ActionArgs:
		success, err := handler.(favorite.FavoriteService).Action(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ActionResult)
		realResult.Success = success
	}
	return nil
}
func newActionArgs() interface{} {
	return &ActionArgs{}
}

func newActionResult() interface{} {
	return &ActionResult{}
}

type ActionArgs struct {
	Req *favorite.ActionRequest
}

func (p *ActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.ActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ActionArgs) Unmarshal(in []byte) error {
	msg := new(favorite.ActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ActionArgs_Req_DEFAULT *favorite.ActionRequest

func (p *ActionArgs) GetReq() *favorite.ActionRequest {
	if !p.IsSetReq() {
		return ActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ActionArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ActionArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ActionResult struct {
	Success *favorite.ActionResponse
}

var ActionResult_Success_DEFAULT *favorite.ActionResponse

func (p *ActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.ActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ActionResult) Unmarshal(in []byte) error {
	msg := new(favorite.ActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ActionResult) GetSuccess() *favorite.ActionResponse {
	if !p.IsSetSuccess() {
		return ActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.ActionResponse)
}

func (p *ActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ActionResult) GetResult() interface{} {
	return p.Success
}

func listHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.ListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).List(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ListArgs:
		success, err := handler.(favorite.FavoriteService).List(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListResult)
		realResult.Success = success
	}
	return nil
}
func newListArgs() interface{} {
	return &ListArgs{}
}

func newListResult() interface{} {
	return &ListResult{}
}

type ListArgs struct {
	Req *favorite.ListRequest
}

func (p *ListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.ListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ListArgs) Unmarshal(in []byte) error {
	msg := new(favorite.ListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListArgs_Req_DEFAULT *favorite.ListRequest

func (p *ListArgs) GetReq() *favorite.ListRequest {
	if !p.IsSetReq() {
		return ListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListResult struct {
	Success *favorite.ListResponse
}

var ListResult_Success_DEFAULT *favorite.ListResponse

func (p *ListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.ListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ListResult) Unmarshal(in []byte) error {
	msg := new(favorite.ListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListResult) GetSuccess() *favorite.ListResponse {
	if !p.IsSetSuccess() {
		return ListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.ListResponse)
}

func (p *ListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListResult) GetResult() interface{} {
	return p.Success
}

func judgeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.JudgeRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).Judge(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *JudgeArgs:
		success, err := handler.(favorite.FavoriteService).Judge(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*JudgeResult)
		realResult.Success = success
	}
	return nil
}
func newJudgeArgs() interface{} {
	return &JudgeArgs{}
}

func newJudgeResult() interface{} {
	return &JudgeResult{}
}

type JudgeArgs struct {
	Req *favorite.JudgeRequest
}

func (p *JudgeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.JudgeRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *JudgeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *JudgeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *JudgeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in JudgeArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *JudgeArgs) Unmarshal(in []byte) error {
	msg := new(favorite.JudgeRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var JudgeArgs_Req_DEFAULT *favorite.JudgeRequest

func (p *JudgeArgs) GetReq() *favorite.JudgeRequest {
	if !p.IsSetReq() {
		return JudgeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *JudgeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *JudgeArgs) GetFirstArgument() interface{} {
	return p.Req
}

type JudgeResult struct {
	Success *favorite.JudgeResponse
}

var JudgeResult_Success_DEFAULT *favorite.JudgeResponse

func (p *JudgeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.JudgeResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *JudgeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *JudgeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *JudgeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in JudgeResult")
	}
	return proto.Marshal(p.Success)
}

func (p *JudgeResult) Unmarshal(in []byte) error {
	msg := new(favorite.JudgeResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *JudgeResult) GetSuccess() *favorite.JudgeResponse {
	if !p.IsSetSuccess() {
		return JudgeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *JudgeResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.JudgeResponse)
}

func (p *JudgeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *JudgeResult) GetResult() interface{} {
	return p.Success
}

func countHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.CountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).Count(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CountArgs:
		success, err := handler.(favorite.FavoriteService).Count(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CountResult)
		realResult.Success = success
	}
	return nil
}
func newCountArgs() interface{} {
	return &CountArgs{}
}

func newCountResult() interface{} {
	return &CountResult{}
}

type CountArgs struct {
	Req *favorite.CountRequest
}

func (p *CountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.CountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CountArgs) Unmarshal(in []byte) error {
	msg := new(favorite.CountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CountArgs_Req_DEFAULT *favorite.CountRequest

func (p *CountArgs) GetReq() *favorite.CountRequest {
	if !p.IsSetReq() {
		return CountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CountResult struct {
	Success *favorite.CountResponse
}

var CountResult_Success_DEFAULT *favorite.CountResponse

func (p *CountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.CountResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CountResult) Unmarshal(in []byte) error {
	msg := new(favorite.CountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CountResult) GetSuccess() *favorite.CountResponse {
	if !p.IsSetSuccess() {
		return CountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CountResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.CountResponse)
}

func (p *CountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CountResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Action(ctx context.Context, Req *favorite.ActionRequest) (r *favorite.ActionResponse, err error) {
	var _args ActionArgs
	_args.Req = Req
	var _result ActionResult
	if err = p.c.Call(ctx, "Action", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) List(ctx context.Context, Req *favorite.ListRequest) (r *favorite.ListResponse, err error) {
	var _args ListArgs
	_args.Req = Req
	var _result ListResult
	if err = p.c.Call(ctx, "List", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Judge(ctx context.Context, Req *favorite.JudgeRequest) (r *favorite.JudgeResponse, err error) {
	var _args JudgeArgs
	_args.Req = Req
	var _result JudgeResult
	if err = p.c.Call(ctx, "Judge", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Count(ctx context.Context, Req *favorite.CountRequest) (r *favorite.CountResponse, err error) {
	var _args CountArgs
	_args.Req = Req
	var _result CountResult
	if err = p.c.Call(ctx, "Count", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
