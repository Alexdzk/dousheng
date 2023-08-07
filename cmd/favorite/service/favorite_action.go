package service

import (
	"context"
	"errors"

	"github.com/Alexdzk/dousheng/dal/cache"
	"github.com/Alexdzk/dousheng/dal/db"
	"github.com/Alexdzk/dousheng/dal/mq"
	"github.com/Alexdzk/dousheng/kitex_gen/favorite"
	"github.com/Alexdzk/dousheng/pkg/constants"
	"github.com/Alexdzk/dousheng/pkg/jwt"
)

type FavoriteActionService struct {
	ctx context.Context
}

// NewFavoriteActionService new FavoriteActionService
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// FavoriteAction implement the like and unlike operations
func (s *FavoriteActionService) FavoriteAction(req *favorite.FavoriteActionRequest) error {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		return err
	}
	currentId := claim.Id

	videos, err := db.QueryVideoByVideoIds(s.ctx, []int64{req.VideoId})
	if err != nil {
		return err
	}
	if len(videos) == 0 {
		return errors.New("video not exist")
	}

	//若ActionType（操作类型）等于1，则向favorite表创建一条记录，同时向video表的目标video增加点赞数
	//若ActionType等于2，则向favorite表删除一条记录，同时向video表的目标video减少点赞数
	//若ActionType不等于1和2，则返回错误
	if req.ActionType == constants.Like {
		favorite := &db.FavoriteRaw{
			UserId:  currentId,
			VideoId: req.VideoId,
		}
		if !cache.FavoriteCheck(favorite.UserId, favorite.VideoId) {
			err := cache.VideoFavoriteAciton(favorite.UserId, favorite.VideoId, 1)
			if err != nil {
				return err
			}
			if err = mq.PublishFavoriteMsg(context.Background(), favorite); err != nil {
				return err
			}
		}
		// err := db.CreateFavorite(s.ctx, favorite, req.VideoId)
		if err != nil {
			return err
		}
	}
	if req.ActionType == constants.Unlike {
		err := db.DeleteFavorite(s.ctx, currentId, req.VideoId)
		if err != nil {
			return err
		}

	}
	if req.ActionType != constants.Like && req.ActionType != constants.Unlike {
		return errors.New("action type no equal 1 and 2")
	}
	return nil
}
