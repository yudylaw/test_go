package handler

import (
	"encoding/json"
	"feed_api/common"
	"feed_api/service"
	api "git.inke.cn/video/panshi/panshi_api"
	"github.com/gin-gonic/gin"
	"strconv"
)

type FeedHandler struct {
	feedService *service.FeedService
}

func NewFeedHandle(service *service.FeedService) *FeedHandler {
	return &FeedHandler{feedService: service}
}

func (feedHandler *FeedHandler) SendFeedHandler(c *gin.Context) {
	uid, err := strconv.Atoi(c.Query("uid"))
	if err != nil || uid < 0 {
		common.RespJSON(c, api.ErrorCode_PARAM_ERROR)
		return
	}

	body, err := c.GetRawData()

	var bodyObj map[string]interface{}
	if err := json.Unmarshal(body, &bodyObj); err != nil {
		common.RespJSON(c, api.ErrorCode_ERROR)
		return
	}

	api.BusinessLog().Infof("SendFeed body", bodyObj)

	feedType := uint32(bodyObj["feed_type"].(float64))
	content := bodyObj["content"].(string)

	resp := feedHandler.feedService.Send(uint32(uid), feedType, content)

	api.BusinessLog().Infof("Send resp", resp)

	if resp.GetResult().GetCode() != api.ErrorCode_SUCCESS {
		common.RespJSON(c, resp.GetResult().GetCode())
		return
	}
	feed := &common.FeedItem{}
	feed.FeedId = resp.GetFeed().GetFeedId()
	feed.Uid = resp.GetFeed().GetUid()
	feed.FeedType = resp.GetFeed().GetFeedType()
	feed.CreateTime = resp.GetFeed().GetCreateTime()
	feed.Content = resp.GetFeed().GetContent()
	feed.Status = resp.GetFeed().GetStatus()
	feed.IsLiked = resp.GetFeed().GetIsLiked()
	feed.LikeNum = resp.GetFeed().GetLikeNum()
	api.GetResponseJson(c, api.ErrorCode_SUCCESS, feed)
}
