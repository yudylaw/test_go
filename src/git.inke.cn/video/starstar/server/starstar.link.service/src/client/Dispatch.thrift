namespace cpp live_dispatch

enum RescodeType
{
	RES_SUCCESS = 0,                	//执行处理成功，返回OK
	RES_DEFAULT_EXCEPTION = 1,  		//失败，默认异常
	RES_PARAM_ERROR = 2,  		//失败，默认异常
}


exception DispatchThriftException
{
        1: i32 whatOp,
	2: string why,
}

struct PublishRequest
{
  	1: i64 uid,
  	2: string cv,
  	3: i16 proto,
  	4: string ip,
  	5: string atom,
}

struct LinkSlotRequest
{
  	1: i64 liveid,
  	2: i16 slot,
  	3: i16 incr,
  	4: string extra,
}
struct LinkAddrRequest
{
  	1: i64 liveid,
  	2: string extra,
}

struct StreamPullAddrRequest {
        1: string url, //主路麦地址
        2: list<string> slot_urls,  //当前辅路麦地址
        3: string extra,
}

struct SlotAddr {
        1: string addr
}

struct StreamPullAddrResponse {
        1: i32 rescode, //主路麦地址
        2: list<SlotAddr> addrs,
}

struct UserSettingRequest
{
  	1: i64 uid,
}


struct PublishAddrResponse {
        1: i64 rescode,
        2: i16 link,
        3: string publish_pre,
        4: string publish_suf,
        5: string stream_pre,
        6: string stream_suf,
        7: string extra;
}

struct LinkSlotResponse {
        1: i64 rescode,
        2: string publish_addr,
        3: string stream_addr,
        4: string stream_addr_cdn,
}

struct LinkAddrResponse {
        1: i64 rescode,
        2: string link_addr,
        3: string extra,
}

struct UserSettingResponse {
        1: i64 rescode,
        2: i16 link,
        3: i16 multi,
        4: i16 force_link
}

service DispatchService
{
	PublishAddrResponse getPublishAddr(1: PublishRequest request) throws (1:DispatchThriftException ex);
	LinkSlotResponse getLinkSlotAddr(1: LinkSlotRequest request) throws (1:DispatchThriftException ex);
	LinkAddrResponse getLinkAddr(1: LinkAddrRequest request) throws (1:DispatchThriftException ex);
	UserSettingResponse getUserSetting(1: UserSettingRequest request) throws (1:DispatchThriftException ex);
	StreamPullAddrResponse getStreamPullAddr(1: StreamPullAddrRequest request) throws (1:DispatchThriftException ex);
}
