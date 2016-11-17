// Code generated by protoc-gen-go.
// source: codelingo.proto
// DO NOT EDIT!

/*
Package codelingo is a generated protocol buffer package.

It is generated from these files:
	codelingo.proto

It has these top-level messages:
	ListFactsRequest
	FactList
	Children
	ListLexiconsRequest
	ListLexiconsReply
	SessionRequest
	SessionReply
	QueryRequest
	QueryReply
	ReviewRequest
	ReviewReply
	Issue
	IssueRange
	Position
*/
package codelingo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ListFactsRequest struct {
	Lexicon string `protobuf:"bytes,1,opt,name=lexicon" json:"lexicon,omitempty"`
}

func (m *ListFactsRequest) Reset()                    { *m = ListFactsRequest{} }
func (m *ListFactsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListFactsRequest) ProtoMessage()               {}
func (*ListFactsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FactList struct {
	Facts map[string]*Children `protobuf:"bytes,1,rep,name=facts" json:"facts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FactList) Reset()                    { *m = FactList{} }
func (m *FactList) String() string            { return proto.CompactTextString(m) }
func (*FactList) ProtoMessage()               {}
func (*FactList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FactList) GetFacts() map[string]*Children {
	if m != nil {
		return m.Facts
	}
	return nil
}

type Children struct {
	Child []string `protobuf:"bytes,1,rep,name=child" json:"child,omitempty"`
}

func (m *Children) Reset()                    { *m = Children{} }
func (m *Children) String() string            { return proto.CompactTextString(m) }
func (*Children) ProtoMessage()               {}
func (*Children) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ListLexiconsRequest struct {
}

func (m *ListLexiconsRequest) Reset()                    { *m = ListLexiconsRequest{} }
func (m *ListLexiconsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLexiconsRequest) ProtoMessage()               {}
func (*ListLexiconsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ListLexiconsReply struct {
	Lexicons []string `protobuf:"bytes,1,rep,name=lexicons" json:"lexicons,omitempty"`
}

func (m *ListLexiconsReply) Reset()                    { *m = ListLexiconsReply{} }
func (m *ListLexiconsReply) String() string            { return proto.CompactTextString(m) }
func (*ListLexiconsReply) ProtoMessage()               {}
func (*ListLexiconsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type SessionRequest struct {
}

func (m *SessionRequest) Reset()                    { *m = SessionRequest{} }
func (m *SessionRequest) String() string            { return proto.CompactTextString(m) }
func (*SessionRequest) ProtoMessage()               {}
func (*SessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type SessionReply struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *SessionReply) Reset()                    { *m = SessionReply{} }
func (m *SessionReply) String() string            { return proto.CompactTextString(m) }
func (*SessionReply) ProtoMessage()               {}
func (*SessionReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// The request message containing the CLQL source code.
type QueryRequest struct {
	Clql string `protobuf:"bytes,1,opt,name=clql" json:"clql,omitempty"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// The query response.
type QueryReply struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *QueryReply) Reset()                    { *m = QueryReply{} }
func (m *QueryReply) String() string            { return proto.CompactTextString(m) }
func (*QueryReply) ProtoMessage()               {}
func (*QueryReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// The request message containing the files or directories to review.
type ReviewRequest struct {
	Key           string   `protobuf:"bytes,12,opt,name=key" json:"key,omitempty"`
	Host          string   `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	Owner         string   `protobuf:"bytes,2,opt,name=owner" json:"owner,omitempty"`
	Repo          string   `protobuf:"bytes,3,opt,name=repo" json:"repo,omitempty"`
	Sha           string   `protobuf:"bytes,4,opt,name=sha" json:"sha,omitempty"`
	FilesAndDirs  []string `protobuf:"bytes,5,rep,name=filesAndDirs" json:"filesAndDirs,omitempty"`
	Recursive     bool     `protobuf:"varint,6,opt,name=recursive" json:"recursive,omitempty"`
	Patches       []string `protobuf:"bytes,7,rep,name=Patches" json:"Patches,omitempty"`
	IsPullRequest bool     `protobuf:"varint,8,opt,name=isPullRequest" json:"isPullRequest,omitempty"`
	PullRequestID int64    `protobuf:"varint,9,opt,name=pullRequestID" json:"pullRequestID,omitempty"`
	Vcs           string   `protobuf:"bytes,10,opt,name=vcs" json:"vcs,omitempty"`
	Dotlingo      string   `protobuf:"bytes,11,opt,name=dotlingo" json:"dotlingo,omitempty"`
}

func (m *ReviewRequest) Reset()                    { *m = ReviewRequest{} }
func (m *ReviewRequest) String() string            { return proto.CompactTextString(m) }
func (*ReviewRequest) ProtoMessage()               {}
func (*ReviewRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// Empty, we use the control queue once a review is started
type ReviewReply struct {
}

func (m *ReviewReply) Reset()                    { *m = ReviewReply{} }
func (m *ReviewReply) String() string            { return proto.CompactTextString(m) }
func (*ReviewReply) ProtoMessage()               {}
func (*ReviewReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// Issue returned from a review.
type Issue struct {
	// The name of the issue.
	Name          string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Position      *IssueRange       `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
	Comment       string            `protobuf:"bytes,3,opt,name=comment" json:"comment,omitempty"`
	CtxBefore     string            `protobuf:"bytes,4,opt,name=ctxBefore" json:"ctxBefore,omitempty"`
	LineText      string            `protobuf:"bytes,5,opt,name=lineText" json:"lineText,omitempty"`
	CtxAfter      string            `protobuf:"bytes,6,opt,name=ctxAfter" json:"ctxAfter,omitempty"`
	Metrics       map[string]string `protobuf:"bytes,7,rep,name=metrics" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tags          []string          `protobuf:"bytes,8,rep,name=tags" json:"tags,omitempty"`
	Link          string            `protobuf:"bytes,9,opt,name=link" json:"link,omitempty"`
	NewCode       bool              `protobuf:"varint,10,opt,name=newCode" json:"newCode,omitempty"`
	Patch         string            `protobuf:"bytes,11,opt,name=patch" json:"patch,omitempty"`
	Err           string            `protobuf:"bytes,12,opt,name=err" json:"err,omitempty"`
	Discard       bool              `protobuf:"varint,13,opt,name=discard" json:"discard,omitempty"`
	DiscardReason string            `protobuf:"bytes,14,opt,name=discardReason" json:"discardReason,omitempty"`
}

func (m *Issue) Reset()                    { *m = Issue{} }
func (m *Issue) String() string            { return proto.CompactTextString(m) }
func (*Issue) ProtoMessage()               {}
func (*Issue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Issue) GetPosition() *IssueRange {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Issue) GetMetrics() map[string]string {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type IssueRange struct {
	Start *Position `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End   *Position `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
}

func (m *IssueRange) Reset()                    { *m = IssueRange{} }
func (m *IssueRange) String() string            { return proto.CompactTextString(m) }
func (*IssueRange) ProtoMessage()               {}
func (*IssueRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *IssueRange) GetStart() *Position {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *IssueRange) GetEnd() *Position {
	if m != nil {
		return m.End
	}
	return nil
}

type Position struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Offset   int64  `protobuf:"varint,2,opt,name=Offset" json:"Offset,omitempty"`
	Line     int64  `protobuf:"varint,3,opt,name=Line" json:"Line,omitempty"`
	Column   int64  `protobuf:"varint,4,opt,name=Column" json:"Column,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*ListFactsRequest)(nil), "codelingo.ListFactsRequest")
	proto.RegisterType((*FactList)(nil), "codelingo.FactList")
	proto.RegisterType((*Children)(nil), "codelingo.Children")
	proto.RegisterType((*ListLexiconsRequest)(nil), "codelingo.ListLexiconsRequest")
	proto.RegisterType((*ListLexiconsReply)(nil), "codelingo.ListLexiconsReply")
	proto.RegisterType((*SessionRequest)(nil), "codelingo.SessionRequest")
	proto.RegisterType((*SessionReply)(nil), "codelingo.SessionReply")
	proto.RegisterType((*QueryRequest)(nil), "codelingo.QueryRequest")
	proto.RegisterType((*QueryReply)(nil), "codelingo.QueryReply")
	proto.RegisterType((*ReviewRequest)(nil), "codelingo.ReviewRequest")
	proto.RegisterType((*ReviewReply)(nil), "codelingo.ReviewReply")
	proto.RegisterType((*Issue)(nil), "codelingo.Issue")
	proto.RegisterType((*IssueRange)(nil), "codelingo.IssueRange")
	proto.RegisterType((*Position)(nil), "codelingo.Position")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CodeLingo service

type CodeLingoClient interface {
	// Initialise session
	Session(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionReply, error)
	// Reviews files
	Review(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewReply, error)
	// Queries source code
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error)
	// Lists available lexicons
	ListLexicons(ctx context.Context, in *ListLexiconsRequest, opts ...grpc.CallOption) (*ListLexiconsReply, error)
	ListFacts(ctx context.Context, in *ListFactsRequest, opts ...grpc.CallOption) (*FactList, error)
}

type codeLingoClient struct {
	cc *grpc.ClientConn
}

func NewCodeLingoClient(cc *grpc.ClientConn) CodeLingoClient {
	return &codeLingoClient{cc}
}

func (c *codeLingoClient) Session(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionReply, error) {
	out := new(SessionReply)
	err := grpc.Invoke(ctx, "/codelingo.CodeLingo/Session", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeLingoClient) Review(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewReply, error) {
	out := new(ReviewReply)
	err := grpc.Invoke(ctx, "/codelingo.CodeLingo/Review", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeLingoClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error) {
	out := new(QueryReply)
	err := grpc.Invoke(ctx, "/codelingo.CodeLingo/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeLingoClient) ListLexicons(ctx context.Context, in *ListLexiconsRequest, opts ...grpc.CallOption) (*ListLexiconsReply, error) {
	out := new(ListLexiconsReply)
	err := grpc.Invoke(ctx, "/codelingo.CodeLingo/ListLexicons", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeLingoClient) ListFacts(ctx context.Context, in *ListFactsRequest, opts ...grpc.CallOption) (*FactList, error) {
	out := new(FactList)
	err := grpc.Invoke(ctx, "/codelingo.CodeLingo/ListFacts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CodeLingo service

type CodeLingoServer interface {
	// Initialise session
	Session(context.Context, *SessionRequest) (*SessionReply, error)
	// Reviews files
	Review(context.Context, *ReviewRequest) (*ReviewReply, error)
	// Queries source code
	Query(context.Context, *QueryRequest) (*QueryReply, error)
	// Lists available lexicons
	ListLexicons(context.Context, *ListLexiconsRequest) (*ListLexiconsReply, error)
	ListFacts(context.Context, *ListFactsRequest) (*FactList, error)
}

func RegisterCodeLingoServer(s *grpc.Server, srv CodeLingoServer) {
	s.RegisterService(&_CodeLingo_serviceDesc, srv)
}

func _CodeLingo_Session_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeLingoServer).Session(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codelingo.CodeLingo/Session",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeLingoServer).Session(ctx, req.(*SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeLingo_Review_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeLingoServer).Review(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codelingo.CodeLingo/Review",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeLingoServer).Review(ctx, req.(*ReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeLingo_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeLingoServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codelingo.CodeLingo/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeLingoServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeLingo_ListLexicons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLexiconsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeLingoServer).ListLexicons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codelingo.CodeLingo/ListLexicons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeLingoServer).ListLexicons(ctx, req.(*ListLexiconsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeLingo_ListFacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeLingoServer).ListFacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codelingo.CodeLingo/ListFacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeLingoServer).ListFacts(ctx, req.(*ListFactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CodeLingo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "codelingo.CodeLingo",
	HandlerType: (*CodeLingoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Session",
			Handler:    _CodeLingo_Session_Handler,
		},
		{
			MethodName: "Review",
			Handler:    _CodeLingo_Review_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _CodeLingo_Query_Handler,
		},
		{
			MethodName: "ListLexicons",
			Handler:    _CodeLingo_ListLexicons_Handler,
		},
		{
			MethodName: "ListFacts",
			Handler:    _CodeLingo_ListFacts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "codelingo.proto",
}

func init() { proto.RegisterFile("codelingo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 857 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x55, 0xe1, 0x6e, 0x23, 0x35,
	0x10, 0xbe, 0x74, 0x9b, 0x36, 0x3b, 0x49, 0x4b, 0x71, 0xdb, 0xc3, 0x84, 0x03, 0x45, 0xab, 0x43,
	0x0a, 0x12, 0x0a, 0xa2, 0x20, 0x01, 0x27, 0xa4, 0xd3, 0xb5, 0x07, 0xd2, 0x49, 0x39, 0x08, 0x06,
	0x89, 0x7f, 0x48, 0xcb, 0x66, 0xd2, 0x98, 0xdb, 0xd8, 0x39, 0xdb, 0x49, 0x9b, 0xa7, 0xe0, 0x1d,
	0x78, 0x19, 0xde, 0x80, 0x7f, 0xbc, 0x0b, 0x1a, 0xdb, 0xbb, 0xd9, 0xdc, 0x85, 0x7f, 0xf3, 0x8d,
	0xbf, 0x19, 0x7b, 0x3e, 0x8f, 0xc7, 0xf0, 0x4e, 0xa1, 0xa7, 0x58, 0x4a, 0x75, 0xab, 0x47, 0x4b,
	0xa3, 0x9d, 0x66, 0x69, 0xed, 0xc8, 0x3e, 0x85, 0xb3, 0xb1, 0xb4, 0xee, 0xfb, 0xbc, 0x70, 0x56,
	0xe0, 0xeb, 0x15, 0x5a, 0xc7, 0x38, 0x1c, 0x97, 0x78, 0x2f, 0x0b, 0xad, 0x78, 0x6b, 0xd0, 0x1a,
	0xa6, 0xa2, 0x82, 0xd9, 0x9f, 0x2d, 0xe8, 0x10, 0x95, 0x42, 0xd8, 0x97, 0xd0, 0x9e, 0x51, 0x18,
	0x6f, 0x0d, 0x92, 0x61, 0xf7, 0xea, 0xa3, 0xd1, 0x76, 0x9b, 0x8a, 0xe3, 0x0d, 0xfb, 0x9d, 0x72,
	0x66, 0x23, 0x02, 0xb9, 0xff, 0x12, 0x60, 0xeb, 0x64, 0x67, 0x90, 0xbc, 0xc2, 0x4d, 0xdc, 0x86,
	0x4c, 0xf6, 0x09, 0xb4, 0xd7, 0x79, 0xb9, 0x42, 0x7e, 0x30, 0x68, 0x0d, 0xbb, 0x57, 0xe7, 0x8d,
	0xac, 0x37, 0x73, 0x59, 0x4e, 0x0d, 0x2a, 0x11, 0x18, 0x4f, 0x0e, 0xbe, 0x6e, 0x65, 0x03, 0xe8,
	0x54, 0x6e, 0x76, 0x01, 0xed, 0x82, 0x6c, 0x7f, 0xa0, 0x54, 0x04, 0x90, 0x5d, 0xc2, 0x39, 0x1d,
	0x65, 0x1c, 0x4a, 0xa8, 0x8a, 0xcc, 0x3e, 0x83, 0x77, 0x77, 0xdd, 0xcb, 0x72, 0xc3, 0xfa, 0xd0,
	0x89, 0xa5, 0xda, 0x98, 0xa4, 0xc6, 0xd9, 0x19, 0x9c, 0xfe, 0x8c, 0xd6, 0x4a, 0xad, 0xaa, 0x14,
	0x03, 0xe8, 0xd5, 0x1e, 0x8a, 0x7e, 0xab, 0x98, 0x2c, 0x83, 0xde, 0x4f, 0x2b, 0x34, 0x9b, 0x4a,
	0x59, 0x06, 0x87, 0x45, 0xf9, 0xba, 0x8c, 0x14, 0x6f, 0x67, 0x8f, 0x01, 0x22, 0x87, 0x72, 0x3c,
	0x84, 0x23, 0x83, 0x76, 0x55, 0xba, 0xc8, 0x89, 0x28, 0xfb, 0xfb, 0x00, 0x4e, 0x04, 0xae, 0x25,
	0xde, 0x55, 0xb9, 0xe2, 0x6e, 0xbd, 0xad, 0x74, 0x0c, 0x0e, 0xe7, 0xda, 0x56, 0x91, 0xde, 0x26,
	0x4d, 0xf4, 0x9d, 0x42, 0xe3, 0xe5, 0x4c, 0x45, 0x00, 0xc4, 0x34, 0xb8, 0xd4, 0x3c, 0x09, 0x4c,
	0xb2, 0x29, 0x9f, 0x9d, 0xe7, 0xfc, 0x30, 0xe4, 0xb3, 0xf3, 0x9c, 0x65, 0xd0, 0x9b, 0xc9, 0x12,
	0xed, 0x33, 0x35, 0x7d, 0x2e, 0x8d, 0xe5, 0x6d, 0xaf, 0xc8, 0x8e, 0x8f, 0x3d, 0x82, 0xd4, 0x60,
	0xb1, 0x32, 0x56, 0xae, 0x91, 0x1f, 0x0d, 0x5a, 0xc3, 0x8e, 0xd8, 0x3a, 0xa8, 0x93, 0x26, 0xb9,
	0x2b, 0xe6, 0x68, 0xf9, 0xb1, 0x0f, 0xae, 0x20, 0x7b, 0x0c, 0x27, 0xd2, 0x4e, 0x56, 0x65, 0x19,
	0xcb, 0xe1, 0x1d, 0x1f, 0xbb, 0xeb, 0x24, 0xd6, 0x72, 0x0b, 0x5f, 0x3c, 0xe7, 0xe9, 0xa0, 0x35,
	0x4c, 0xc4, 0xae, 0x93, 0x4e, 0xbe, 0x2e, 0x2c, 0x87, 0x70, 0xf2, 0x75, 0x61, 0xe9, 0x1e, 0xa7,
	0xda, 0xf9, 0xae, 0xe1, 0x5d, 0xef, 0xae, 0x71, 0x76, 0x02, 0xdd, 0x4a, 0xc8, 0x65, 0xb9, 0xc9,
	0xfe, 0x4d, 0xa0, 0xfd, 0xc2, 0xda, 0x15, 0x92, 0x28, 0x2a, 0x5f, 0x60, 0x25, 0x1f, 0xd9, 0xec,
	0x73, 0xe8, 0x2c, 0xb5, 0x95, 0x4e, 0x6a, 0x15, 0x1b, 0xf2, 0xb2, 0xd1, 0x90, 0x3e, 0x4e, 0xe4,
	0xea, 0x16, 0x45, 0x4d, 0xa3, 0x9a, 0x0b, 0xbd, 0x58, 0xa0, 0x72, 0x51, 0xde, 0x0a, 0x92, 0x56,
	0x85, 0xbb, 0xbf, 0xc6, 0x99, 0x36, 0x18, 0x75, 0xde, 0x3a, 0x7c, 0xef, 0x49, 0x85, 0xbf, 0xe0,
	0xbd, 0xe3, 0xed, 0x70, 0xe6, 0x0a, 0xd3, 0x5a, 0xe1, 0xee, 0x9f, 0xcd, 0x1c, 0x1a, 0x2f, 0x72,
	0x2a, 0x6a, 0xcc, 0xbe, 0x82, 0xe3, 0x05, 0x3a, 0x23, 0x8b, 0xa0, 0x71, 0xf7, 0xea, 0xc3, 0x37,
	0x4f, 0x38, 0x7a, 0x19, 0xd6, 0xc3, 0x3b, 0xac, 0xd8, 0x54, 0xaf, 0xcb, 0x6f, 0x2d, 0xef, 0xf8,
	0x9b, 0xf1, 0x36, 0xf9, 0x4a, 0xa9, 0x5e, 0x79, 0x9d, 0x53, 0xe1, 0x6d, 0x2a, 0x48, 0xe1, 0xdd,
	0x8d, 0x9e, 0xa2, 0x97, 0xb8, 0x23, 0x2a, 0x48, 0xcd, 0xb5, 0xa4, 0xfb, 0x8c, 0x1a, 0x07, 0x40,
	0xd7, 0x81, 0xc6, 0x54, 0x8d, 0x89, 0xc6, 0x50, 0x86, 0xa9, 0xb4, 0x45, 0x6e, 0xa6, 0xfc, 0x24,
	0x64, 0x88, 0x90, 0x2e, 0x38, 0x9a, 0x02, 0x73, 0xab, 0x15, 0x3f, 0xf5, 0x51, 0xbb, 0xce, 0xfe,
	0x13, 0xe8, 0x35, 0x4b, 0xd8, 0x33, 0x35, 0x2e, 0x9a, 0x53, 0x23, 0x6d, 0x0e, 0x88, 0xdf, 0x00,
	0xb6, 0xd7, 0x44, 0xd3, 0xc5, 0xba, 0xdc, 0x84, 0x37, 0xb2, 0x3b, 0x5d, 0x26, 0xf1, 0x02, 0x45,
	0x60, 0xb0, 0x8f, 0x21, 0x41, 0x35, 0xdd, 0x33, 0x86, 0x6a, 0x22, 0xad, 0x67, 0x7f, 0x40, 0xa7,
	0x72, 0xd0, 0x35, 0xd1, 0xe3, 0x68, 0x74, 0x51, 0x8d, 0xe9, 0x61, 0xff, 0x38, 0x9b, 0x59, 0x74,
	0x3e, 0x63, 0x22, 0x22, 0x22, 0xc5, 0xc7, 0x52, 0xa1, 0xef, 0x95, 0x44, 0x78, 0x9b, 0xb8, 0x37,
	0xba, 0x5c, 0x2d, 0x94, 0xef, 0x92, 0x44, 0x44, 0x74, 0xf5, 0xcf, 0x01, 0xa4, 0x24, 0xfc, 0x98,
	0xce, 0xc1, 0x9e, 0xc2, 0x71, 0x1c, 0x3f, 0xec, 0xfd, 0xc6, 0xf1, 0x76, 0x87, 0x54, 0xff, 0xbd,
	0x7d, 0x4b, 0xd4, 0xf8, 0x0f, 0xd8, 0xb7, 0x70, 0x14, 0x5e, 0x02, 0xe3, 0x0d, 0xd2, 0xce, 0x94,
	0xe9, 0x3f, 0xdc, 0xb3, 0x12, 0xa2, 0xbf, 0x81, 0xb6, 0x9f, 0x5b, 0xac, 0xb9, 0x43, 0x73, 0xda,
	0xf5, 0x2f, 0xdf, 0x5e, 0x08, 0xa1, 0x3f, 0x40, 0xaf, 0x39, 0x7b, 0x59, 0xf3, 0xeb, 0xd8, 0x33,
	0xab, 0xfb, 0x8f, 0xfe, 0x77, 0x3d, 0xe4, 0x7b, 0x0a, 0x69, 0xfd, 0x89, 0xb1, 0x0f, 0xde, 0x20,
	0x37, 0xbf, 0xb6, 0xfe, 0xf9, 0x9e, 0x4f, 0x2a, 0x7b, 0x70, 0x3d, 0x82, 0x0b, 0xa9, 0x1b, 0x4b,
	0xb7, 0x66, 0x59, 0x8c, 0x8a, 0xf2, 0xfa, 0xb4, 0x56, 0x7b, 0x42, 0x1f, 0xe7, 0xa4, 0xf5, 0xd7,
	0x41, 0x72, 0x33, 0xfe, 0xf5, 0xf7, 0x23, 0xff, 0x8f, 0x7e, 0xf1, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc2, 0x7d, 0xe9, 0x38, 0x5a, 0x07, 0x00, 0x00,
}
