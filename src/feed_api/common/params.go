package common

type User struct {
	NiceName    int64  `json:"nickname""`
	Gender      int32  `json:"gender"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Audit       string `json:"audit"`
}

type FeedItem struct {
	FeedId     uint32 `json:"feed_id""`
	Uid        uint32 `json:"uid"`
	FeedType   uint32 `json:"feed_type"`
	Status     uint32 `json:"status"`
	CreateTime uint32 `json:"create_time"`
	Content    string `json:"content"`
	IsLiked    bool   `json:"isliked"`
	LikeNum    uint32 `json:"like_num"`
}
