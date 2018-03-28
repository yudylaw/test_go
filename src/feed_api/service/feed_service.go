package service

import (
	_ "flag"
	rpc "git.inke.cn/inkelogic/rpc-go"
	api "git.inke.cn/video/panshi/panshi_api"
	"git.inke.cn/video/panshi/panshi_api/feed"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
)

type FeedService struct {
	client feed.FeedServiceClient
}

func NewFeedService() *FeedService {
	requestOption := rpc.NewRequestOptional()
	requestOption.SetTimeOut(30000)

	conn, err := rpc.DialService("feed_service", requestOption)
	if err != nil {
		api.ErrorLog().Error("failed to init feed_service client, err:", err)
		return nil
	}

	client := feed.NewFeedServiceClient(conn)
	obj := &FeedService{client: client}

	return obj
}

func (m *FeedService) GetClient() payment.FeedServiceClient {
	return m.client
}

func (m *FeedService) Send(uid uint32, feedType uint32, content string) *feed.SendResponse {
	r := feed.SendRequest{Uid: proto.Uint32(uid), Content: proto.String(content), FeedType: proto.Uint32(feedType)}

	rsp, er := m.client.Send(context.TODO(), &r)
	if er != nil {
		api.ErrorLog().Error("call Send error ,err:", er, rsp)
		rsp = &feed.SendResponse{}
		rsp.Result = &api.BaseResponse{Code: api.ErrorCode_ERROR.Enum()}
	}

	return rsp
}
