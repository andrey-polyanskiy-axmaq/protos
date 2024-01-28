// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0--rc1
// source: sso/sso.proto

package ssov1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FeedbackClient is the client API for Feedback service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedbackClient interface {
	Suggestion(ctx context.Context, in *SuggestionRequest, opts ...grpc.CallOption) (*SuggestionResponse, error)
	Complaint(ctx context.Context, in *ComplaintRequest, opts ...grpc.CallOption) (*ComplaintResponse, error)
	SugComAnswer(ctx context.Context, in *AnswerRequest, opts ...grpc.CallOption) (*AnswerResponse, error)
}

type feedbackClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedbackClient(cc grpc.ClientConnInterface) FeedbackClient {
	return &feedbackClient{cc}
}

func (c *feedbackClient) Suggestion(ctx context.Context, in *SuggestionRequest, opts ...grpc.CallOption) (*SuggestionResponse, error) {
	out := new(SuggestionResponse)
	err := c.cc.Invoke(ctx, "/feedback.Feedback/Suggestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackClient) Complaint(ctx context.Context, in *ComplaintRequest, opts ...grpc.CallOption) (*ComplaintResponse, error) {
	out := new(ComplaintResponse)
	err := c.cc.Invoke(ctx, "/feedback.Feedback/Complaint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackClient) SugComAnswer(ctx context.Context, in *AnswerRequest, opts ...grpc.CallOption) (*AnswerResponse, error) {
	out := new(AnswerResponse)
	err := c.cc.Invoke(ctx, "/feedback.Feedback/SugComAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedbackServer is the server API for Feedback service.
// All implementations must embed UnimplementedFeedbackServer
// for forward compatibility
type FeedbackServer interface {
	Suggestion(context.Context, *SuggestionRequest) (*SuggestionResponse, error)
	Complaint(context.Context, *ComplaintRequest) (*ComplaintResponse, error)
	SugComAnswer(context.Context, *AnswerRequest) (*AnswerResponse, error)
	mustEmbedUnimplementedFeedbackServer()
}

// UnimplementedFeedbackServer must be embedded to have forward compatible implementations.
type UnimplementedFeedbackServer struct {
}

func (UnimplementedFeedbackServer) Suggestion(context.Context, *SuggestionRequest) (*SuggestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Suggestion not implemented")
}
func (UnimplementedFeedbackServer) Complaint(context.Context, *ComplaintRequest) (*ComplaintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Complaint not implemented")
}
func (UnimplementedFeedbackServer) SugComAnswer(context.Context, *AnswerRequest) (*AnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SugComAnswer not implemented")
}
func (UnimplementedFeedbackServer) mustEmbedUnimplementedFeedbackServer() {}

// UnsafeFeedbackServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedbackServer will
// result in compilation errors.
type UnsafeFeedbackServer interface {
	mustEmbedUnimplementedFeedbackServer()
}

func RegisterFeedbackServer(s grpc.ServiceRegistrar, srv FeedbackServer) {
	s.RegisterService(&Feedback_ServiceDesc, srv)
}

func _Feedback_Suggestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServer).Suggestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.Feedback/Suggestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServer).Suggestion(ctx, req.(*SuggestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feedback_Complaint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComplaintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServer).Complaint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.Feedback/Complaint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServer).Complaint(ctx, req.(*ComplaintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feedback_SugComAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServer).SugComAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.Feedback/SugComAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServer).SugComAnswer(ctx, req.(*AnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Feedback_ServiceDesc is the grpc.ServiceDesc for Feedback service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Feedback_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "feedback.Feedback",
	HandlerType: (*FeedbackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Suggestion",
			Handler:    _Feedback_Suggestion_Handler,
		},
		{
			MethodName: "Complaint",
			Handler:    _Feedback_Complaint_Handler,
		},
		{
			MethodName: "SugComAnswer",
			Handler:    _Feedback_SugComAnswer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/sso.proto",
}
