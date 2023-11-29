package cache

import (
	"VideoWeb2/database/db/dao"
	"VideoWeb2/database/db/model"
	"VideoWeb2/response"
	"VideoWeb2/types"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func CreateVideoByCache(ctx context.Context, video model.Video) error {
	key := VideoCacheKey(video.ID)
	jsonStr, err := json.Marshal(video)
	if err != nil {
		return err
	}
	RedisClient.Set(ctx, key, string(jsonStr), 300)
	return nil
}

func FindVideoByVidInCache(ctx context.Context, vid uint) (model.Video, error) {
	key := VideoCacheKey(vid)
	res := RedisClient.Get(ctx, key)
	var video model.Video
	//如果redis没有，则去数据库查找，查完后加入redis
	if res.Err() == redis.Nil {
		videoDao := dao.NewVideoDao(ctx)
		video, err := videoDao.FindVideoByVid(vid)
		if err != nil {
			return video, err
		}
		//加入缓存
		jsonStr, _ := json.Marshal(video)
		//fmt.Println(err)
		RedisClient.Set(ctx, key, string(jsonStr), 300)
		return video, nil
	}
	jsonStr := res.Val()
	_ = json.Unmarshal([]byte(jsonStr), &video)

	return video, nil
}

// FindVideoRanks 查找Rank中的video总数和所有vid
func FindVideoRanks(ctx context.Context) (int64, []response.VideoRankResp) {
	resp := make([]response.VideoRankResp, 0)
	count, _ := RedisClient.ZCard(ctx, "Rank").Result()
	keys, _ := RedisClient.ZRevRange(ctx, "Rank", 0, -1).Result()
	for _, key := range keys {
		val := RedisClient.ZScore(ctx, "Rank", key).Val()
		vid, _ := strconv.Atoi(key)
		resp = append(resp, response.VideoRankResp{
			Vid:   vid,
			Views: int64(int(val)),
		})
	}
	return count, resp
}

func SaveQueryVideoResult(ctx context.Context, list response.Response) error {
	u, _ := types.GetUserInfo(ctx)
	key := QueryVideoHistoryKey(u.ID)
	jsonStr, err := json.Marshal(list)
	if err != nil {
		return err
	}
	RedisClient.LPush(ctx, key, string(jsonStr))
	return nil
}

func QueryVideoResult(ctx context.Context) []response.Response {
	u, _ := types.GetUserInfo(ctx)
	key := QueryVideoHistoryKey(u.ID)
	res := RedisClient.LRange(ctx, key, 0, -1).Val()
	var resps []response.Response
	for _, v := range res {
		var resp response.Response
		_ = json.Unmarshal([]byte(v), &resp)
		resps = append(resps, resp)
	}
	return resps
}

// View 点击量
func View(ctx context.Context, vid uint) uint64 {
	countStr, _ := RedisClient.Get(ctx, VideoViewKey(vid)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func AddView(ctx context.Context, video *model.Video) {
	//增加点击量
	RedisClient.Incr(ctx, VideoViewKey(video.ID))
	//增加排行点击数
	RedisClient.ZIncrBy(ctx, "Rank", 1, strconv.Itoa(int(video.ID)))
}
