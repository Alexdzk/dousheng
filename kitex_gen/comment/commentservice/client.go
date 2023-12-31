// Code generated by Kitex v0.3.2. DO NOT EDIT.

package commentservice

import (
	"context"
	"github.com/Alexdzk/dousheng/kitex_gen/comment"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateComment(ctx context.Context, Req *comment.CreateCommentRequest, callOptions ...callopt.Option) (r *comment.CreateCommentResponse, err error)
	DeleteComment(ctx context.Context, Req *comment.DeleteCommentRequest, callOptions ...callopt.Option) (r *comment.DeleteCommentResponse, err error)
	CommentList(ctx context.Context, Req *comment.CommentListRequest, callOptions ...callopt.Option) (r *comment.CommentListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCommentServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCommentServiceClient struct {
	*kClient
}

func (p *kCommentServiceClient) CreateComment(ctx context.Context, Req *comment.CreateCommentRequest, callOptions ...callopt.Option) (r *comment.CreateCommentResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateComment(ctx, Req)
}

func (p *kCommentServiceClient) DeleteComment(ctx context.Context, Req *comment.DeleteCommentRequest, callOptions ...callopt.Option) (r *comment.DeleteCommentResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteComment(ctx, Req)
}

func (p *kCommentServiceClient) CommentList(ctx context.Context, Req *comment.CommentListRequest, callOptions ...callopt.Option) (r *comment.CommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, Req)
}
