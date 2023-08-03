// Code generated by Kitex v0.3.2. DO NOT EDIT.

package feedservice

import (
	"context"
	"github.com/Alexdzk/dousheng/kitex_gen/feed"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Feed(ctx context.Context, Req *feed.FeedRequest, callOptions ...callopt.Option) (r *feed.FeedResponse, err error)
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
	return &kFeedServiceClient{
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

type kFeedServiceClient struct {
	*kClient
}

func (p *kFeedServiceClient) Feed(ctx context.Context, Req *feed.FeedRequest, callOptions ...callopt.Option) (r *feed.FeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Feed(ctx, Req)
}
