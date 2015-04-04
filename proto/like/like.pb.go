// Code generated by protoc-gen-go.
// source: proto/like/like.proto
// DO NOT EDIT!

/*
Package like is a generated protocol buffer package.

It is generated from these files:
	proto/like/like.proto

It has these top-level messages:
	LikeRequest
	LikeResponse
	Like
*/
package like

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type LikeRequest struct {
	Trace  string `protobuf:"bytes,1,opt,name=trace" json:"trace,omitempty"`
	From   string `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	PostID int32  `protobuf:"varint,3,opt,name=postID" json:"postID,omitempty"`
	UserID int32  `protobuf:"varint,4,opt,name=userID" json:"userID,omitempty"`
}

func (m *LikeRequest) Reset()         { *m = LikeRequest{} }
func (m *LikeRequest) String() string { return proto.CompactTextString(m) }
func (*LikeRequest) ProtoMessage()    {}

type LikeResponse struct {
	Trace string `protobuf:"bytes,1,opt,name=trace" json:"trace,omitempty"`
	From  string `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	Like  *Like  `protobuf:"bytes,3,opt,name=like" json:"like,omitempty"`
}

func (m *LikeResponse) Reset()         { *m = LikeResponse{} }
func (m *LikeResponse) String() string { return proto.CompactTextString(m) }
func (*LikeResponse) ProtoMessage()    {}

func (m *LikeResponse) GetLike() *Like {
	if m != nil {
		return m.Like
	}
	return nil
}

type Like struct {
	PostID int32 `protobuf:"varint,1,opt,name=postID" json:"postID,omitempty"`
	Count  int32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *Like) Reset()         { *m = Like{} }
func (m *Like) String() string { return proto.CompactTextString(m) }
func (*Like) ProtoMessage()    {}

func init() {
}

// Client API for LikeService service

type LikeServiceClient interface {
	RecordLike(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error)
}

type likeServiceClient struct {
	cc *grpc.ClientConn
}

func NewLikeServiceClient(cc *grpc.ClientConn) LikeServiceClient {
	return &likeServiceClient{cc}
}

func (c *likeServiceClient) RecordLike(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error) {
	out := new(LikeResponse)
	err := grpc.Invoke(ctx, "/.LikeService/RecordLike", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LikeService service

type LikeServiceServer interface {
	RecordLike(context.Context, *LikeRequest) (*LikeResponse, error)
}

func RegisterLikeServiceServer(s *grpc.Server, srv LikeServiceServer) {
	s.RegisterService(&_LikeService_serviceDesc, srv)
}

func _LikeService_RecordLike_Handler(srv interface{}, ctx context.Context, buf []byte) (interface{}, error) {
	in := new(LikeRequest)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(LikeServiceServer).RecordLike(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _LikeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: ".LikeService",
	HandlerType: (*LikeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordLike",
			Handler:    _LikeService_RecordLike_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}