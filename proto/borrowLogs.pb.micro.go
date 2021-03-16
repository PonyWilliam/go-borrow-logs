// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/borrowLogs.proto

//
//ID int64 `gorm:"primary_key;auto_increment;not_null" json:"id"`
//ReqWID int64 `json:"req_wid"`//转借者ID
//RspWID int64 `json:"rsp_wid"`//收到者ID
//Confirm bool `json:"confirm"`//对方是否已确认
//Reason string `json:"reason"`//原因说明
//Boss int8 `json:"boss"`//是否需要boss确认，1代表不需要，2代表需要且未确认，3代表需要且确认

package borrowLogs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for BorrowLogs service

func NewBorrowLogsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for BorrowLogs service

type BorrowLogsService interface {
	ToOther(ctx context.Context, in *ReqToOther, opts ...client.CallOption) (*Rsp_All, error)
	Confirm(ctx context.Context, in *Req_Confirm, opts ...client.CallOption) (*Rsp_All, error)
}

type borrowLogsService struct {
	c    client.Client
	name string
}

func NewBorrowLogsService(name string, c client.Client) BorrowLogsService {
	return &borrowLogsService{
		c:    c,
		name: name,
	}
}

func (c *borrowLogsService) ToOther(ctx context.Context, in *ReqToOther, opts ...client.CallOption) (*Rsp_All, error) {
	req := c.c.NewRequest(c.name, "BorrowLogs.to_other", in)
	out := new(Rsp_All)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borrowLogsService) Confirm(ctx context.Context, in *Req_Confirm, opts ...client.CallOption) (*Rsp_All, error) {
	req := c.c.NewRequest(c.name, "BorrowLogs.confirm", in)
	out := new(Rsp_All)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BorrowLogs service

type BorrowLogsHandler interface {
	ToOther(context.Context, *ReqToOther, *Rsp_All) error
	Confirm(context.Context, *Req_Confirm, *Rsp_All) error
}

func RegisterBorrowLogsHandler(s server.Server, hdlr BorrowLogsHandler, opts ...server.HandlerOption) error {
	type borrowLogs interface {
		ToOther(ctx context.Context, in *ReqToOther, out *Rsp_All) error
		Confirm(ctx context.Context, in *Req_Confirm, out *Rsp_All) error
	}
	type BorrowLogs struct {
		borrowLogs
	}
	h := &borrowLogsHandler{hdlr}
	return s.Handle(s.NewHandler(&BorrowLogs{h}, opts...))
}

type borrowLogsHandler struct {
	BorrowLogsHandler
}

func (h *borrowLogsHandler) ToOther(ctx context.Context, in *ReqToOther, out *Rsp_All) error {
	return h.BorrowLogsHandler.ToOther(ctx, in, out)
}

func (h *borrowLogsHandler) Confirm(ctx context.Context, in *Req_Confirm, out *Rsp_All) error {
	return h.BorrowLogsHandler.Confirm(ctx, in, out)
}
