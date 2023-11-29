package service

import (
	"VideoWeb2/database/cache"
	"VideoWeb2/database/db/dao"
	"VideoWeb2/database/db/model"
	"VideoWeb2/es"
	"VideoWeb2/response"
	"VideoWeb2/types"
	"context"
	"strconv"
	"time"
)

type VideoService struct {
}

func (s *VideoService) VideoCreate(ctx context.Context, req *types.VideoCreateRequest) (interface{}, error) {
	videoDao := dao.NewVideoDao(ctx)
	key, err := Upload(req.FilePath)
	if err != nil {
		return response.BadResponse(err.Error()), err
	}
	u, err := types.GetUserInfo(ctx)
	if err != nil {
		return response.BadResponse("用户错误"), err
	}
	video := model.Video{
		Uid:      u.ID,
		Title:    req.Title,
		Types:    req.Types,
		Key:      key,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	err = videoDao.CreateVideo(&video)
	if err != nil {
		return response.BadResponse("存入数据库失败"), err
	}
	err = cache.CreateVideoByCache(ctx, video)
	if err != nil {
		return response.BadResponse("redis缓存失败"), err
	}
	videoForEs := model.VideoForEs{
		ID:       video.ID,
		Uid:      video.Uid,
		Title:    video.Title,
		Types:    video.Types,
		Views:    0,
		Key:      video.Key,
		CreateAt: video.CreateAt,
	}

	err = es.CreateVideoDocument(model.Video{}.IndexName(), &videoForEs)
	if err != nil {
		return response.BadResponse("es创建文档失败"), err
	}
	return response.SuccessResponse(), nil
}

func (s *VideoService) VideoWatch(ctx context.Context, req *types.VideoWatchRequest) (interface{}, error) {
	video, err := cache.FindVideoByVidInCache(ctx, req.Vid)
	if err != nil {
		return response.BadResponse("数据库中无该视频"), err
	}
	URL, err := GetURL(video.Key)
	if err != nil {
		return response.BadResponse("获取URL失败"), err
	}
	cache.AddView(ctx, &video)
	views := cache.View(ctx, video.ID)
	resp := response.VideoResp{
		ID:       video.ID,
		Uid:      video.Uid,
		Title:    video.Title,
		Types:    video.Types,
		URL:      URL,
		Views:    int(views),
		CreateAt: video.CreateAt,
	}
	DocID := strconv.Itoa(int(video.ID))
	err = es.UpdateViews(DocID, int(views))
	if err != nil {
		return response.BadResponse("es更新失败"), nil
	}
	return response.SuccessResponseWithData(resp), nil
}

func (s *VideoService) VideoRank(ctx context.Context, req *types.VideoRankRequest) (interface{}, error) {
	count, resp := cache.FindVideoRanks(ctx)
	var list []response.VideoResp
	for _, v := range resp {
		video, _ := cache.FindVideoByVidInCache(ctx, uint(v.Vid))
		url, _ := GetURL(video.Key)
		list = append(list, response.VideoResp{
			ID:       uint(v.Vid),
			Uid:      video.Uid,
			Title:    video.Title,
			Types:    video.Types,
			URL:      url,
			Views:    int(v.Views),
			CreateAt: video.CreateAt,
		})
	}
	return response.ListResp(list, count), nil
}

func (s *VideoService) QueryVideos(ctx context.Context, filter *types.Filter) (interface{}, error) {

	videos, count, err := es.QueryVideoDocuments(model.Video{}.IndexName(), filter)
	if err != nil {
		return response.BadResponse("es查询失败"), err
	}
	list := make([]*response.VideoResp, 0)
	for _, v := range videos {
		url, _ := GetURL(v.Key)
		list = append(list, &response.VideoResp{
			ID:       v.ID,
			Uid:      v.Uid,
			Title:    v.Title,
			Types:    v.Types,
			URL:      url,
			Views:    v.Views,
			CreateAt: v.CreateAt,
		})
	}
	res := response.ListResp(list, count)
	err = cache.SaveQueryVideoResult(ctx, res)
	if err != nil {
		return response.BadResponse("保存redis失败"), err
	}
	return res, nil
}
