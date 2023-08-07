package cache

import (
	"context"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}

func VideoFavoriteAciton(userId int64, videoId int64, status int64) error {
	if FavoriteCheck(userId, videoId) {
		return nil
	}
	err := rdb.HIncrBy(context.Background(), "videoFavoriteCount", strconv.FormatInt(videoId, 10), status).Err()
	if err != nil {
		return err
	}
	key := "FavoriteCheck:" + strconv.FormatInt(videoId, 10)
	err = rdb.Do(context.Background(), "setbit", key, strconv.FormatInt(userId, 10), status).Err()
	if err != nil {
		klog.Error("FavoriteCheck error" + err.Error())
		return err
	}
	return nil
}

func GetVideoFavoriteCount(videoId int64) (int64, error) {
	val, err := rdb.HGet(context.Background(), "videoFavoriteCount", strconv.FormatInt(videoId, 10)).Result()
	res, _ := strconv.ParseInt(val, 10, 64)
	return res, err
}

func FavoriteCheck(userId int64, videoId int64) bool {
	key := "FavoriteCheck:" + strconv.FormatInt(videoId, 10)
	var val int
	val, err := rdb.Do(context.Background(), "EXISTS", key).Int()
	if val == 0 {
		return false
	}

	val, err = rdb.Do(context.Background(), "getbit", key, strconv.FormatInt(userId, 10), 1).Int()
	if err != nil {
		klog.Error("FavoriteCheck error" + err.Error())
		return true
	}
	if val == 1 {
		return true
	} else {
		return false
	}
}
