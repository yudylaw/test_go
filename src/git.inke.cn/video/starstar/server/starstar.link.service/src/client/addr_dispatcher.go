package client

import (
	"context"
	_ "fmt"
	"net"
	"time"

	"git.inke.cn/video/starstar/server/starstar.link.service/src/define"
	"git.inke.cn/video/starstar/server/starstar.link.service/src/dispatch"
	log "git.inke.cn/BackendPlatform/golang/logging"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/silenceper/pool"
	"git.inke.cn/video/starstar/server/starstar.common.api"
)

var (
	ThriftDeal *AddrThrift
)

type AddrThrift struct {
	thriftClient pool.Pool
}

func InitThriftDeal() error {
	ThriftDeal = &AddrThrift{}
	return ThriftDeal.Init()
}

func (g *AddrThrift) Init() error {
	poolConfig := &pool.PoolConfig{
		//InitialCap:  32,
		//MaxCap:      64, thrift.TClient
		InitialCap:  12,
		MaxCap:      12,
		IdleTimeout: 180 * time.Second,
		Factory: func() (interface{}, error) {
			trans, err := thrift.NewTSocketTimeout(net.JoinHostPort(define.Config.AddrThriftHost,
				define.Config.AddrThriftPort),
				5000*time.Millisecond)
			if err != nil {
				return nil, err
			}

			if err := trans.Open(); err != nil {
				return nil, err
			}

			protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
			client := dispatch.NewDispatchServiceClientFactory(trans, protocolFactory)

			return client, nil
		},

		Close: func(v interface{}) error {
			//todo
			//v.(*dispatch.DispatchServiceClient)..Close()
			return nil
		},
	}

	p, err := pool.NewChannelPool(poolConfig)
	g.thriftClient = p
	if err != nil {
		log.Errorf("InitThrift error:%v", err)
		return err
	}
	log.Infof("InitThrift success")
	return nil
}

func (g *AddrThrift) GetLinkSlotAddr(ctx context.Context, liveid int64, slot int16, incr int16, extra string) (interface {}, api.ErrorCode) {
	req := dispatch.LinkSlotRequest{}
	req.Liveid = liveid
	req.Slot = slot
	req.Incr = incr
	req.Extra = extra

	client, err := g.thriftClient.Get()
	defer g.thriftClient.Put(client)

	resp, err := client.(*dispatch.DispatchServiceClient).GetLinkSlotAddr(&req)
	if err != nil {
		log.Errorf("GetLinkSlotAddr error, liveid:%+v, err:%+v", liveid, err)
		return nil, api.ErrorCode_ERROR
	}

	if resp.Rescode != 0 {
		log.Errorf("GetLinkSlotAddr resp.Rescode error, liveidv, resp:%+v, err:%+v", liveid, resp, err)
		return nil, api.ErrorCode_ERROR
	}

	addr := define.LinkSlotAddrResponse{}
	addr.PublishAddr = resp.PublishAddr
	addr.StreamAddr = resp.StreamAddr
	addr.StreamAddrCdn = resp.StreamAddrCdn

	return addr, api.ErrorCode_SUCCESS
}
