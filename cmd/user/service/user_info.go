package service

import (
	"context"
	"errors"

	"github.com/Alexdzk/dousheng/dal/db"
	"github.com/Alexdzk/dousheng/dal/pack"
	"github.com/Alexdzk/dousheng/kitex_gen/user"
	"github.com/Alexdzk/dousheng/pkg/constants"
	"github.com/Alexdzk/dousheng/pkg/jwt"
)

type UserInfoService struct {
	ctx context.Context
}

// NewUserInfoService new UserInfoService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{
		ctx: ctx,
	}
}

func (s *UserInfoService) UserInfo(req *user.UserInfoRequest) (*user.User, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, err := Jwt.CheckToken(req.Token)
	if err != nil {
		return nil, err
	}

	userIds := []int64{req.UserId}
	users, err := db.QueryUserByIds(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("user not exist")
	}
	user := users[0]

	relationMap, err := db.QueryRelationByIds(s.ctx, currentId, userIds)
	if err != nil {
		return nil, err
	}

	var isFollow bool
	_, ok := relationMap[req.UserId]
	if ok {
		isFollow = true
	} else {
		isFollow = false
	}

	userInfo := pack.UserInfo(user, isFollow)
	return userInfo, nil
}
