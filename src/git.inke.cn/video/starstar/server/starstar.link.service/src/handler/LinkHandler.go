package handler

import (
	"golang.org/x/net/context"
	"net/http"
	"git.inke.cn/video/starstar/server/starstar.common.api"
	"git.inke.cn/video/starstar/server/starstar.link.service/src/client"
	log "git.inke.cn/BackendPlatform/golang/logging"
)

type LinkSlotAddrHandler struct{}

//多人连麦推流地址接口
func (self *LinkSlotAddrHandler) Serve(ctx context.Context, request *http.Request) (interface{}, int) {

	resp, code := client.ThriftDeal.GetLinkSlotAddr(ctx, 1521532180522220, 1, 1, "")

	log.Debugf("GetLinkSlotAddr, resp=%v, code=%d", resp, code)

	return api.GetResponse(api.ErrorCode_SUCCESS)
}