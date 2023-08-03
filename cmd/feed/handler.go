package main

import (
	"context"

	"github.com/Alexdzk/dousheng/cmd/feed/service"
	"github.com/Alexdzk/dousheng/dal/pack"
	"github.com/Alexdzk/dousheng/kitex_gen/feed"
	"github.com/Alexdzk/dousheng/pkg/errno"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	resp = new(feed.FeedResponse)

	if req.LatestTime <= 0 {
		resp.BaseResp = pack.BuildFeedBaseResp(errno.ParamErr)
		return resp, nil
	}

	videos, nextTime, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp.BaseResp = pack.BuildFeedBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildFeedBaseResp(errno.Success)
	resp.VideoList = videos
	resp.NextTime = nextTime
	return resp, nil
}
