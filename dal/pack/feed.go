package pack

import (
	"time"

	"github.com/Alexdzk/dousheng/dal/db"
	"github.com/Alexdzk/dousheng/kitex_gen/feed"
)

// VideoInfo pack video list info
func VideoInfo(currentId int64, videoData []*db.VideoRaw, userMap map[int64]*db.UserRaw, favoriteMap map[int64]*db.FavoriteRaw, relationMap map[int64]*db.RelationRaw) ([]*feed.Video, int64) {
	videoList := make([]*feed.Video, 0)
	var nextTime int64
	for _, video := range videoData {
		videoUser, ok := userMap[video.UserId]
		if !ok {
			videoUser = &db.UserRaw{
				Name:          "未知用户",
				FollowCount:   0,
				FollowerCount: 0,
			}
			videoUser.ID = 0
		}

		var isFavorite bool = false
		var isFollow bool = false

		if currentId != -1 {
			_, ok := favoriteMap[int64(video.ID)]
			if ok {
				isFavorite = true
			}
			_, ok = relationMap[video.UserId]
			if ok {
				isFollow = true
			}
		}
		videoList = append(videoList, &feed.Video{
			Id: int64(video.ID),
			Author: &feed.User{
				Id:            int64(videoUser.ID),
				Name:          videoUser.Name,
				FollowCount:   videoUser.FollowCount,
				FollowerCount: videoUser.FollowerCount,
				IsFollow:      isFollow,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		})
	}

	if len(videoData) == 0 {
		nextTime = time.Now().UnixMilli()
	} else {
		nextTime = videoData[len(videoData)-1].UpdatedAt.UnixMilli()
	}

	return videoList, nextTime
}
